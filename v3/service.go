package v3

type Service struct {
	ID      string `json:"@id"`
	Type    string `json:"@type"`
	Profile string `json:"profile"`
}
