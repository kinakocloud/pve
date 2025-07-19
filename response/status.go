package response

type Time struct {
	Data struct {
		Time      int64  `json:"time"`
		LocalTime int64  `json:"localtime"`
		TimeZone  string `json:"timezone"`
	} `json:"data"`
}

type VMStatusData struct {
	Clipboard      any                   `json:"clipboard"`
	Disk           int64                 `json:"disk"`
	Mem            int64                 `json:"mem"`
	Serial         int                   `json:"serial"`
	BlockStat      map[string]BlockStats `json:"blockstat"`
	QMPStatus      string                `json:"qmpstatus"`
	PID            int                   `json:"pid"`
	FreeMem        int64                 `json:"freemem"`
	CPU            float64               `json:"cpu"`
	Balloon        int64                 `json:"balloon"`
	Name           string                `json:"name"`
	BalloonInfo    BalloonInfo           `json:"ballooninfo"`
	Uptime         int64                 `json:"uptime"`
	ProxmoxSupport ProxmoxSupport        `json:"proxmox-support"`
	RunningMachine string                `json:"running-machine"`
	Agent          int                   `json:"agent"`
	NetIn          uint64                `json:"netin"`
	NetOut         uint64                `json:"netout"`
	Status         string                `json:"status"`
	DiskWrite      int64                 `json:"diskwrite"`
	CPUs           int                   `json:"cpus"`
	NICs           map[string]NICStats   `json:"nics"`
	DiskRead       int64                 `json:"diskread"`
	MaxDisk        int64                 `json:"maxdisk"`
	MaxMem         int64                 `json:"maxmem"`
	RunningQemu    string                `json:"running-qemu"`
	HA             HAInfo                `json:"ha"`
	VMID           int                   `json:"vmid"`
}
type VMStatus struct {
	Data VMStatusData `json:"data"`
}

type BlockStats struct {
	InvalidFlushOperations      int   `json:"invalid_flush_operations"`
	AccountFailed               bool  `json:"account_failed"`
	UnmapMerged                 int   `json:"unmap_merged"`
	InvalidRdOperations         int   `json:"invalid_rd_operations"`
	InvalidWrOperations         int   `json:"invalid_wr_operations"`
	UnmapBytes                  int64 `json:"unmap_bytes"`
	ZoneAppendOperations        int   `json:"zone_append_operations"`
	FailedUnmapOperations       int   `json:"failed_unmap_operations"`
	FailedWrOperations          int   `json:"failed_wr_operations"`
	AccountInvalid              bool  `json:"account_invalid"`
	UnmapTotalTimeNs            int64 `json:"unmap_total_time_ns"`
	FlushOperations             int   `json:"flush_operations"`
	ZoneAppendMerged            int   `json:"zone_append_merged"`
	ZoneAppendTotalTimeNs       int64 `json:"zone_append_total_time_ns"`
	WrMerged                    int   `json:"wr_merged"`
	RdTotalTimeNs               int64 `json:"rd_total_time_ns"`
	WrOperations                int   `json:"wr_operations"`
	RdMerged                    int   `json:"rd_merged"`
	RdOperations                int   `json:"rd_operations"`
	WrTotalTimeNs               int64 `json:"wr_total_time_ns"`
	InvalidUnmapOperations      int   `json:"invalid_unmap_operations"`
	FailedRdOperations          int   `json:"failed_rd_operations"`
	ZoneAppendBytes             int64 `json:"zone_append_bytes"`
	FlushTotalTimeNs            int64 `json:"flush_total_time_ns"`
	IdleTimeNs                  int64 `json:"idle_time_ns"`
	WrHighestOffset             int64 `json:"wr_highest_offset"`
	UnmapOperations             int   `json:"unmap_operations"`
	FailedZoneAppendOperations  int   `json:"failed_zone_append_operations"`
	RdBytes                     int64 `json:"rd_bytes"`
	TimedStats                  []any `json:"timed_stats"`
	FailedFlushOperations       int   `json:"failed_flush_operations"`
	InvalidZoneAppendOperations int   `json:"invalid_zone_append_operations"`
	WrBytes                     int64 `json:"wr_bytes"`
}

type BalloonInfo struct {
	Actual          int64 `json:"actual"`
	TotalMem        int64 `json:"total_mem"`
	MemSwappedOut   int64 `json:"mem_swapped_out"`
	LastUpdate      int64 `json:"last_update"`
	MaxMem          int64 `json:"max_mem"`
	FreeMem         int64 `json:"free_mem"`
	MinorPageFaults int64 `json:"minor_page_faults"`
	MemSwappedIn    int64 `json:"mem_swapped_in"`
	MajorPageFaults int64 `json:"major_page_faults"`
}

type ProxmoxSupport struct {
	QueryBitmapInfo         bool   `json:"query-bitmap-info"`
	BackupMaxWorkers        bool   `json:"backup-max-workers"`
	BackupFleecing          bool   `json:"backup-fleecing"`
	PBSLibraryVersion       string `json:"pbs-library-version"`
	PBSDirtyBitmap          bool   `json:"pbs-dirty-bitmap"`
	PBSDirtyBitmapSavevm    bool   `json:"pbs-dirty-bitmap-savevm"`
	PBSDirtyBitmapMigration bool   `json:"pbs-dirty-bitmap-migration"`
	PBSMasterkey            bool   `json:"pbs-masterkey"`
}

type NICStats struct {
	NetOut int64 `json:"netout"`
	NetIn  int64 `json:"netin"`
}

type HAInfo struct {
	Managed int `json:"managed"`
}

type StatusData struct {
	Data string `json:"data"`
}
