package v3

type Services struct {
	ID      string    `json:"@id"`
	Type    string    `json:"@type"`
	Profile string    `json:"profile"`
	Label   string    `json:"label"`
	Service []Service `json:"service"`
}
