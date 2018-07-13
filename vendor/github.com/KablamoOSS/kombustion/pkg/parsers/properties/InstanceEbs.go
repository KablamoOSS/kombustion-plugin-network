package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// InstanceEbs Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-blockdev-template.html
type InstanceEbs struct {
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	Encrypted           interface{} `yaml:"Encrypted,omitempty"`
	Iops                interface{} `yaml:"Iops,omitempty"`
	SnapshotId          interface{} `yaml:"SnapshotId,omitempty"`
	VolumeSize          interface{} `yaml:"VolumeSize,omitempty"`
	VolumeType          interface{} `yaml:"VolumeType,omitempty"`
}

// InstanceEbs validation
func (resource InstanceEbs) Validate() []error {
	errs := []error{}

	return errs
}
