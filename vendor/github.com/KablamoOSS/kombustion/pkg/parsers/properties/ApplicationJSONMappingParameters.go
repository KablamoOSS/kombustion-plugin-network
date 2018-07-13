package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ApplicationJSONMappingParameters Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-jsonmappingparameters.html
type ApplicationJSONMappingParameters struct {
	RecordRowPath interface{} `yaml:"RecordRowPath"`
}

// ApplicationJSONMappingParameters validation
func (resource ApplicationJSONMappingParameters) Validate() []error {
	errs := []error{}

	if resource.RecordRowPath == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RecordRowPath'"))
	}
	return errs
}
