package v3

type AnnotationPage struct {
	ID    string       `json:"id"`
	Type  string       `json:"type"`
	Items []Annotation `json:"items"`
}
