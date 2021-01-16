package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/EsmaeilMazahery/wild/infrastructure/constant"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// D: A BSON document. This type should be used in situations where order matters, such as MongoDB commands.
// M: An unordered map. It is the same as D, except it does not preserve order.
// A: A BSON array.
// E: A single element inside a D.

//GetClient connection to Database
func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(constant.Dburi())
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

//IsDupError check error is duplicate
func IsDupError(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}

// GetDatabases ...
func GetDatabases(ctx context.Context, client *mongo.Client) ([]string, error) {
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	return databases, nil
}

// ObjectList convert string ids to object ids
func ObjectList(list []string) (*[]primitive.ObjectID, error) {
	var objects []primitive.ObjectID
	for _, item := range list {
		object, err := primitive.ObjectIDFromHex(item)
		if err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}

	return &objects, nil
}
