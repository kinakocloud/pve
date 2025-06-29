package response

type IPSets struct {
	Data []IPSetData `json:"data"`
}

type IPSetData struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Digest  string `json:"digest"`
}

type IPSetEntry struct {
	Data []IPSetEntryData `json:"data"`
}

type IPSetEntryData struct {
	Digest  string `json:"digest"`
	Comment string `json:"comment"`
	Cidr    string `json:"cidr"`
}
