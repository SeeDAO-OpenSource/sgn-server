package erc721

import (
	"encoding/json"
)

type TokenInfo struct {
	ID        string        `json:"id" bson:"_id"`
	Contract  string        `json:"contract" bson:"contract"`
	Name      string        `json:"name" bson:"name"`
	TokenId   int64         `json:"token_id" bson:"token_id"`
	TokenURI  string        `json:"token_uri" bson:"token_uri"`
	Metadata  TokenMetadata `json:"metadata" bson:"metadata"`
	TimeStamp int64         `json:"timestamp" bson:"timestamp"`
	Owner     string        `json:"owner" bson:"owner"`
}

type TokenMetadata struct {
	Attributes []AttrItem `json:"attributes" bson:"attributes"`
	Image      string     `json:"image" bson:"image"`
}

type AttrItem struct {
	TraitType string `json:"trait_type" bson:"trait_type"`
	Value     string `json:"value" bson:"value"`
}

func ParseMetadata(metadata string) TokenMetadata {
	m := TokenMetadata{}
	json.Unmarshal([]byte(metadata), &m)
	return m
}
