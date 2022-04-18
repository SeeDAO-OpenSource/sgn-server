package nftv1

import (
	"context"

	"github.com/waite-lee/nftserver/internal/apiserver/pkg/erc721"
	"github.com/waite-lee/nftserver/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NftTokenRepo interface {
	InsertMany(data []*erc721.TokenInfo) error
	GetList(address string, page int, pageSize int) ([]erc721.TokenInfo, error)
}

type MongoDbNftTokenRepo struct {
	MongoClient *mongo.Client
}

func NewMongoDbNftTokenRepo(mongoClient *mongo.Client) NftTokenRepo {
	return &MongoDbNftTokenRepo{
		MongoClient: mongoClient,
	}
}

func (r *MongoDbNftTokenRepo) InsertMany(data []*erc721.TokenInfo) error {
	collection := r.tranfersCollection()
	insertData := make([]interface{}, len(data))
	for index, v := range data {
		insertData[index] = v
	}
	_, err := collection.InsertMany(context.Background(), insertData, options.InsertMany().SetOrdered(false))
	return mongodb.CheckInsertError(err)
}

func (r *MongoDbNftTokenRepo) GetList(address string, page int, pageSize int) ([]erc721.TokenInfo, error) {
	collection := r.tranfersCollection()
	var data []erc721.TokenInfo
	filter := bson.D{{"contract", address}}
	findOptions := options.Find().SetLimit(int64(pageSize)).SetSkip(int64((page - 1) * pageSize))
	result, err := collection.Find(context.TODO(), filter, findOptions)
	result.All(context.TODO(), &data)
	return data, err
}

func (r *MongoDbNftTokenRepo) tranfersCollection() *mongo.Collection {
	return r.MongoClient.Database("nft").Collection("token_seedao")

}
