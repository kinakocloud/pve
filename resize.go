package pve

import "strconv"

func (v *VM) ResizeDisk(diskName string, size string) error {
	params := map[string]any{
		"disk": diskName,
		"size": size, // K
	}

	if err := v.Machine.PostFormJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/resize",
		params, nil,
	); err != nil {
		return err
	}

	return nil
}
