package pve

import "strconv"

// ResizeDisk resizes the disk of the VM to the specified size.
// The disk name must be specified, and the size should be in a format
// recognized by Proxmox (e.g., "10G" for 10 gigabytes).
func (v *VM) ResizeDisk(diskName string, size string) error {
	params := map[string]any{
		"disk": diskName,
		"size": size,
	}

	if err := v.Machine.PutFormJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/resize",
		params, nil,
	); err != nil {
		return err
	}

	return nil
}
