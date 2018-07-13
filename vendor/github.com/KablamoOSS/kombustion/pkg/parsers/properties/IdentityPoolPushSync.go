package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// IdentityPoolPushSync Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-pushsync.html
type IdentityPoolPushSync struct {
	RoleArn         interface{} `yaml:"RoleArn,omitempty"`
	ApplicationArns interface{} `yaml:"ApplicationArns,omitempty"`
}

// IdentityPoolPushSync validation
func (resource IdentityPoolPushSync) Validate() []error {
	errs := []error{}

	return errs
}
