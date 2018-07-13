package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// PipelineField Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects-fields.html
type PipelineField struct {
	Key         interface{} `yaml:"Key"`
	RefValue    interface{} `yaml:"RefValue,omitempty"`
	StringValue interface{} `yaml:"StringValue,omitempty"`
}

// PipelineField validation
func (resource PipelineField) Validate() []error {
	errs := []error{}

	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	return errs
}
