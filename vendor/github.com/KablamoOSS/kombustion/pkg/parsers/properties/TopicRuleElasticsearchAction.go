package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TopicRuleElasticsearchAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-elasticsearchaction.html
type TopicRuleElasticsearchAction struct {
	Endpoint interface{} `yaml:"Endpoint"`
	Id       interface{} `yaml:"Id"`
	Index    interface{} `yaml:"Index"`
	RoleArn  interface{} `yaml:"RoleArn"`
	Type     interface{} `yaml:"Type"`
}

// TopicRuleElasticsearchAction validation
func (resource TopicRuleElasticsearchAction) Validate() []error {
	errs := []error{}

	if resource.Endpoint == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Endpoint'"))
	}
	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	if resource.Index == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Index'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
