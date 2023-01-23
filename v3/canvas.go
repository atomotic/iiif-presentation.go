package v3

type Canvas struct {
	ID          string           `json:"id"`
	Type        string           `json:"type"`
	Label       Label            `json:"label"`
	Height      int              `json:"height"`
	Width       int              `json:"width"`
	Items       []AnnotationPage `json:"items"`
	Annotations []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"annotations,omitempty"`
}
