package v3

type RequiredStatement struct {
	Label struct {
		En []string `json:"en"`
	} `json:"label"`
	Value struct {
		En []string `json:"en"`
	} `json:"value"`
}
