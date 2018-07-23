package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ApiGatewayGatewayResponse Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-gatewayresponse.html
type ApiGatewayGatewayResponse struct {
	Type       string                              `yaml:"Type"`
	Properties ApiGatewayGatewayResponseProperties `yaml:"Properties"`
	Condition  interface{}                         `yaml:"Condition,omitempty"`
	Metadata   interface{}                         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                         `yaml:"DependsOn,omitempty"`
}

// ApiGatewayGatewayResponse Properties
type ApiGatewayGatewayResponseProperties struct {
	ResponseType       interface{} `yaml:"ResponseType"`
	RestApiId          interface{} `yaml:"RestApiId"`
	StatusCode         interface{} `yaml:"StatusCode,omitempty"`
	ResponseParameters interface{} `yaml:"ResponseParameters,omitempty"`
	ResponseTemplates  interface{} `yaml:"ResponseTemplates,omitempty"`
}

// NewApiGatewayGatewayResponse constructor creates a new ApiGatewayGatewayResponse
func NewApiGatewayGatewayResponse(properties ApiGatewayGatewayResponseProperties, deps ...interface{}) ApiGatewayGatewayResponse {
	return ApiGatewayGatewayResponse{
		Type:       "AWS::ApiGateway::GatewayResponse",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseApiGatewayGatewayResponse parses ApiGatewayGatewayResponse
func ParseApiGatewayGatewayResponse(
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
	var resource ApiGatewayGatewayResponse
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

// ParseApiGatewayGatewayResponse validator
func (resource ApiGatewayGatewayResponse) Validate() []error {
	return resource.Properties.Validate()
}

// ParseApiGatewayGatewayResponseProperties validator
func (resource ApiGatewayGatewayResponseProperties) Validate() []error {
	errors := []error{}
	if resource.ResponseType == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'ResponseType'"))
	}
	if resource.RestApiId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'RestApiId'"))
	}
	return errors
}