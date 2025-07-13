package pve

import (
	"strconv"

	"github.com/kinakocloud/pve/response"
)

// Get
func (v *VM) GetConfig() (*response.VMConfig, error) {
	var data response.VMConfig
	if err := v.Machine.getQueryJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/config",
		nil,
		&data,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

// Edit
func (v *VM) EditConfig(config map[string]any) error {
	if err := v.Machine.PostFormJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/config",
		config,
		nil,
	); err != nil {
		return err
	}

	return nil
}
