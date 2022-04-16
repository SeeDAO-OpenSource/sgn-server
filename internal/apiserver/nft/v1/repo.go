package nftv1

import (
	"context"

	"github.com/waite-lee/nftserver/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Erc721TransferLogRepo interface {
	InsertMany(data []ERC721Transfer) error
	GetAddresses() ([]string, error)
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
	collection := r.collection()
	insertData := make([]interface{}, len(data))
	for index, v := range data {
		insertData[index] = v
	}
	_, err := collection.InsertMany(context.Background(), insertData, options.InsertMany().SetOrdered(false))
	return mongodb.CheckInsertError(err)

}

func (r *MongoErc721TransferLogRepo) GetAddresses() ([]string, error) {
	collection := r.collection()
	context := context.TODO()
	data, err := collection.Distinct(context, "contract_address", bson.D{})
	if err != nil {
		return nil, err
	}
	result := make([]string, len(data))
	for index, v := range data {
		result[index] = v.(string)
	}
	return result, err
}

func (r *MongoErc721TransferLogRepo) collection() *mongo.Collection {
	return r.MongoClient.Database("nft").Collection("transfer_log")
}
