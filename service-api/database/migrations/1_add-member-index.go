package migrations

import (
	"context"
	"log"

	"github.com/EsmaeilMazahery/wild/model"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	log.Printf("init migration : 1_add-member-index")
	migrate.Register(func(db *mongo.Database) error {
		opt := options.Index().SetName(model.MemberMongoType().Mobile).SetUnique(true)
		keys := bson.D{{Key: model.MemberMongoType().Mobile, Value: 1}}
		m := mongo.IndexModel{Keys: keys, Options: opt}
		_, err := db.Collection("members").Indexes().CreateOne(context.TODO(), m)
		if err != nil {
			return err
		}

		opt = options.Index().SetName(model.MemberMongoType().Email).SetUnique(true)
		opt = opt.SetPartialFilterExpression(bson.M{model.MemberMongoType().Email: bson.M{"$exists": true, "$gt": ""}})
		keys = bson.D{{Key: model.MemberMongoType().Email, Value: 1}}
		m = mongo.IndexModel{Keys: keys, Options: opt}
		_, err = db.Collection("members").Indexes().CreateOne(context.TODO(), m)
		if err != nil {
			return err
		}

		opt = options.Index().SetName(model.MemberMongoType().Username).SetUnique(true)
		keys = bson.D{{Key: model.MemberMongoType().Username, Value: 1}}
		m = mongo.IndexModel{Keys: keys, Options: opt}
		_, err = db.Collection("members").Indexes().CreateOne(context.TODO(), m)
		if err != nil {
			return err
		}

		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection("members").Indexes().DropOne(context.TODO(), model.MemberMongoType().Mobile)
		if err != nil {
			return err
		}

		_, err = db.Collection("members").Indexes().DropOne(context.TODO(), model.MemberMongoType().Email)
		if err != nil {
			return err
		}

		_, err = db.Collection("members").Indexes().DropOne(context.TODO(), model.MemberMongoType().Username)
		if err != nil {
			return err
		}

		return nil
	})
}
