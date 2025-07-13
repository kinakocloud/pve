package pve

import (
	"strconv"

	"github.com/kinakocloud/pve/response"
)

func (v *VM) GetVmStatusCurrent() (*response.VMStatus, error) {
	var data response.VMStatus
	if err := v.Machine.getQueryJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/status/current",
		nil,
		&data,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

func (v *VM) VmStart() error {
	var params = map[string]any{
		"timeout": 30, // seconds
	}

	if err := v.Machine.PostFormJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/status/start",
		params,
		nil,
	); err != nil {
		return err
	}
	return nil
}
