package erc721

import (
	"encoding/json"
)

type TokenInfo struct {
	ID       string `json:"id" bson:"_id"`
	Contract string
	Name     string
	TokenId  int64         `json:"token_id" bson:"token_id"`
	TokenURI string        `json:"token_uri" bson:"token_uri"`
	Metadata TokenMetadata `json:"metadata" bson:"metadata"`
}

type TokenMetadata struct {
	Attributes []AttrItem
	Image      string
}

type AttrItem struct {
	TraitType string `json:"trait_type" bson:"trait_type"`
	Value     string
}

func ParseMetadata(metadata string) TokenMetadata {
	m := TokenMetadata{}
	json.Unmarshal([]byte(metadata), &m)
	return m
}
