package v3

type Provider struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Label struct {
		En []string `json:"en"`
	} `json:"label"`
	Homepage []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Label struct {
			En []string `json:"en"`
		} `json:"label"`
		Format string `json:"format"`
	} `json:"homepage"`
	Logo []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Format  string `json:"format"`
		Service []struct {
			ID      string `json:"id"`
			Type    string `json:"type"`
			Profile string `json:"profile"`
		} `json:"service"`
	} `json:"logo"`
	SeeAlso []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Format  string `json:"format"`
		Profile string `json:"profile"`
	} `json:"seeAlso"`
}
