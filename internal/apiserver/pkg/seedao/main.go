package seedao

import "encoding/json"

type SeedaoMetadata struct {
	Attributes []AttrItem
	Image      string
}

type AttrItem struct {
	TraitType string `json:"trait_type"`
	Value     string
}

func ParseMetadata(metadata string) SeedaoMetadata {
	m := SeedaoMetadata{}
	json.Unmarshal([]byte(metadata), &m)
	return m
}
