package pve

var Status = struct {
	Start    string
	Stop     string
	Suspend  string
	Resume   string
	Reboot   string
	Shutdown string
	Reset    string
}{
	Start:    "start",
	Stop:     "stop",
	Suspend:  "suspend",
	Resume:   "resume",
	Reboot:   "reboot",
	Shutdown: "shutdown",
	Reset:    "reset",
}

var TimeFrame = struct {
	Hour  string
	Day   string
	Week  string
	Month string
	Year  string
}{
	Hour:  "hour",
	Day:   "day",
	Week:  "week",
	Month: "month",
	Year:  "year",
}
