package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// MethodIntegration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration.html
type MethodIntegration struct {
	CacheNamespace        interface{} `yaml:"CacheNamespace,omitempty"`
	ContentHandling       interface{} `yaml:"ContentHandling,omitempty"`
	Credentials           interface{} `yaml:"Credentials,omitempty"`
	IntegrationHttpMethod interface{} `yaml:"IntegrationHttpMethod,omitempty"`
	PassthroughBehavior   interface{} `yaml:"PassthroughBehavior,omitempty"`
	Type                  interface{} `yaml:"Type,omitempty"`
	Uri                   interface{} `yaml:"Uri,omitempty"`
	RequestParameters     interface{} `yaml:"RequestParameters,omitempty"`
	RequestTemplates      interface{} `yaml:"RequestTemplates,omitempty"`
	CacheKeyParameters    interface{} `yaml:"CacheKeyParameters,omitempty"`
	IntegrationResponses  interface{} `yaml:"IntegrationResponses,omitempty"`
}

// MethodIntegration validation
func (resource MethodIntegration) Validate() []error {
	errs := []error{}

	return errs
}
