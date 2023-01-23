package v3

type Metadata struct {
	Label struct {
		En []string `json:"en"`
	} `json:"label"`
	Value struct {
		None []string `json:"@none"`
	} `json:"value,omitempty"`
	// Value0 struct {
	// 	En []string `json:"en"`
	// 	Fr []string `json:"fr"`
	// } `json:"value,omitempty"`
	// Value1 struct {
	// 	En []string `json:"en"`
	// } `json:"value,omitempty"`
}
