package pve

import (
	"strconv"

	"github.com/kinakocloud/pve/response"
)

func (v *VM) GetRRDData(tf string) ([]response.RRDData, error) {
	var data response.RRDResponse
	if err := v.Machine.getQueryJSON(
		"/nodes/"+v.Machine.NodeId+"/qemu/"+strconv.FormatInt(v.ID, 10)+"/rrddata",
		map[string]any{
			"timeframe": tf,
		},
		&data,
	); err != nil {
		return data.Data, err
	}
	return data.Data, nil
}
