package pve

import (
	"strconv"

	"github.com/kinakocloud/pve/response"
)

func (v *VM) GetIpSets() (*response.IPSets, error) {
	var data response.IPSets

	if err := v.Machine.getQueryJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/firewall/ipset",
		nil,
		&data,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

func (v *VM) GetIpSet(setName string) (*response.IPSetEntry, error) {
	var data response.IPSetEntry

	if err := v.Machine.getQueryJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/firewall/ipset/"+setName,
		nil,
		&data,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

func (v *VM) DeleteIpSetAddress(setName, address string) error {
	if err := v.Machine.deleteQuery(
		"/nodes/" + v.Machine.NodeId + "/qemu/" + strconv.FormatInt(v.ID, 10) + "/firewall/ipset/" + setName + "/" + address,
	); err != nil {
		return err
	}

	return nil
}

func (v *VM) DeleteIpSet(setName string) error {
	if err := v.Machine.deleteQuery(
		"/nodes/" + v.Machine.NodeId + "/qemu/" + strconv.FormatInt(v.ID, 10) + "/firewall/ipset/" + setName,
	); err != nil {
		return err
	}

	return nil
}

func (v *VM) CreateIpSet(setName string) error {
	params := map[string]any{
		"name":    setName,
		"comment": "Created by KNKC",
	}

	if err := v.Machine.PostFormJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/firewall/ipset",
		params,
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (v *VM) AddIpSetAddress(setName, address string) error {
	params := map[string]any{
		"cidr":    address,
		"nomatch": false,
		"comment": "Added by KNKC",
	}

	if err := v.Machine.PostFormJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/firewall/ipset/"+setName,
		params,
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (v *VM) UpdateFirewallOptions(options map[string]any) error {
	if err := v.Machine.PutFormJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/firewall/options",
		options,
		nil,
	); err != nil {
		return err
	}

	return nil
}
