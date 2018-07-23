package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SNSSubscription Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html
type SNSSubscription struct {
	Type       string                    `yaml:"Type"`
	Properties SNSSubscriptionProperties `yaml:"Properties"`
	Condition  interface{}               `yaml:"Condition,omitempty"`
	Metadata   interface{}               `yaml:"Metadata,omitempty"`
	DependsOn  interface{}               `yaml:"DependsOn,omitempty"`
}

// SNSSubscription Properties
type SNSSubscriptionProperties struct {
	Endpoint interface{} `yaml:"Endpoint,omitempty"`
	Protocol interface{} `yaml:"Protocol,omitempty"`
	TopicArn interface{} `yaml:"TopicArn,omitempty"`
}

// NewSNSSubscription constructor creates a new SNSSubscription
func NewSNSSubscription(properties SNSSubscriptionProperties, deps ...interface{}) SNSSubscription {
	return SNSSubscription{
		Type:       "AWS::SNS::Subscription",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSNSSubscription parses SNSSubscription
func ParseSNSSubscription(
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
	var resource SNSSubscription
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

// ParseSNSSubscription validator
func (resource SNSSubscription) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSNSSubscriptionProperties validator
func (resource SNSSubscriptionProperties) Validate() []error {
	errors := []error{}
	return errors
}