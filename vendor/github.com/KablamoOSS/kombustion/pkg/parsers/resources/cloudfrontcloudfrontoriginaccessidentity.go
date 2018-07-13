package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CloudFrontCloudFrontOriginAccessIdentity Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html
type CloudFrontCloudFrontOriginAccessIdentity struct {
	Type       string                                             `yaml:"Type"`
	Properties CloudFrontCloudFrontOriginAccessIdentityProperties `yaml:"Properties"`
	Condition  interface{}                                        `yaml:"Condition,omitempty"`
	Metadata   interface{}                                        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                        `yaml:"DependsOn,omitempty"`
}

// CloudFrontCloudFrontOriginAccessIdentity Properties
type CloudFrontCloudFrontOriginAccessIdentityProperties struct {
	CloudFrontOriginAccessIdentityConfig *properties.CloudFrontOriginAccessIdentityCloudFrontOriginAccessIdentityConfig `yaml:"CloudFrontOriginAccessIdentityConfig"`
}

// NewCloudFrontCloudFrontOriginAccessIdentity constructor creates a new CloudFrontCloudFrontOriginAccessIdentity
func NewCloudFrontCloudFrontOriginAccessIdentity(properties CloudFrontCloudFrontOriginAccessIdentityProperties, deps ...interface{}) CloudFrontCloudFrontOriginAccessIdentity {
	return CloudFrontCloudFrontOriginAccessIdentity{
		Type:       "AWS::CloudFront::CloudFrontOriginAccessIdentity",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCloudFrontCloudFrontOriginAccessIdentity parses CloudFrontCloudFrontOriginAccessIdentity
func ParseCloudFrontCloudFrontOriginAccessIdentity(name string, data string) (cf types.TemplateObject, err error) {
	var resource CloudFrontCloudFrontOriginAccessIdentity
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CloudFrontCloudFrontOriginAccessIdentity - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseCloudFrontCloudFrontOriginAccessIdentity validator
func (resource CloudFrontCloudFrontOriginAccessIdentity) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCloudFrontCloudFrontOriginAccessIdentityProperties validator
func (resource CloudFrontCloudFrontOriginAccessIdentityProperties) Validate() []error {
	errs := []error{}
	if resource.CloudFrontOriginAccessIdentityConfig == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CloudFrontOriginAccessIdentityConfig'"))
	} else {
		errs = append(errs, resource.CloudFrontOriginAccessIdentityConfig.Validate()...)
	}
	return errs
}