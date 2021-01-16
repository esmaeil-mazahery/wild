package data

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/EsmaeilMazahery/wild/database"
	"github.com/EsmaeilMazahery/wild/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DatabaseCommentStore ...
type DatabaseCommentStore struct {
	dbDatabase *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

//NewDatabaseCommentStore ...
func NewDatabaseCommentStore(dnname string) *DatabaseCommentStore {

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT_DB"))
	if err != nil {
		log.Fatalln("TIMEOUT_DB is incorrect:", err)
	}

	db := database.GetClient().Database(dnname)

	return &DatabaseCommentStore{
		dbDatabase: db,
		collection: db.Collection("comments"),
		timeout:    time.Duration(timeout) * time.Millisecond,
	}
}

//Add save Comment data to the store
func (store *DatabaseCommentStore) Add(ctx context.Context, comment *model.Comment) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	comment.ID = primitive.NewObjectID()
	comment.Likes = []primitive.ObjectID{}
	insertResult, err := store.collection.InsertOne(ctx, comment)
	if err != nil {
		return "", errors.New("Error inserting new comment: " + err.Error())
	}

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

//List ...
func (store *DatabaseCommentStore) List(ctx context.Context, postID string, page int64) (*[]model.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectPostID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, errors.New("Error get comment list: " + err.Error())
	}

	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "members"},
			{"localField", model.CommentMongoType().MemberID},
			{"foreignField", model.MemberMongoType().ID},
			{"as", "member"},
		},
	}}

	unwindStage := bson.D{{
		"$unwind", bson.D{
			{"path", "$member"},
			{"preserveNullAndEmptyArrays", false},
		},
	}}

	matchStage := bson.D{{
		"$match", bson.D{
			{model.CommentMongoType().PostID, objectPostID},
		},
	}}

	limitStage := bson.D{{"$limit", 20}}
	skipStage := bson.D{{"$skip", (page - 1) * 20}}
	sortStage := bson.D{{
		"$sort", bson.D{{model.CommentMongoType().RegisterDate, -1}},
	}}

	pipeline := mongo.Pipeline{lookupStage, matchStage, unwindStage, limitStage, skipStage, sortStage}

	// Aggregate
	cur, err := store.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	// Here's an array in which you can store the decoded documents
	var results []model.Comment

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

//MyComments ...
func (store *DatabaseCommentStore) MyComments(ctx context.Context, memberID string, page int64) (*[]model.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return nil, errors.New("Error get comment list: " + err.Error())
	}

	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "members"},
			{"localField", model.CommentMongoType().MemberID},
			{"foreignField", model.MemberMongoType().ID},
			{"as", "member"},
		},
	}}

	unwindStage := bson.D{{
		"$unwind", bson.D{
			{"path", "$member"},
			{"preserveNullAndEmptyArrays", false},
		},
	}}

	matchStage := bson.D{{
		"$match", bson.D{
			{model.CommentMongoType().MemberID, objectMemberID},
		},
	}}

	limitStage := bson.D{{"$limit", 20}}
	skipStage := bson.D{{"$skip", (page - 1) * 20}}
	sortStage := bson.D{{
		"$sort", bson.D{{model.CommentMongoType().RegisterDate, -1}},
	}}

	pipeline := mongo.Pipeline{lookupStage, matchStage, unwindStage, limitStage, skipStage, sortStage}

	// Aggregate
	cur, err := store.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	// Here's an array in which you can store the decoded documents
	var results []model.Comment

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

//GetLike ...
func (store *DatabaseCommentStore) GetLike(ctx context.Context, memberID string, commentID string) (bool, int, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectCommentID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return false, 0, errors.New("Error get comment like: " + err.Error())
	}

	findOptions := options.FindOne()
	findOptions.SetProjection(bson.M{
		model.CommentMongoType().Likes: 1,
	})

	filter := bson.M{
		model.CommentMongoType().ID: objectCommentID,
	}

	comment := model.Comment{}

	// FindOne
	err = store.collection.FindOne(ctx, filter, findOptions).Decode(&comment)
	if err != nil {
		return false, 0, err
	}

	for _, item := range comment.Likes {
		if item.Hex() == memberID {
			return true, len(comment.Likes), nil
		}
	}

	return false, len(comment.Likes), nil
}

//EditLike ...
func (store *DatabaseCommentStore) EditLike(ctx context.Context, memberID string, commentID string, like bool) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	oldLike, countLike, err := store.GetLike(ctx, memberID, commentID)
	if err != nil {
		return false, err
	}
	if oldLike == like {
		return like, nil
	}

	objectCommentID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return false, errors.New("Error get comment like: " + err.Error())
	}

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return false, errors.New("Error edit comment like: " + err.Error())
	}

	filter := bson.M{
		model.CommentMongoType().ID: objectCommentID,
	}

	update := bson.M{}
	if like {
		if countLike == 0 {
			update = bson.M{
				"$set": bson.M{
					model.CommentMongoType().CountLikes: 1,
					model.CommentMongoType().Likes:      []primitive.ObjectID{objectMemberID},
				},
			}
		} else {
			update = bson.M{
				"$push": bson.M{
					model.CommentMongoType().Likes: objectMemberID,
				},
				"$set": bson.M{
					model.CommentMongoType().CountLikes: countLike + 1,
				},
			}
		}
	} else {
		update = bson.M{
			"$pull": bson.M{
				model.CommentMongoType().Likes: objectMemberID,
			},
			"$set": bson.M{
				model.CommentMongoType().CountLikes: countLike - 1,
			},
		}
	}

	_, err = store.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, errors.New("Error edit comment like: " + err.Error())
	}

	return like, nil
}
