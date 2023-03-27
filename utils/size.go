package utils

import (
	"context"
	"fmt"

	"github.com/carlmjohnson/requests"
)

type Image struct {
	Context  string `json:"@context"`
	Protocol string `json:"protocol"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Sizes    []struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"sizes"`
	Tiles []struct {
		Width        int   `json:"width"`
		Height       int   `json:"height"`
		ScaleFactors []int `json:"scaleFactors"`
	} `json:"tiles"`
	ID             string   `json:"id"`
	Type           string   `json:"type"`
	Profile        string   `json:"profile"`
	MaxWidth       int      `json:"maxWidth"`
	MaxHeight      int      `json:"maxHeight"`
	ExtraQualities []string `json:"extraQualities"`
	ExtraFeatures  []string `json:"extraFeatures"`
}

func GetSize(imageapi string) ([]int, error) {
	var image Image
	ctx := context.Background()
	err := requests.URL(imageapi).ToJSON(&image).Fetch(ctx)

	if err != nil {
		return nil, fmt.Errorf("error getting image size: %v", err)
	}
	return []int{image.Width, image.Height}, nil
}
