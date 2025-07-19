package pve

import (
	"errors"
	"strconv"

	"github.com/kinakocloud/pve/response"
)

var ErrInvalidVMStatus = errors.New("invalid VM status, allowed values are: start, stop, suspend, shutdown, reset, reboot, resume")

// Machine Time
func (m *Machine) GetTime() (*response.Time, error) {
	var data response.Time
	if err := m.getQueryJSON("/nodes/"+m.NodeId+"/time", nil, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (v *VM) GetStatusCurrent() (*response.VMStatus, error) {
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

// SetStatus sets the status of the VM to one of the allowed values.
//
// Use pve.Status.Start, pve.Status.Stop...
// It sends a request to the Proxmox API to change the VM status.
//
// Returns ErrInvalidVMStatus if the status is not allowed.
func (vm *VM) SetStatus(status string) error {
	var response response.StatusData
	err := vm.Machine.PostFormJSON(
		"/nodes/"+vm.Machine.NodeId+"/qemu/"+strconv.FormatInt(vm.ID, 10)+"/status/"+status,
		map[string]any{},
		&response,
	)

	return err
}
