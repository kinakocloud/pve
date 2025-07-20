package response

type RRDData struct {
	Time      int64   `json:"time"`
	Mem       float64 `json:"mem"`
	Disk      float64 `json:"disk"`
	Netout    float64 `json:"netout"`
	MaxCpu    int     `json:"maxcpu"`
	DiskRead  float64 `json:"diskread"`
	CPU       float64 `json:"cpu"`
	MaxDisk   int64   `json:"maxdisk"`
	MaxMem    int64   `json:"maxmem"`
	NetIn     float64 `json:"netin"`
	DiskWrite float64 `json:"diskwrite"`
}

type RRDResponse struct {
	Data []RRDData `json:"data"`
}
type NodeRRDData struct {
	Time      int64   `json:"time"`
	NetIn     float64 `json:"netin"`
	NetOut    float64 `json:"netout"`
	Cpu       float64 `json:"cpu"`
	Iowait    float64 `json:"iowait"`
	Loadavg   float64 `json:"loadavg"`
	SwapUsed  float64 `json:"swapused"`
	SwapTotal float64 `json:"swaptotal"`
	MemUsed   float64 `json:"memused"`
	MemTotal  float64 `json:"memtotal"`
	RootUsed  float64 `json:"rootused"`
	RootTotal float64 `json:"roottotal"`
	MaxCpu    int     `json:"maxcpu"`
}

type NodeRRDResponse struct {
	Data []NodeRRDData `json:"data"`
}
