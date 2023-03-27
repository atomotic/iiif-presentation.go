package v3

type Annotation struct {
	ID         string `json:"id,omitempty"`
	Type       string `json:"type,omitempty"`
	Motivation string `json:"motivation,omitempty"`
	Body       Body   `json:"body,omitempty"`
	Target     string `json:"target,omitempty"`
}

type Body struct {
	ID      string    `json:"id,omitempty"`
	Type    string    `json:"type,omitempty"`
	Format  string    `json:"format,omitempty"`
	Service []Service `json:"service,omitempty"`
	Height  int       `json:"height,omitempty"`
	Width   int       `json:"width,omitempty"`
}
