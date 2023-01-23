package v3

type Structure struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Label struct {
		En []string `json:"en"`
	} `json:"label"`
	Items []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Label struct {
			En []string `json:"en"`
		} `json:"label"`
		Supplementary struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"supplementary"`
		Items []struct {
			ID       string `json:"id,omitempty"`
			Type     string `json:"type"`
			Source   string `json:"source,omitempty"`
			Selector struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"selector,omitempty"`
		} `json:"items"`
	} `json:"items"`
}
