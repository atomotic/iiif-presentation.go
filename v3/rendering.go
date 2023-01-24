package v3

type Rendering struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Label struct {
		En []string `json:"en,omitempty"`
	} `json:"label,omitempty"`
	Format string `json:"format,omitempty"`
}
