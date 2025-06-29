package pve

import (
	"strconv"

	"github.com/kinakocloud/pve/response"
)

func (m *Machine) ResetUserPassword(vmId int64, username string, password string) error {
	params := map[string]any{
		"username": username,
		"password": password,
	}

	if err := m.PostFormJSON("/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/agent/set-user-password", params, nil); err != nil {
		return err
	}

	return nil
}

func (m *Machine) GetOSInfo(vmId int64) (*response.OSInfo, error) {
	var data response.OSInfo
	if err := m.getQueryJSON(
		"/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/agent/osinfo",
		nil,
		&data,
	); err != nil {
		return nil, err
	}

	return &data, nil
}
