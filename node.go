package pve

import "github.com/kinakocloud/pve/response"

func (m *Machine) GetNodeStatus(nodeId string) (*response.NodeStatus, error) {
	var data response.NodeStatus
	if err := m.getQueryJSON("/nodes/"+nodeId+"/status", nil, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
