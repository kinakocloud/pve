package response

type OSInfo struct {
	Result struct {
		Name          string `json:"name"`
		KernelVersion string `json:"kernel-version"`
		ID            string `json:"id"`
		PrettyName    string `json:"pretty-name"`
		VersionID     string `json:"version-id"`
		Version       string `json:"version"`
		KernelRelease string `json:"kernel-release"`
		Machine       string `json:"machine"`
	} `json:"result"`
}
