// +build plugin

package outputs

import (
	"github.com/KablamoOSS/kombustion/kombustion-plugin-kablamo-network/common"
	"github.com/KablamoOSS/kombustion/types"
	yaml "gopkg.in/yaml.v2"
)

func ParseNetworkVPC(name string, data string) (cf types.ValueMap, err error) {
	var config common.NetworkVPCConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// create a group of objects (each to be validated)
	cf = make(types.ValueMap)

	return
}