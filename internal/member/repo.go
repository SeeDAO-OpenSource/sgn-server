package member

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MemberRepository interface {
	GetList(page int, pageSize int) ([]Member, error)
	GetByAddress(address string) (Member, error)
	GetByAddresses(address []string) ([]Member, error)
	Insert(member *Member) error
	InsertManay(members []Member) error
	Update(member *Member) error
	Delete(address string) error

	Init() error
}

const (
	MemberDBName         = "sgn"
	MemberCollectionName = "members"
)

type mongoMemberRepository struct {
	client *mongo.Client
}

func NewMongoMemberRepository(client *mongo.Client) MemberRepository {
	return &mongoMemberRepository{
		client: client,
	}
}

func (r mongoMemberRepository) GetList(page int, pageSize int) ([]Member, error) {
	collection := r.collection()
	context := context.TODO()
	options := options.Find().
		SetSort(bson.D{{Key: "createdAt", Value: -1}}).
		SetLimit(int64(pageSize)).
		SetSkip(int64((page - 1) * pageSize))
	data, err := collection.Find(context, bson.D{}, options)
	if err != nil {
		return nil, err
	}
	result := make([]Member, 0)
	err = data.All(context, &result)
	return result, err
}

func (r mongoMemberRepository) GetByAddress(address string) (Member, error) {
	collection := r.collection()
	context := context.TODO()
	data := Member{}
	err := collection.FindOne(context, bson.D{{Key: "address", Value: address}}).Decode(&data)
	return data, err
}

func (r mongoMemberRepository) GetByAddresses(address []string) ([]Member, error) {
	collection := r.collection()
	context := context.TODO()
	data, err := collection.Find(context, bson.D{{Key: "address", Value: bson.M{"$in": address}}})
	if err != nil {
		return nil, err
	}
	result := make([]Member, 0)
	err = data.All(context, &result)
	return result, err
}

func (r mongoMemberRepository) Insert(member *Member) error {
	collection := r.collection()
	context := context.TODO()
	member.CreatedAt = time.Now().Unix()
	_, err := collection.InsertOne(context, &member)
	return err
}

func (r mongoMemberRepository) InsertManay(members []Member) error {
	collection := r.collection()
	context := context.TODO()
	options := options.InsertMany().SetOrdered(false)
	data := make([]interface{}, len(members))
	for i, v := range members {
		data[i] = v
	}
	_, err := collection.InsertMany(context, data, options)
	return err
}

func (r mongoMemberRepository) Update(member *Member) error {
	collection := r.collection()
	context := context.TODO()
	member.UpdatedAt = time.Now().Unix()
	_, err := collection.UpdateOne(context, bson.D{{Key: "address", Value: member.Address}}, &member)
	return err
}
func (r mongoMemberRepository) Delete(address string) error {
	collection := r.collection()
	context := context.TODO()
	_, err := collection.DeleteOne(context, bson.D{{Key: "address", Value: address}})
	return err
}

func (r mongoMemberRepository) Init() error {
	collection := r.collection()
	context := context.TODO()
	options := options.Index().SetUnique(true)
	_, err := collection.Indexes().CreateOne(context, mongo.IndexModel{Keys: bson.D{{Key: "address", Value: 1}}, Options: options})
	return err
}

func (r mongoMemberRepository) collection() *mongo.Collection {
	collection := r.client.Database(MemberDBName).Collection(MemberCollectionName)
	return collection
}
