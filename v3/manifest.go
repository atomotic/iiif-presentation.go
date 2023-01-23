package v3

import "time"

type Manifest struct {
	Context           string            `json:"@context"`
	ID                string            `json:"id"`
	Type              string            `json:"type"`
	Label             Label             `json:"label"`
	Metadata          []Metadata        `json:"metadata"`
	Summary           Summary           `json:"summary"`
	Thumbnail         []Thumbnail       `json:"thumbnail"`
	ViewingDirection  string            `json:"viewingDirection"`
	Behavior          []string          `json:"behavior"`
	NavDate           time.Time         `json:"navDate"`
	Rights            string            `json:"rights"`
	RequiredStatement RequiredStatement `json:"requiredStatement"`
	Provider          []Provider        `json:"provider"`
	Homepage          []Homepage        `json:"homepage"`
	Service           []Service         `json:"service"`
	SeeAlso           []SeeAlso         `json:"seeAlso"`
	Rendering         []Rendering       `json:"rendering"`
	PartOf            []PartOf          `json:"partOf"`
	Start             Start             `json:"start"`
	Services          []Services        `json:"services"`
	Items             []Canvas          `json:"items"`
	Structures        []Structure       `json:"structures"`
	Annotations       []Annotation      `json:"annotations"`
}
