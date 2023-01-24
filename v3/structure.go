package v3

type Structure struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Label Label  `json:"label,omitempty"`
	Items []struct {
		ID    string `json:"id,omitempty"`
		Type  string `json:"type,omitempty"`
		Label struct {
			En []string `json:"en,omitempty"`
		} `json:"label,omitempty"`
		Supplementary struct {
			ID   string `json:"id,omitempty"`
			Type string `json:"type,omitempty"`
		} `json:"supplementary,omitempty"`
		Items []struct {
			ID       string `json:"id,omitempty"`
			Type     string `json:"type,omitempty"`
			Source   string `json:"source,omitempty"`
			Selector struct {
				Type  string `json:"type,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"selector,omitempty"`
		} `json:"items,omitempty"`
	} `json:"items,omitempty"`
}
