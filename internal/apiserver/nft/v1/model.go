package nftv1

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
	ID                string `json:"id" bson:"_id"`
	BlockNumber       int
	TimeStamp         time.Time
	Hash              string `gorm:"szie,80"`
	Nonce             int
	BlockHash         string `gorm:"szie,80"`
	From              string `gorm:"szie,80"`
	ContractAddress   string `gorm:"szie,80"`
	To                string `gorm:"szie,80"`
	TokenID           int64
	TokenName         string `gorm:"szie,255"`
	TokenSymbol       string `gorm:"szie,255"`
	TokenDecimal      uint8
	TransactionIndex  int
	Gas               int
	GasPrice          int64
	GasUsed           int
	CumulativeGasUsed int
	Input             string `gorm:"szie,255"`
	Confirmations     int
}
