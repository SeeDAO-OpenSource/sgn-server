package sgnv1

import (
	"time"

	"gorm.io/datatypes"
)

type Metadata struct {
	ID        int64  `gorm:"primarykey"`
	Uri       string `gorm:"size,512"`
	content   datatypes.JSON
	CreatedAt time.Time
}

type ERC721Transfer struct {
	ID                string    `json:"id" bson:"_id"`
	BlockNumber       int       `json:"block_number" bson:"block_number"`
	TimeStamp         time.Time `json:"time_tamp" bson:"time_tamp"`
	Hash              string    `gorm:"szie,80" json:"hash" bson:"hash"`
	Nonce             int
	BlockHash         string `gorm:"szie,80" json:"block_hash" bson:"block_hash"`
	From              string `gorm:"szie,80"`
	ContractAddress   string `gorm:"szie,80" json:"contract_address" bson:"contract_address"`
	To                string `gorm:"szie,80"`
	TokenID           int64  `json:"token_id" bson:"token_id"`
	TokenName         string `gorm:"szie,255" json:"token_name" bson:"token_name"`
	TokenSymbol       string `gorm:"szie,255" json:"token_symbol" bson:"token_symbol"`
	TokenDecimal      uint8  `json:"token_decimal" bson:"token_decimal"`
	TransactionIndex  int    `json:"transaction_index" bson:"transaction_index"`
	Gas               int    `json:"gas" bson:"gas"`
	GasPrice          int64  `json:"gas_price" bson:"gas_price"`
	GasUsed           int    `json:"gas_used" bson:"gas_used"`
	CumulativeGasUsed int    `json:"cumulative_gas_used" bson:"cumulative_gas_used"`
	Input             string `gorm:"szie,255"`
	Confirmations     int    `json:"confirmations" bson:"confirmations"`
}
