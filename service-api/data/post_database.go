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

//DatabasePostStore store and retrive Post data in database
type DatabasePostStore struct {
	dbDatabase *mongo.Database
	collection *mongo.Collection
	timeout    time.Duration
}

//NewDatabasePostStore return a new DatabasePostStore
func NewDatabasePostStore(dnname string) *DatabasePostStore {

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT_DB"))
	if err != nil {
		log.Fatalln("TIMEOUT_DB is incorrect:", err)
	}

	db := database.GetClient().Database(dnname)

	return &DatabasePostStore{
		dbDatabase: db,
		collection: db.Collection("posts"),
		timeout:    time.Duration(timeout) * time.Millisecond,
	}
}

//Register save Post data to the store
func (store *DatabasePostStore) Register(ctx context.Context, post *model.Post) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	post.ID = primitive.NewObjectID()
	post.Likes = []primitive.ObjectID{}
	insertResult, err := store.collection.InsertOne(ctx, post)
	if err != nil {
		return "", errors.New("Error inserting new post: " + err.Error())
	}

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

//List ...
func (store *DatabasePostStore) List(ctx context.Context, memberID string, page int64) (*[]model.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return nil, errors.New("Error get post list: " + err.Error())
	}

	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "members"},
			{"localField", model.PostMongoType().MemberID},
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

	lookupFollowerStage := bson.D{{
		"$lookup", bson.D{
			{"from", "following"},
			{"localField", model.PostMongoType().MemberID},
			{"foreignField", model.FollowingMongoType().FollowingID},
			{"as", "followerslist"},
		},
	}}

	matchStage := bson.D{{
		"$match", bson.M{
			"$or": bson.A{
				bson.D{{model.PostMongoType().MemberID, objectMemberID}},
				bson.D{{"followerslist", bson.D{{
					"$elemMatch", bson.D{{
						model.FollowingMongoType().FollowerID, bson.D{{
							"$eq", objectMemberID,
						}},
					}},
				}},
				}},
			},
		},
	}}

	limitStage := bson.D{{"$limit", 21}}
	skipStage := bson.D{{"$skip", (page - 1) * 20}}
	sortStage := bson.D{{
		"$sort", bson.D{{model.PostMongoType().RegisterDate, -1}},
	}}

	pipeline := mongo.Pipeline{lookupStage, lookupFollowerStage, matchStage, unwindStage, limitStage, skipStage, sortStage}

	// Aggregate
	cur, err := store.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	// Here's an array in which you can store the decoded documents
	var results []model.Post

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

//MyPosts ...
func (store *DatabasePostStore) MyPosts(ctx context.Context, memberID string, page int64) (*[]model.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return nil, errors.New("Error get post list: " + err.Error())
	}

	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "members"},
			{"localField", model.PostMongoType().MemberID},
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
			{model.PostMongoType().MemberID, objectMemberID},
		},
	}}

	limitStage := bson.D{{"$limit", 20}}
	skipStage := bson.D{{"$skip", (page - 1) * 20}}
	sortStage := bson.D{{
		"$sort", bson.D{{model.PostMongoType().RegisterDate, -1}},
	}}

	pipeline := mongo.Pipeline{lookupStage, matchStage, unwindStage, limitStage, skipStage, sortStage}

	// Aggregate
	cur, err := store.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	// Here's an array in which you can store the decoded documents
	var results []model.Post

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

//Search ...
func (store *DatabasePostStore) Search(ctx context.Context, term string, page int64) (*[]model.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "members"},
			{"localField", model.PostMongoType().MemberID},
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
			// {model.PostMongoType().MemberID, objectMemberID},
			{model.PostMongoType().Content, primitive.Regex{Pattern: ".*" + term + ".*", Options: "gi"}},
		},
	}}

	limitStage := bson.D{{"$limit", 21}}
	skipStage := bson.D{{"$skip", (page - 1) * 20}}
	sortStage := bson.D{{
		"$sort", bson.D{{model.PostMongoType().RegisterDate, -1}},
	}}

	pipeline := mongo.Pipeline{lookupStage, matchStage, unwindStage, limitStage, skipStage, sortStage}

	// Aggregate
	cur, err := store.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	// Here's an array in which you can store the decoded documents
	var results []model.Post

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

//GetLike ...
func (store *DatabasePostStore) GetLike(ctx context.Context, memberID string, postID string) (bool, int, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectPostID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return false, 0, errors.New("Error get post like: " + err.Error())
	}

	findOptions := options.FindOne()
	findOptions.SetProjection(bson.M{
		model.PostMongoType().Likes: 1,
	})

	filter := bson.M{
		model.PostMongoType().ID: objectPostID,
	}

	post := model.Post{}

	// FindOne
	err = store.collection.FindOne(ctx, filter, findOptions).Decode(&post)
	if err != nil {
		return false, 0, err
	}

	for _, item := range post.Likes {
		if item.Hex() == memberID {
			return true, len(post.Likes), nil
		}
	}

	return false, len(post.Likes), nil
}

//EditLike ...
func (store *DatabasePostStore) EditLike(ctx context.Context, memberID string, postID string, like bool) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	oldLike, countLike, err := store.GetLike(ctx, memberID, postID)
	if err != nil {
		return false, err
	}
	if oldLike == like {
		return like, nil
	}

	objectPostID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return false, errors.New("Error get post like: " + err.Error())
	}

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return false, errors.New("Error edit post like: " + err.Error())
	}

	filter := bson.M{
		model.PostMongoType().ID: objectPostID,
	}

	update := bson.M{}
	if like {
		if countLike == 0 {
			update = bson.M{
				"$set": bson.M{
					model.PostMongoType().CountLikes: 1,
					model.PostMongoType().Likes:      []primitive.ObjectID{objectMemberID},
				},
			}
		} else {
			update = bson.M{
				"$push": bson.M{
					model.PostMongoType().Likes: objectMemberID,
				},
				"$set": bson.M{
					model.PostMongoType().CountLikes: countLike + 1,
				},
			}
		}
	} else {
		update = bson.M{
			"$pull": bson.M{
				model.PostMongoType().Likes: objectMemberID,
			},
			"$set": bson.M{
				model.PostMongoType().CountLikes: countLike - 1,
			},
		}
	}

	_, err = store.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, errors.New("Error edit post like: " + err.Error())
	}

	return like, nil
}
