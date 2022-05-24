package sgnv1

import (
	"context"

	"github.com/SeeDAO-OpenSource/sgn/internal/apiserver/pkg/erc721"
	"github.com/SeeDAO-OpenSource/sgn/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SgnTokenRepo interface {
	InsertMany(data []*erc721.TokenInfo) error
	GetList(address string, page int, pageSize int) ([]erc721.TokenInfo, error)
	Get(token int64, address string) (erc721.TokenInfo, error)
}

type MongoDbSgnTokenRepo struct {
	MongoClient *mongo.Client
}

func NewMongoDbSgnTokenRepo(mongoClient *mongo.Client) SgnTokenRepo {
	return &MongoDbSgnTokenRepo{
		MongoClient: mongoClient,
	}
}

func (r *MongoDbSgnTokenRepo) InsertMany(data []*erc721.TokenInfo) error {
	collection := r.tranfersCollection()
	insertData := make([]interface{}, len(data))
	for index, v := range data {
		insertData[index] = v
	}
	_, err := collection.InsertMany(context.Background(), insertData, options.InsertMany().SetOrdered(false))
	return mongodb.CheckInsertError(err)
}

func (r *MongoDbSgnTokenRepo) GetList(address string, page int, pageSize int) ([]erc721.TokenInfo, error) {
	collection := r.tranfersCollection()
	var data []erc721.TokenInfo
	filter := bson.D{{Key: "contract", Value: address}}
	findOptions := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}}).SetLimit(int64(pageSize)).SetSkip(int64((page - 1) * pageSize))
	result, err := collection.Find(context.TODO(), filter, findOptions)
	result.All(context.TODO(), &data)
	return data, err
}

func (r *MongoDbSgnTokenRepo) Get(token int64, address string) (erc721.TokenInfo, error) {
	collection := r.tranfersCollection()
	tokenInfo := erc721.TokenInfo{}
	err := collection.FindOne(context.TODO(), bson.D{{Key: "token_id", Value: token}, {Key: "contract", Value: address}}).Decode(&tokenInfo)
	return tokenInfo, err
}

func (r *MongoDbSgnTokenRepo) tranfersCollection() *mongo.Collection {
	return r.MongoClient.Database("sgn").Collection("token_seedao")

}
