package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/db/mongodb"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"go.mongodb.org/mongo-driver/mongo"
)

var mongoOptions = &mongodb.MongoOptions{
	URL: "mongodb://localhost:27017",
}

func AddMongoClient(ac *app.AppBuilder) {
	ac.BindOptions("Mongo", mongoOptions)
	ac.ConfigureServices(func() error {
		services.AddValue(mongoOptions)
		services.AddTransient(func(c *services.Container) *mongo.Client {
			client, err := mongodb.GetClient(mongoOptions)
			if err != nil {
				return nil
			}
			return client
		})
		return nil
	})
}
