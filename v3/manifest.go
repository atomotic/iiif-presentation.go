package v3

import (
	"encoding/json"
	"fmt"

	"github.com/atomotic/iiif-presentation.go/utils"
)

type Manifest struct {
	Context           string             `json:"@context"`
	ID                string             `json:"id"`
	Type              string             `json:"type"`
	Label             Label              `json:"label,omitempty"`
	Metadata          []Metadata         `json:"metadata,omitempty"`
	Summary           *Summary           `json:"summary,omitempty"`
	Thumbnail         []Thumbnail        `json:"thumbnail,omitempty"`
	ViewingDirection  string             `json:"viewingDirection,omitempty"`
	Behavior          []string           `json:"behavior,omitempty"`
	NavDate           string             `json:"navDate,omitempty"`
	Rights            string             `json:"rights,omitempty"`
	RequiredStatement *RequiredStatement `json:"requiredStatement,omitempty"`
	Provider          []Provider         `json:"provider,omitempty"`
	Homepage          []Homepage         `json:"homepage,omitempty"`
	Service           []Service          `json:"service,omitempty"`
	SeeAlso           []SeeAlso          `json:"seeAlso,omitempty"`
	Rendering         []Rendering        `json:"rendering,omitempty"`
	PartOf            []PartOf           `json:"partOf,omitempty"`
	Start             *Start             `json:"start,omitempty"`
	Services          []Services         `json:"services,omitempty"`
	Items             []Canvas           `json:"items,omitempty"`
	Structures        []Structure        `json:"structures,omitempty"`
	Annotations       []Annotation       `json:"annotations,omitempty"`
}

func (m *Manifest) Serialize() string {
	iiif, _ := json.MarshalIndent(m, " ", " ")
	return string(iiif)
}

func (m *Manifest) NewItem(canvasid string, label string, body string, size []int, imageservice string) error {
	var err error
	if imageservice != "" {
		size, err = utils.GetSize(imageservice)
		if err != nil {
			return err
		}
	}

	c := Canvas{
		ID:     fmt.Sprintf("%s/canvas/%s", m.ID, canvasid),
		Type:   "Canvas",
		Width:  size[0],
		Height: size[1],
		Label:  Label{"it": {label}},
	}

	ap := AnnotationPage{
		ID:   fmt.Sprintf("%s/page/%s", m.ID, canvasid),
		Type: "AnnotationPage",
	}

	a := Annotation{
		ID:         fmt.Sprintf("%s/annotation/%s", m.ID, canvasid),
		Type:       "Annotation",
		Motivation: "painting",
		Target:     fmt.Sprintf("%s/canvas/%s", m.ID, canvasid),
	}

	b := Body{
		ID:     body,
		Type:   "Image",
		Format: "image/jpeg",
		Width:  size[0],
		Height: size[1],
	}

	s := Service{
		// ID: imageservice,
		ID:      body,
		Type:    "ImageService3",
		Profile: "level2",
	}

	b.Service = append(b.Service, s)
	a.Body = b
	ap.Items = append(ap.Items, a)
	c.Items = append(c.Items, ap)
	m.Items = append(m.Items, c)
	return nil
}

func NewManifest(id, base string) (*Manifest, error) {
	manifest := Manifest{
		Context: "http://iiif.io/api/presentation/3/context.json",
		Type:    "Manifest",
		ID:      fmt.Sprintf("%s/%s", base, id),
	}
	return &manifest, nil
}
