package v3

type Thumbnail struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Format  string `json:"format"`
	Service []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Profile string `json:"profile"`
	} `json:"service"`
}
