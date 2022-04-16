package nftv1

import (
	"context"

	"github.com/waite-lee/nftserver/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Erc721TransferLogRepo interface {
	InsertMany(data []ERC721Transfer) error
}

type MongoErc721TransferLogRepo struct {
	MongoClient *mongo.Client
}

func NewMongoErc721TransferLogRepo(mongoClient *mongo.Client) Erc721TransferLogRepo {
	return &MongoErc721TransferLogRepo{
		MongoClient: mongoClient,
	}
}

func (r *MongoErc721TransferLogRepo) InsertMany(data []ERC721Transfer) error {
	collection := r.MongoClient.Database("nft").Collection("transfer_log")
	insertData := make([]interface{}, len(data))
	for index, v := range data {
		insertData[index] = v
	}
	_, err := collection.InsertMany(context.Background(), insertData, options.InsertMany().SetOrdered(false))
	return mongodb.CheckInsertError(err)

}
