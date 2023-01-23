package v3

type Homepage struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Label struct {
		En []string `json:"en"`
	} `json:"label"`
	Format string `json:"format"`
}