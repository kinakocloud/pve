package pve

import (
	"strconv"

	"github.com/kinakocloud/pve/response"
)

func (v *VM) ResetUserPassword(username string, password string) error {
	params := map[string]any{
		"username": username,
		"password": password,
	}

	if err := v.Machine.PostFormJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/agent/set-user-password",
		params, nil,
	); err != nil {
		return err
	}

	return nil
}

func (v *VM) GetOSInfo() (*response.OSInfo, error) {
	var data response.OSInfo
	if err := v.Machine.getQueryJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/agent/osinfo",
		nil,
		&data,
	); err != nil {
		return nil, err
	}

	return &data, nil
}
