package pve

import "strconv"

func (m *Machine) Clone(newId, vmId int64) error {
	params := map[string]any{
		"newid":   newId,
		"vmid":    vmId,
		"node":    m.NodeId,
		"target":  m.NodeId,
		"full":    true,
		"storage": m.Storage,
	}

	if err := m.PostFormJSON("/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/clone", params, nil); err != nil {
		return err
	}

	return nil
}
