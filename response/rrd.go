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
