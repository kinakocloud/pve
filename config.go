package pve

import (
	"strconv"

	"github.com/kinakocloud/pve/response"
)

// Get
func (m *Machine) GetVmConfig(vmId int64) (*response.VMConfig, error) {
	var data response.VMConfig
	if err := m.getQueryJSON(
		"/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/config",
		nil,
		&data,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

// Edit
func (m *Machine) EditVmConfig(vmId int64, config map[string]any) error {
	if err := m.PostFormJSON(
		"/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/config",
		config,
		nil,
	); err != nil {
		return err
	}

	return nil
}
