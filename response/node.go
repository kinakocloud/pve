package response

type NodeStatus struct {
	Data NodeStatusData `json:"data"`
}

type NodeStatusData struct {
	CPUInfo       CPUInfo       `json:"cpuinfo"`
	BootInfo      BootInfo      `json:"boot-info"`
	CPU           float64       `json:"cpu"`
	CurrentKernel CurrentKernel `json:"current-kernel"`
	PVEVersion    string        `json:"pveversion"`
	Swap          MemoryInfo    `json:"swap"`
	Wait          float64       `json:"wait"`
	Uptime        int64         `json:"uptime"`
	KVersion      string        `json:"kversion"`
	RootFS        DiskInfo      `json:"rootfs"`
	KSM           KSMInfo       `json:"ksm"`
	Memory        MemoryInfo    `json:"memory"`
	Idle          float64       `json:"idle"`
	LoadAvg       []string      `json:"loadavg"`
}

type CPUInfo struct {
	Cores   int    `json:"cores"`
	CPUs    int    `json:"cpus"`
	UserHz  int    `json:"user_hz"`
	Flags   string `json:"flags"`
	HVM     string `json:"hvm"`
	Sockets int    `json:"sockets"`
	MHz     string `json:"mhz"`
	Model   string `json:"model"`
}

type BootInfo struct {
	Mode       string `json:"mode"`
	SecureBoot int    `json:"secureboot"`
}

type CurrentKernel struct {
	Version string `json:"version"`
	Release string `json:"release"`
	Machine string `json:"machine"`
	Sysname string `json:"sysname"`
}

type MemoryInfo struct {
	Total int64 `json:"total"`
	Free  int64 `json:"free"`
	Used  int64 `json:"used"`
}

type DiskInfo struct {
	Available int64 `json:"avail"`
	Used      int64 `json:"used"`
	Total     int64 `json:"total"`
	Free      int64 `json:"free"`
}

type KSMInfo struct {
	Shared int64 `json:"shared"`
}
