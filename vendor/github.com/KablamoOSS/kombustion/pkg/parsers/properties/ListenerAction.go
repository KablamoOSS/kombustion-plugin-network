package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ListenerAction Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-defaultactions.html
type ListenerAction struct {
	TargetGroupArn interface{} `yaml:"TargetGroupArn"`
	Type           interface{} `yaml:"Type"`
}

// ListenerAction validation
func (resource ListenerAction) Validate() []error {
	errs := []error{}

	if resource.TargetGroupArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetGroupArn'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}