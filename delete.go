package pve

import "fmt"

func (vm *VM) Delete() error {
	if err := vm.Machine.deleteQuery(
		fmt.Sprintf("/nodes/%s/qemu/%d", vm.Machine.NodeId, vm.ID),
	); err != nil {
		return fmt.Errorf("failed to delete VM: %w", err)
	}

	return nil
}
