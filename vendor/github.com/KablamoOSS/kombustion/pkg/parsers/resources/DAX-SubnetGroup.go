package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// DAXSubnetGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dax-subnetgroup.html
type DAXSubnetGroup struct {
	Type       string                   `yaml:"Type"`
	Properties DAXSubnetGroupProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// DAXSubnetGroup Properties
type DAXSubnetGroupProperties struct {
	Description     interface{} `yaml:"Description,omitempty"`
	SubnetGroupName interface{} `yaml:"SubnetGroupName,omitempty"`
	SubnetIds       interface{} `yaml:"SubnetIds"`
}

// NewDAXSubnetGroup constructor creates a new DAXSubnetGroup
func NewDAXSubnetGroup(properties DAXSubnetGroupProperties, deps ...interface{}) DAXSubnetGroup {
	return DAXSubnetGroup{
		Type:       "AWS::DAX::SubnetGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDAXSubnetGroup parses DAXSubnetGroup
func ParseDAXSubnetGroup(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource DAXSubnetGroup
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseDAXSubnetGroup validator
func (resource DAXSubnetGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDAXSubnetGroupProperties validator
func (resource DAXSubnetGroupProperties) Validate() []error {
	errors := []error{}
	if resource.SubnetIds == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errors
}