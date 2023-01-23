package v3

type Annotation struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Motivation string `json:"motivation"`
	Body       struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Format  string `json:"format"`
		Service []struct {
			ID      string `json:"id"`
			Type    string `json:"type"`
			Profile string `json:"profile"`
			Service []struct {
				ID   string `json:"@id"`
				Type string `json:"@type"`
			} `json:"service"`
		} `json:"service"`
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"body"`
	Target string `json:"target"`
}
