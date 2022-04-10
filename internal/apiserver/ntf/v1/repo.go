package nftv1

import (
	"context"
)

type EthlogRepo interface {
	// Get(ctx context.Context, name string) (*Transaction, error)
	// GetList(ctx context.Context) ([]*Transaction, error)
	// Update(ctx context.Context, metadata *Transaction) (*Transaction, error)
	// Create(ctx context.Context, metadata *Transaction) error
	GetERC721TransferLogs(ctx context.Context, address string, page int, pageSize int) ([]ERC721Transfer, error)
	Insert(data []*ERC721Transfer) error
	Clear()
}
