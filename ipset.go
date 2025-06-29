package pve

import (
	"strconv"

	"github.com/kinakocloud/pve/response"
)

func (m *Machine) GetIpSets(vmId int64) (*response.IPSets, error) {
	var data response.IPSets

	if err := m.getQueryJSON(
		"/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/firewall/ipset",
		nil,
		&data,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *Machine) GetIpSet(vmId int64, setName string) (*response.IPSetEntry, error) {
	var data response.IPSetEntry

	if err := m.getQueryJSON(
		"/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/firewall/ipset/"+setName,
		nil,
		&data,
	); err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *Machine) DeleteIpSetAddress(vmId int64, setName, address string) error {
	if err := m.deleteQuery(
		"/nodes/" + m.NodeId + "/qemu/" + strconv.FormatInt(vmId, 10) + "/firewall/ipset/" + setName + "/" + address,
	); err != nil {
		return err
	}

	return nil
}

func (m *Machine) DeleteIpSet(vmId int64, setName string) error {
	if err := m.deleteQuery(
		"/nodes/" + m.NodeId + "/qemu/" + strconv.FormatInt(vmId, 10) + "/firewall/ipset/" + setName,
	); err != nil {
		return err
	}

	return nil
}

func (m *Machine) CreateIpSet(vmId int64, setName string) error {
	params := map[string]any{
		"name":    setName,
		"comment": "Created by KNKC",
	}

	if err := m.PostFormJSON(
		"/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/firewall/ipset",
		params,
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (m *Machine) AddIpSetAddress(vmId int64, setName, address string) error {
	params := map[string]any{
		"cidr":    address,
		"nomatch": false,
		"comment": "Added by KNKC",
	}

	if err := m.PostFormJSON(
		"/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/firewall/ipset/"+setName,
		params,
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (m *Machine) UpdateFirewallOptions(vmId int64, options map[string]any) error {
	if err := m.PostFormJSON(
		"/nodes/"+m.NodeId+"/qemu/"+strconv.FormatInt(vmId, 10)+"/firewall/options",
		options,
		nil,
	); err != nil {
		return err
	}

	return nil
}
