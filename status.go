package pve

import (
	"strconv"

	"github.com/kinakocloud/pve/response"
)

func (m *Machine) GetVmStatusCurrent(vmId int64) (*response.VMStatus, error) {
	var data response.VMStatus
	if err := m.getQueryJSON(
		"/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/status/current",
		nil,
		&data,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *Machine) VmStart(vmId int64) error {
	var params = map[string]any{
		"timeout": 30, // seconds
	}

	if err := m.PostFormJSON(
		"/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/status/start",
		params,
		nil,
	); err != nil {
		return err
	}
	return nil
}
