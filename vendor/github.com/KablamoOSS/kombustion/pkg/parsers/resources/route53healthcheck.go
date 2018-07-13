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

// Route53HealthCheck Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
type Route53HealthCheck struct {
	Type       string                       `yaml:"Type"`
	Properties Route53HealthCheckProperties `yaml:"Properties"`
	Condition  interface{}                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                  `yaml:"DependsOn,omitempty"`
}

// Route53HealthCheck Properties
type Route53HealthCheckProperties struct {
	HealthCheckTags   interface{}                              `yaml:"HealthCheckTags,omitempty"`
	HealthCheckConfig *properties.HealthCheckHealthCheckConfig `yaml:"HealthCheckConfig"`
}

// NewRoute53HealthCheck constructor creates a new Route53HealthCheck
func NewRoute53HealthCheck(properties Route53HealthCheckProperties, deps ...interface{}) Route53HealthCheck {
	return Route53HealthCheck{
		Type:       "AWS::Route53::HealthCheck",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseRoute53HealthCheck parses Route53HealthCheck
func ParseRoute53HealthCheck(name string, data string) (cf types.TemplateObject, err error) {
	var resource Route53HealthCheck
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: Route53HealthCheck - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseRoute53HealthCheck validator
func (resource Route53HealthCheck) Validate() []error {
	return resource.Properties.Validate()
}

// ParseRoute53HealthCheckProperties validator
func (resource Route53HealthCheckProperties) Validate() []error {
	errs := []error{}
	if resource.HealthCheckConfig == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HealthCheckConfig'"))
	} else {
		errs = append(errs, resource.HealthCheckConfig.Validate()...)
	}
	return errs
}
