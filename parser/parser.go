package parser

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
	"log"
	"multy-go/resources/common"
	"multy-go/validate"
	"multy-go/variables"
)

type Parser struct {
	CliVars variables.CommandLineVariables
}

type MultyResource struct {
	Type            string   `hcl:"type,label"`
	ID              string   `hcl:"id,label"`
	HCLBody         hcl.Body `hcl:",remain"`
	Dependencies    []MultyResourceDependency
	DefinitionRange hcl.Range
}

type MultyResourceDependency struct {
	From        *MultyResource
	To          *MultyResource
	SourceRange hcl.Range
}

type MultyConfig struct {
	Location                 string
	DefaultResourceGroupName *hcl.Expression
	Clouds                   []string
}

type ParsedConfig struct {
	Variables      []ParsedVariable
	MultyResources []*MultyResource
	GlobalConfig   hcl.Body
}

type ParsedVariable struct {
	Name  string
	Value cty.Value
}

type multyConfig struct {
	HCLBody hcl.Body `hcl:",remain"`
}

type config struct {
	Variables      []variables.Variable `hcl:"variable,block"`
	MultyResources []*MultyResource     `hcl:"multy,block"`
	MultyConfig    *multyConfig         `hcl:"config,block"`
}

func (p *Parser) Parse(filepaths []string) ParsedConfig {

	var configs []config
	parser := hclparse.NewParser()
	resourcesById := map[string]*MultyResource{}
	for _, filepath := range filepaths {
		var c config
		blocks := p.parseSingleFile(filepath, parser, &c)
		for _, r := range c.MultyResources {
			resourceId := fmt.Sprintf("%s", r.ID)
			if dup, ok := resourcesById[resourceId]; ok {
				validate.LogFatalWithSourceRange(dup.DefinitionRange, "duplicate resource found")
			}
			resourcesById[resourceId] = r
		}
		for _, block := range blocks {
			resourcesById[fmt.Sprintf("%s", block.Labels[1])].DefinitionRange = block.DefRange
		}
		configs = append(configs, c)
	}

	c := mergeConfigs(configs)

	for _, r := range c.MultyResources {
		r.Dependencies = findDependencies(r, resourcesById)
	}

	result := ParsedConfig{
		Variables:      convertVars(c.Variables, p.CliVars),
		MultyResources: c.MultyResources,
	}
	if c.MultyConfig != nil {
		result.GlobalConfig = c.MultyConfig.HCLBody
	}
	return result

}

func mergeConfigs(configs []config) config {
	var c config
	for _, fileConfig := range configs {
		c.MultyResources = append(c.MultyResources, fileConfig.MultyResources...)
		c.Variables = append(c.Variables, fileConfig.Variables...)
		if fileConfig.MultyConfig != nil {
			if c.MultyConfig != nil {
				log.Fatalf("multiple multy configs found, but only one is allowed")
			}
			c.MultyConfig = fileConfig.MultyConfig
		}
	}
	return c
}

func (p *Parser) parseSingleFile(filepath string, parser *hclparse.Parser, c *config) hcl.Blocks {
	f, diags := parser.ParseHCLFile(filepath)
	if diags != nil {
		validate.LogFatalWithDiags(diags, "unable to parse file %s", filepath)
	}
	diags = gohcl.DecodeBody(f.Body, nil, c)
	if diags != nil {
		validate.LogFatalWithDiags(diags, "failed to decode file %s", filepath)
	}

	rootSchema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "multy", LabelNames: []string{"type", "id"}},
		},
	}

	content, _, diags := f.Body.PartialContent(rootSchema)
	if diags != nil {
		validate.LogFatalWithDiags(diags, "Failed to load configuration.")
	}

	return content.Blocks
}

func findDependencies(resource *MultyResource, resourcesById map[string]*MultyResource) []MultyResourceDependency {
	var result []MultyResourceDependency
	attrs, diags := resource.HCLBody.JustAttributes()
	if diags != nil {
		validate.LogFatalWithDiags(diags, "unable to get attributes %s")
	}
	for _, attr := range attrs {
		for _, v := range attr.Expr.Variables() {
			if _, ok := common.AsCloudProvider(v.RootName()); ok {
				if traverseAttr, ok := v[1].(hcl.TraverseAttr); ok {
					if _, ok := resourcesById[traverseAttr.Name]; ok {
						result = append(result, MultyResourceDependency{
							From:        resource,
							To:          resourcesById[traverseAttr.Name],
							SourceRange: v.SourceRange(),
						})
					}
				} else {
					validate.LogFatalWithSourceRange(v.SourceRange(), "expected attr lookup when referencing a cloud name (aws, az, ...)")
				}
			} else {
				if _, ok := resourcesById[v.RootName()]; ok {
					result = append(result, MultyResourceDependency{
						From:        resource,
						To:          resourcesById[v.RootName()],
						SourceRange: v.SourceRange(),
					})
				}
			}
		}
	}
	return result
}

func convertVars(allVars []variables.Variable, cliVars variables.CommandLineVariables) []ParsedVariable {
	cliVarsByName := map[string]string{}
	for _, cliVar := range cliVars {
		cliVarsByName[cliVar.Name] = cliVar.Value
	}
	var result []ParsedVariable
	for _, v := range allVars {
		var val cty.Value
		if cliVar, ok := cliVarsByName[v.Name]; ok {
			val = v.GetValue(&cliVar)
		}
		val = v.GetValue(nil)
		result = append(result, ParsedVariable{
			Name:  v.Name,
			Value: val,
		})
	}
	return result
}
