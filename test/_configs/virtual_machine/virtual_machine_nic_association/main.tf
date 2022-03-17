resource "aws_vpc" "example_vn_aws" {
  tags = {
    "Name" = "example_vn"
  }

  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
}
resource "aws_internet_gateway" "example_vn_aws" {
  tags = {
    "Name" = "example_vn"
  }

  vpc_id = aws_vpc.example_vn_aws.id
}
resource "aws_default_security_group" "example_vn_aws" {
  tags = {
    "Name" = "example_vn"
  }

  vpc_id = aws_vpc.example_vn_aws.id

  ingress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
    self        = true
  }

  egress {
    protocol    = "-1"
    from_port   = 0
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
    self        = true
  }
}
resource "aws_network_interface" "nic_aws" {
  tags = {
    "Name" = "test-nic"
  }

  subnet_id = "${aws_subnet.subnet_aws.id}"
}
resource "aws_subnet" "subnet_aws" {
  tags = {
    "Name" = "subnet"
  }

  cidr_block        = "10.0.2.0/24"
  vpc_id            = aws_vpc.example_vn_aws.id
  availability_zone = "eu-west-1b"
}
resource "aws_iam_role" "vm_aws" {
  tags               = { "Name" = "test-vm" }
  name               = "iam_for_vm_vm"
  assume_role_policy = "{\"Statement\":[{\"Action\":[\"sts:AssumeRole\"],\"Effect\":\"Allow\",\"Principal\":{\"Service\":\"ec2.amazonaws.com\"}}],\"Version\":\"2012-10-17\"}"
}
resource "aws_iam_instance_profile" "vm_aws" {
  name = "iam_for_vm_vm"
  role = aws_iam_role.vm_aws.name
}
resource "aws_instance" "vm_aws" {
  tags = {
    "Name" = "test-vm"
  }

  ami              = "ami-09d4a659cdd8677be"
  instance_type    = "t2.nano"
  user_data_base64 = base64encode("echo 'Hello World'")

  network_interface {
    network_interface_id = "${aws_network_interface.nic_aws.id}"
    device_index         = 0
  }
  iam_instance_profile = aws_iam_instance_profile.vm_aws.id
}
resource "azurerm_virtual_network" "example_vn_azure" {
  resource_group_name = azurerm_resource_group.vn-rg.name
  name                = "example_vn"
  location            = "northeurope"
  address_space       = ["10.0.0.0/16"]
}
resource "azurerm_route_table" "example_vn_azure" {
  resource_group_name = azurerm_resource_group.vn-rg.name
  name                = "example_vn"
  location            = "northeurope"

  route {
    name           = "local"
    address_prefix = "0.0.0.0/0"
    next_hop_type  = "VnetLocal"
  }
}
resource "azurerm_network_interface" "nic_azure" {
  resource_group_name = azurerm_resource_group.nic-rg.name
  name                = "test-nic"
  location            = "northeurope"

  ip_configuration {
    name                          = "internal"
    private_ip_address_allocation = "Dynamic"
    subnet_id                     = "${azurerm_subnet.subnet_azure.id}"
    primary                       = true
  }
}
resource "azurerm_resource_group" "nic-rg" {
  name     = "nic-rg"
  location = "northeurope"
}
resource "azurerm_subnet" "subnet_azure" {
  resource_group_name  = azurerm_resource_group.vn-rg.name
  name                 = "subnet"
  address_prefixes     = ["10.0.2.0/24"]
  virtual_network_name = azurerm_virtual_network.example_vn_azure.name
}
resource "azurerm_subnet_route_table_association" "subnet_azure" {
  subnet_id      = "${azurerm_subnet.subnet_azure.id}"
  route_table_id = "${azurerm_route_table.example_vn_azure.id}"
}
resource "random_password" "vm_azure" {
  length  = 16
  special = true
  upper   = true
  lower   = true
  number  = true
}
resource "azurerm_linux_virtual_machine" "vm_azure" {
  resource_group_name   = azurerm_resource_group.vm-rg.name
  name                  = "test-vm"
  location              = "northeurope"
  size                  = "Standard_B1ls"
  network_interface_ids = ["${azurerm_network_interface.nic_azure.id}"]
  custom_data           = base64encode("echo 'Hello World'")

  os_disk {
    caching              = "None"
    storage_account_type = "Standard_LRS"
  }

  admin_username = "adminuser"
  admin_password = random_password.vm_azure.result

  source_image_reference {
    publisher = "OpenLogic"
    offer     = "CentOS"
    sku       = "7_9-gen2"
    version   = "latest"
  }
  identity {
    type = "SystemAssigned"
  }
  disable_password_authentication = false
}
resource "azurerm_resource_group" "vm-rg" {
  name     = "vm-rg"
  location = "northeurope"
}
resource "azurerm_resource_group" "vn-rg" {
  name     = "vn-rg"
  location = "northeurope"
}
provider "aws" {
  region = "eu-west-1"
}
provider "azurerm" {
  features {}
}