package v3

type Provider struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Label struct {
		En []string `json:"en,omitempty"`
	} `json:"label,omitempty"`
	Homepage []struct {
		ID    string `json:"id,omitempty"`
		Type  string `json:"type,omitempty"`
		Label struct {
			En []string `json:"en,omitempty"`
		} `json:"label,omitempty"`
		Format string `json:"format,omitempty"`
	} `json:"homepage,omitempty"`
	Logo []struct {
		ID      string `json:"id,omitempty"`
		Type    string `json:"type,omitempty"`
		Format  string `json:"format,omitempty"`
		Service []struct {
			ID      string `json:"id,omitempty"`
			Type    string `json:"type,omitempty"`
			Profile string `json:"profile,omitempty"`
		} `json:"service,omitempty"`
	} `json:"logo,omitempty"`
	SeeAlso []struct {
		ID      string `json:"id,omitempty"`
		Type    string `json:"type,omitempty"`
		Format  string `json:"format,omitempty"`
		Profile string `json:"profile,omitempty"`
	} `json:"seeAlso,omitempty"`
}
