package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// RuleKinesisParameters Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-kinesisparameters.html
type RuleKinesisParameters struct {
	PartitionKeyPath interface{} `yaml:"PartitionKeyPath"`
}

// RuleKinesisParameters validation
func (resource RuleKinesisParameters) Validate() []error {
	errs := []error{}

	if resource.PartitionKeyPath == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PartitionKeyPath'"))
	}
	return errs
}
