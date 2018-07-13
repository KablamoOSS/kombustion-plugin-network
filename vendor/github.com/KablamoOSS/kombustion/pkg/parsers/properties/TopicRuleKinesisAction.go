package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TopicRuleKinesisAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-kinesisaction.html
type TopicRuleKinesisAction struct {
	PartitionKey interface{} `yaml:"PartitionKey,omitempty"`
	RoleArn      interface{} `yaml:"RoleArn"`
	StreamName   interface{} `yaml:"StreamName"`
}

// TopicRuleKinesisAction validation
func (resource TopicRuleKinesisAction) Validate() []error {
	errs := []error{}

	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.StreamName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StreamName'"))
	}
	return errs
}