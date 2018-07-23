package outputs

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ParseDMSReplicationInstance Documentation http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html
func ParseDMSReplicationInstance(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-outputs"

	var resource, output types.TemplateObject

	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-DMSReplicationInstance-" + name,
				},
			},
		},
	}

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "ReplicationInstancePrivateIpAddresses"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-DMSReplicationInstance-" + name + "-ReplicationInstancePrivateIpAddresses",
			},
		},
	}

	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}

	outputs[name+"ReplicationInstancePrivateIpAddresses"] = output

	output = types.TemplateObject{
		"Description": name + " Object",
		"Value": map[string]interface{}{
			"Fn::GetAtt": []string{name, "ReplicationInstancePublicIpAddresses"},
		},
		"Export": map[string]interface{}{
			"Name": map[string]interface{}{
				"Fn::Sub": "${AWS::StackName}-DMSReplicationInstance-" + name + "-ReplicationInstancePublicIpAddresses",
			},
		},
	}

	if condition, ok := resource["Condition"]; ok {
		output["Condition"] = condition
	}

	outputs[name+"ReplicationInstancePublicIpAddresses"] = output

	return
}