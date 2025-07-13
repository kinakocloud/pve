package pve

import "strconv"

type Machine struct {
	API     string
	User    string
	Token   string
	NodeId  string
	Storage string
}

func NewMachine(host string, port int, user, token, nodeId, storage string) *Machine {
	return &Machine{
		API:     "https://" + host + ":" + strconv.Itoa(port),
		User:    user,
		Token:   token,
		NodeId:  nodeId,
		Storage: storage,
	}
}

type VM struct {
	ID      int64
	Machine *Machine
}

func (m *Machine) VM(id int64) *VM {
	return &VM{
		ID:      id,
		Machine: m,
	}
}
