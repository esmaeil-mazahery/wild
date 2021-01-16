package data

import (
	"context"
	"errors"
	"time"

	"github.com/EsmaeilMazahery/wild/database"
	"github.com/EsmaeilMazahery/wild/infrastructure/constant"
	"github.com/EsmaeilMazahery/wild/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SmslogStore is an interface to store Service data
type SmslogStore interface {
	//Save save Service data to the store
	AddLog(ctx context.Context, member *model.Member) (string, error)
}

//DatabaseSmslogStore store member data in memory
type DatabaseSmslogStore struct {
	dbDatabase *mongo.Database
}

//NewDatabaseSmslogStore return a new SmslogStore
func NewDatabaseSmslogStore() *DatabaseSmslogStore {
	return &DatabaseSmslogStore{
		dbDatabase: database.GetClient().Database(constant.Dbname()),
	}
}

//AddLog save Member data to the store
func (store *DatabaseSmslogStore) AddLog(ctx context.Context, member *model.Member) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	collection := store.dbDatabase.Collection("members")
	insertResult, err := collection.InsertOne(ctx, member)
	if err != nil {
		return "", errors.New("Error inserting newUser: " + err.Error())
	}

	return insertResult.InsertedID.(primitive.ObjectID).String(), nil
}
