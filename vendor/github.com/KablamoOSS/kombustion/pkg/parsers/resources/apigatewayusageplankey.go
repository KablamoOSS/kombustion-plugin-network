package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ApiGatewayUsagePlanKey Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html
type ApiGatewayUsagePlanKey struct {
	Type       string                           `yaml:"Type"`
	Properties ApiGatewayUsagePlanKeyProperties `yaml:"Properties"`
	Condition  interface{}                      `yaml:"Condition,omitempty"`
	Metadata   interface{}                      `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                      `yaml:"DependsOn,omitempty"`
}

// ApiGatewayUsagePlanKey Properties
type ApiGatewayUsagePlanKeyProperties struct {
	KeyId       interface{} `yaml:"KeyId"`
	KeyType     interface{} `yaml:"KeyType"`
	UsagePlanId interface{} `yaml:"UsagePlanId"`
}

// NewApiGatewayUsagePlanKey constructor creates a new ApiGatewayUsagePlanKey
func NewApiGatewayUsagePlanKey(properties ApiGatewayUsagePlanKeyProperties, deps ...interface{}) ApiGatewayUsagePlanKey {
	return ApiGatewayUsagePlanKey{
		Type:       "AWS::ApiGateway::UsagePlanKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayUsagePlanKey parses ApiGatewayUsagePlanKey
func ParseApiGatewayUsagePlanKey(name string, data string) (cf types.TemplateObject, err error) {
	var resource ApiGatewayUsagePlanKey
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayUsagePlanKey - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseApiGatewayUsagePlanKey validator
func (resource ApiGatewayUsagePlanKey) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayUsagePlanKeyProperties validator
func (resource ApiGatewayUsagePlanKeyProperties) Validate() []error {
	errs := []error{}
	if resource.KeyId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeyId'"))
	}
	if resource.KeyType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeyType'"))
	}
	if resource.UsagePlanId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'UsagePlanId'"))
	}
	return errs
}
