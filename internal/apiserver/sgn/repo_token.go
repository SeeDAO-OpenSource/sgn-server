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
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "contract", Value: address}}}}
	groupStage := bson.D{{Key: "$group", Value: bson.D{{Key: "_id", Value: "$token_id"}, {Key: "token", Value: bson.D{{Key: "$first", Value: "$$ROOT"}}}}}}
	// 将所有token属性变为token数组
	projectStage := bson.D{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: "$token"}}}}
	sortStage := bson.D{{Key: "$sort", Value: bson.D{{Key: "timestamp", Value: -1}}}}
	skipStage := bson.D{{Key: "$skip", Value: int64((page - 1) * pageSize)}}
	limitStage := bson.D{{Key: "$limit", Value: int64(pageSize)}}
	// 查询所有tokenId最新的交易信息
	pipeline := []bson.D{matchStage, groupStage, projectStage, sortStage, skipStage, limitStage}
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &data)
	return data, err
}

func (r *MongoDbSgnTokenRepo) Get(token int64, address string) (erc721.TokenInfo, error) {
	collection := r.tranfersCollection()
	findOptions := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}}).SetLimit(1)
	curror, err := collection.Find(context.TODO(), bson.D{{Key: "token_id", Value: token}, {Key: "contract", Value: address}}, findOptions)
	if err != nil {
		return erc721.TokenInfo{}, err
	}
	tokens := []erc721.TokenInfo{}
	if err := curror.All(context.TODO(), &tokens); err != nil {
		return erc721.TokenInfo{}, err
	}
	if len(tokens) == 0 {
		return erc721.TokenInfo{}, mongo.ErrNoDocuments
	}
	return tokens[0], err
}

func (r *MongoDbSgnTokenRepo) tranfersCollection() *mongo.Collection {
	return r.MongoClient.Database(DatabaseName).Collection(TokenCollectionName)
}
