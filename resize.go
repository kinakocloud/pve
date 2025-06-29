package pve

import "strconv"

func (m *Machine) ResizeDisk(vmId int64, diskName string, size string) error {
	params := map[string]any{
		"disk": diskName,
		"size": size, // K
	}

	if err := m.PostFormJSON("/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/resize", params, nil); err != nil {
		return err
	}

	return nil
}
