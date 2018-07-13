package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ClassifierGrokClassifier Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-grokclassifier.html
type ClassifierGrokClassifier struct {
	Classification interface{} `yaml:"Classification"`
	CustomPatterns interface{} `yaml:"CustomPatterns,omitempty"`
	GrokPattern    interface{} `yaml:"GrokPattern"`
	Name           interface{} `yaml:"Name,omitempty"`
}

// ClassifierGrokClassifier validation
func (resource ClassifierGrokClassifier) Validate() []error {
	errs := []error{}

	if resource.Classification == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Classification'"))
	}
	if resource.GrokPattern == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'GrokPattern'"))
	}
	return errs
}
