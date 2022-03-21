// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: api/proto/resources/subnet.proto

package resources

import (
	common "github.com/multycloud/multy/api/proto/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateSubnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resources []*CloudSpecificSubnetArgs `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *CreateSubnetRequest) Reset() {
	*x = CreateSubnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_subnet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubnetRequest) ProtoMessage() {}

func (x *CreateSubnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_subnet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubnetRequest.ProtoReflect.Descriptor instead.
func (*CreateSubnetRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_subnet_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSubnetRequest) GetResources() []*CloudSpecificSubnetArgs {
	if x != nil {
		return x.Resources
	}
	return nil
}

type ReadSubnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ReadSubnetRequest) Reset() {
	*x = ReadSubnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_subnet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSubnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSubnetRequest) ProtoMessage() {}

func (x *ReadSubnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_subnet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSubnetRequest.ProtoReflect.Descriptor instead.
func (*ReadSubnetRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_subnet_proto_rawDescGZIP(), []int{1}
}

func (x *ReadSubnetRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type UpdateSubnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string                     `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Resources  []*CloudSpecificSubnetArgs `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *UpdateSubnetRequest) Reset() {
	*x = UpdateSubnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_subnet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSubnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSubnetRequest) ProtoMessage() {}

func (x *UpdateSubnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_subnet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSubnetRequest.ProtoReflect.Descriptor instead.
func (*UpdateSubnetRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_subnet_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSubnetRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateSubnetRequest) GetResources() []*CloudSpecificSubnetArgs {
	if x != nil {
		return x.Resources
	}
	return nil
}

type DeleteSubnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeleteSubnetRequest) Reset() {
	*x = DeleteSubnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_subnet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubnetRequest) ProtoMessage() {}

func (x *DeleteSubnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_subnet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubnetRequest.ProtoReflect.Descriptor instead.
func (*DeleteSubnetRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_subnet_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteSubnetRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type CloudSpecificSubnetArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters *common.CloudSpecificResourceCommonArgs `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	Name             string                                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CidrBlock        string                                  `protobuf:"bytes,3,opt,name=cidr_block,json=cidrBlock,proto3" json:"cidr_block,omitempty"`
	VirtualNetworkId string                                  `protobuf:"bytes,4,opt,name=virtual_network_id,json=virtualNetworkId,proto3" json:"virtual_network_id,omitempty"`
	AvailabilityZone int32                                   `protobuf:"varint,5,opt,name=availability_zone,json=availabilityZone,proto3" json:"availability_zone,omitempty"`
}

func (x *CloudSpecificSubnetArgs) Reset() {
	*x = CloudSpecificSubnetArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_subnet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudSpecificSubnetArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSpecificSubnetArgs) ProtoMessage() {}

func (x *CloudSpecificSubnetArgs) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_subnet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSpecificSubnetArgs.ProtoReflect.Descriptor instead.
func (*CloudSpecificSubnetArgs) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_subnet_proto_rawDescGZIP(), []int{4}
}

func (x *CloudSpecificSubnetArgs) GetCommonParameters() *common.CloudSpecificResourceCommonArgs {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *CloudSpecificSubnetArgs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CloudSpecificSubnetArgs) GetCidrBlock() string {
	if x != nil {
		return x.CidrBlock
	}
	return ""
}

func (x *CloudSpecificSubnetArgs) GetVirtualNetworkId() string {
	if x != nil {
		return x.VirtualNetworkId
	}
	return ""
}

func (x *CloudSpecificSubnetArgs) GetAvailabilityZone() int32 {
	if x != nil {
		return x.AvailabilityZone
	}
	return 0
}

type CloudSpecificSubnetResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters *common.CloudSpecificCommonResourceParameters `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	Name             string                                        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CidrBlock        string                                        `protobuf:"bytes,3,opt,name=cidr_block,json=cidrBlock,proto3" json:"cidr_block,omitempty"`
	VirtualNetworkId string                                        `protobuf:"bytes,4,opt,name=virtual_network_id,json=virtualNetworkId,proto3" json:"virtual_network_id,omitempty"`
	AvailabilityZone int32                                         `protobuf:"varint,5,opt,name=availability_zone,json=availabilityZone,proto3" json:"availability_zone,omitempty"`
}

func (x *CloudSpecificSubnetResource) Reset() {
	*x = CloudSpecificSubnetResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_subnet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudSpecificSubnetResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudSpecificSubnetResource) ProtoMessage() {}

func (x *CloudSpecificSubnetResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_subnet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudSpecificSubnetResource.ProtoReflect.Descriptor instead.
func (*CloudSpecificSubnetResource) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_subnet_proto_rawDescGZIP(), []int{5}
}

func (x *CloudSpecificSubnetResource) GetCommonParameters() *common.CloudSpecificCommonResourceParameters {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *CloudSpecificSubnetResource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CloudSpecificSubnetResource) GetCidrBlock() string {
	if x != nil {
		return x.CidrBlock
	}
	return ""
}

func (x *CloudSpecificSubnetResource) GetVirtualNetworkId() string {
	if x != nil {
		return x.VirtualNetworkId
	}
	return ""
}

func (x *CloudSpecificSubnetResource) GetAvailabilityZone() int32 {
	if x != nil {
		return x.AvailabilityZone
	}
	return 0
}

type SubnetResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonParameters *common.CommonResourceParameters `protobuf:"bytes,1,opt,name=common_parameters,json=commonParameters,proto3" json:"common_parameters,omitempty"`
	Resources        []*CloudSpecificSubnetResource   `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *SubnetResource) Reset() {
	*x = SubnetResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_resources_subnet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubnetResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetResource) ProtoMessage() {}

func (x *SubnetResource) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_resources_subnet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetResource.ProtoReflect.Descriptor instead.
func (*SubnetResource) Descriptor() ([]byte, []int) {
	return file_api_proto_resources_subnet_proto_rawDescGZIP(), []int{6}
}

func (x *SubnetResource) GetCommonParameters() *common.CommonResourceParameters {
	if x != nil {
		return x.CommonParameters
	}
	return nil
}

func (x *SubnetResource) GetResources() []*CloudSpecificSubnetResource {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_api_proto_resources_subnet_proto protoreflect.FileDescriptor

var file_api_proto_resources_subnet_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a,
	0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x34, 0x0a, 0x11, 0x52, 0x65, 0x61,
	0x64, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22,
	0x82, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x65,
	0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x53,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x87, 0x02, 0x0a,
	0x17, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x53, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x5e, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x12, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0x91, 0x02, 0x0a, 0x1b, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x2c, 0x0a, 0x12, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x2b, 0x0a,
	0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x7a, 0x6f,
	0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x22, 0xb9, 0x01, 0x0a, 0x0e, 0x53,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x57, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d,
	0x75, 0x6c, 0x74, 0x79, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x64, 0x65, 0x76, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x79, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x5c, 0x0a, 0x17, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x79, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x75, 0x6c, 0x74, 0x79,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_resources_subnet_proto_rawDescOnce sync.Once
	file_api_proto_resources_subnet_proto_rawDescData = file_api_proto_resources_subnet_proto_rawDesc
)

func file_api_proto_resources_subnet_proto_rawDescGZIP() []byte {
	file_api_proto_resources_subnet_proto_rawDescOnce.Do(func() {
		file_api_proto_resources_subnet_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_resources_subnet_proto_rawDescData)
	})
	return file_api_proto_resources_subnet_proto_rawDescData
}

var file_api_proto_resources_subnet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_resources_subnet_proto_goTypes = []interface{}{
	(*CreateSubnetRequest)(nil),                          // 0: dev.multy.resources.CreateSubnetRequest
	(*ReadSubnetRequest)(nil),                            // 1: dev.multy.resources.ReadSubnetRequest
	(*UpdateSubnetRequest)(nil),                          // 2: dev.multy.resources.UpdateSubnetRequest
	(*DeleteSubnetRequest)(nil),                          // 3: dev.multy.resources.DeleteSubnetRequest
	(*CloudSpecificSubnetArgs)(nil),                      // 4: dev.multy.resources.CloudSpecificSubnetArgs
	(*CloudSpecificSubnetResource)(nil),                  // 5: dev.multy.resources.CloudSpecificSubnetResource
	(*SubnetResource)(nil),                               // 6: dev.multy.resources.SubnetResource
	(*common.CloudSpecificResourceCommonArgs)(nil),       // 7: dev.multy.common.CloudSpecificResourceCommonArgs
	(*common.CloudSpecificCommonResourceParameters)(nil), // 8: dev.multy.common.CloudSpecificCommonResourceParameters
	(*common.CommonResourceParameters)(nil),              // 9: dev.multy.common.CommonResourceParameters
}
var file_api_proto_resources_subnet_proto_depIdxs = []int32{
	4, // 0: dev.multy.resources.CreateSubnetRequest.resources:type_name -> dev.multy.resources.CloudSpecificSubnetArgs
	4, // 1: dev.multy.resources.UpdateSubnetRequest.resources:type_name -> dev.multy.resources.CloudSpecificSubnetArgs
	7, // 2: dev.multy.resources.CloudSpecificSubnetArgs.common_parameters:type_name -> dev.multy.common.CloudSpecificResourceCommonArgs
	8, // 3: dev.multy.resources.CloudSpecificSubnetResource.common_parameters:type_name -> dev.multy.common.CloudSpecificCommonResourceParameters
	9, // 4: dev.multy.resources.SubnetResource.common_parameters:type_name -> dev.multy.common.CommonResourceParameters
	5, // 5: dev.multy.resources.SubnetResource.resources:type_name -> dev.multy.resources.CloudSpecificSubnetResource
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_resources_subnet_proto_init() }
func file_api_proto_resources_subnet_proto_init() {
	if File_api_proto_resources_subnet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_resources_subnet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubnetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resources_subnet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSubnetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resources_subnet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSubnetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resources_subnet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSubnetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resources_subnet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudSpecificSubnetArgs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resources_subnet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudSpecificSubnetResource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_resources_subnet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubnetResource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_resources_subnet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_resources_subnet_proto_goTypes,
		DependencyIndexes: file_api_proto_resources_subnet_proto_depIdxs,
		MessageInfos:      file_api_proto_resources_subnet_proto_msgTypes,
	}.Build()
	File_api_proto_resources_subnet_proto = out.File
	file_api_proto_resources_subnet_proto_rawDesc = nil
	file_api_proto_resources_subnet_proto_goTypes = nil
	file_api_proto_resources_subnet_proto_depIdxs = nil
}