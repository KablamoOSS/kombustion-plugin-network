package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// StreamingDistributionStreamingDistributionConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html
type StreamingDistributionStreamingDistributionConfig struct {
	Comment        interface{}                          `yaml:"Comment"`
	Enabled        interface{}                          `yaml:"Enabled"`
	PriceClass     interface{}                          `yaml:"PriceClass,omitempty"`
	TrustedSigners *StreamingDistributionTrustedSigners `yaml:"TrustedSigners"`
	S3Origin       *StreamingDistributionS3Origin       `yaml:"S3Origin"`
	Logging        *StreamingDistributionLogging        `yaml:"Logging,omitempty"`
	Aliases        interface{}                          `yaml:"Aliases,omitempty"`
}

// StreamingDistributionStreamingDistributionConfig validation
func (resource StreamingDistributionStreamingDistributionConfig) Validate() []error {
	errs := []error{}

	if resource.Comment == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Comment'"))
	}
	if resource.Enabled == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Enabled'"))
	}
	if resource.TrustedSigners == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TrustedSigners'"))
	} else {
		errs = append(errs, resource.TrustedSigners.Validate()...)
	}
	if resource.S3Origin == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3Origin'"))
	} else {
		errs = append(errs, resource.S3Origin.Validate()...)
	}
	return errs
}
