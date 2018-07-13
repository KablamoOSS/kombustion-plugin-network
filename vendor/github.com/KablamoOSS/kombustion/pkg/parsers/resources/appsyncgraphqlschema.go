package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// AppSyncGraphQLSchema Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-graphqlschema.html
type AppSyncGraphQLSchema struct {
	Type       string                         `yaml:"Type"`
	Properties AppSyncGraphQLSchemaProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// AppSyncGraphQLSchema Properties
type AppSyncGraphQLSchemaProperties struct {
	ApiId                interface{} `yaml:"ApiId"`
	Definition           interface{} `yaml:"Definition,omitempty"`
	DefinitionS3Location interface{} `yaml:"DefinitionS3Location,omitempty"`
}

// NewAppSyncGraphQLSchema constructor creates a new AppSyncGraphQLSchema
func NewAppSyncGraphQLSchema(properties AppSyncGraphQLSchemaProperties, deps ...interface{}) AppSyncGraphQLSchema {
	return AppSyncGraphQLSchema{
		Type:       "AWS::AppSync::GraphQLSchema",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAppSyncGraphQLSchema parses AppSyncGraphQLSchema
func ParseAppSyncGraphQLSchema(name string, data string) (cf types.TemplateObject, err error) {
	var resource AppSyncGraphQLSchema
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AppSyncGraphQLSchema - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseAppSyncGraphQLSchema validator
func (resource AppSyncGraphQLSchema) Validate() []error {
	return resource.Properties.Validate()
}

// ParseAppSyncGraphQLSchemaProperties validator
func (resource AppSyncGraphQLSchemaProperties) Validate() []error {
	errs := []error{}
	if resource.ApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApiId'"))
	}
	return errs
}
