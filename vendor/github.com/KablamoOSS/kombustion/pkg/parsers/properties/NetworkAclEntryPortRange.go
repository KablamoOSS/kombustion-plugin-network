package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// NetworkAclEntryPortRange Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-portrange.html
type NetworkAclEntryPortRange struct {
	From interface{} `yaml:"From,omitempty"`
	To   interface{} `yaml:"To,omitempty"`
}

// NetworkAclEntryPortRange validation
func (resource NetworkAclEntryPortRange) Validate() []error {
	errs := []error{}

	return errs
}