package data

import (
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/EsmaeilMazahery/wild/database"
	"github.com/EsmaeilMazahery/wild/enums"
	"github.com/EsmaeilMazahery/wild/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//DatabaseNotifyStore store and retrive Post data in database
type DatabaseNotifyStore struct {
	dbDatabase        *mongo.Database
	collection        *mongo.Collection
	collectionPost    *mongo.Collection
	collectionComment *mongo.Collection
	timeout           time.Duration
}

//NewDatabaseNotifyStore return a new DatabaseNotifyStore
func NewDatabaseNotifyStore(dnname string) *DatabaseNotifyStore {

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT_DB"))
	if err != nil {
		log.Fatalln("TIMEOUT_DB is incorrect:", err)
	}

	db := database.GetClient().Database(dnname)

	return &DatabaseNotifyStore{
		dbDatabase:        db,
		collection:        db.Collection("notify"),
		collectionPost:    db.Collection("posts"),
		collectionComment: db.Collection("comments"),
		timeout:           time.Duration(timeout) * time.Millisecond,
	}
}

//Add save Notify data to the store
func (store *DatabaseNotifyStore) Add(ctx context.Context, notify *model.Notify) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	if notify.OwnerMemberID.Hex() == notify.TargetMemberID.Hex() {
		return "", nil
	}

	notify.ID = primitive.NewObjectID()
	notify.RegisterDate = time.Now()
	insertResult, err := store.collection.InsertOne(ctx, notify)
	if err != nil {
		return "", errors.New("Error inserting new notify: " + err.Error())
	}

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (store *DatabaseNotifyStore) AddLikePost(ctx context.Context, postID string, targetMemberID string) error {

	objectPostID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return errors.New("id is incorrect: " + err.Error())
	}

	filter := bson.M{
		model.PostMongoType().ID: objectPostID,
	}

	post := model.Post{}
	err = store.collectionPost.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return err
	}

	objectTargetMemberID, err := primitive.ObjectIDFromHex(targetMemberID)
	if err != nil {
		return errors.New("id is incorrect: " + err.Error())
	}

	_, err = store.Add(ctx, &model.Notify{
		OwnerMemberID:  post.MemberID,
		Content:        "مطلب شما رو لایک کرد",
		Type:           enums.NotifyTypeLike,
		Read:           false,
		TargetMemberID: objectTargetMemberID,
	})

	return err

}
func (store *DatabaseNotifyStore) AddLikeComment(ctx context.Context, commentID string, targetMemberID string) error {

	objectCommentID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return errors.New("id is incorrect: " + err.Error())
	}

	filter := bson.M{
		model.CommentMongoType().ID: objectCommentID,
	}

	comment := model.Comment{}
	err = store.collectionComment.FindOne(ctx, filter).Decode(&comment)
	if err != nil {
		return err
	}

	objectTargetMemberID, err := primitive.ObjectIDFromHex(targetMemberID)
	if err != nil {
		return errors.New("id is incorrect: " + err.Error())
	}

	_, err = store.Add(ctx, &model.Notify{
		OwnerMemberID:  comment.MemberID,
		Content:        "مطلب شما رو لایک کرد",
		Type:           enums.NotifyTypeLike,
		Read:           false,
		TargetMemberID: objectTargetMemberID,
	})

	return err

}
func (store *DatabaseNotifyStore) AddFollow(ctx context.Context, followerID string, followingID string) error {
	objectTargetMemberID, err := primitive.ObjectIDFromHex(followerID)
	if err != nil {
		return errors.New("id is incorrect: " + err.Error())
	}

	objectOwnerMemberID, err := primitive.ObjectIDFromHex(followingID)
	if err != nil {
		return errors.New("id is incorrect: " + err.Error())
	}

	_, err = store.Add(ctx, &model.Notify{
		OwnerMemberID:  objectOwnerMemberID,
		Content:        "شما را دنبال کرد",
		Type:           enums.NotifyTypeFollow,
		Read:           false,
		TargetMemberID: objectTargetMemberID,
	})

	return err
}
func (store *DatabaseNotifyStore) AddComment(ctx context.Context, postID string, targetMemberID string, comment *model.Comment) error {
	objectPostID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return errors.New("id is incorrect: " + err.Error())
	}

	filter := bson.M{
		model.PostMongoType().ID: objectPostID,
	}

	post := model.Post{}
	err = store.collectionPost.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return err
	}

	objectTargetMemberID, err := primitive.ObjectIDFromHex(targetMemberID)
	if err != nil {
		return errors.New("id is incorrect: " + err.Error())
	}

	_, err = store.Add(ctx, &model.Notify{
		OwnerMemberID:  post.MemberID,
		Content:        comment.Content,
		Type:           enums.NotifyTypeComment,
		Read:           false,
		TargetMemberID: objectTargetMemberID,
	})

	return err
}

//List ...
func (store *DatabaseNotifyStore) List(ctx context.Context, memberID string, page int64) (*[]model.Notify, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return nil, errors.New("Error get post list: " + err.Error())
	}

	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "members"},
			{"localField", model.NotifyMongoType().TargetMemberID},
			{"foreignField", model.MemberMongoType().ID},
			{"as", "target_member"},
		},
	}}

	unwindStage := bson.D{{
		"$unwind", bson.D{
			{"path", "$target_member"},
			{"preserveNullAndEmptyArrays", false},
		},
	}}

	matchStage := bson.D{{
		"$match", bson.D{
			{model.NotifyMongoType().OwnerMemberID, objectMemberID},
			// {model.NotifyMongoType().Read, true},
		},
	}}

	limitStage := bson.D{{"$limit", 21}}
	skipStage := bson.D{{"$skip", (page - 1) * 20}}
	sortStage := bson.D{{
		"$sort", bson.D{{model.NotifyMongoType().RegisterDate, -1}},
	}}

	pipeline := mongo.Pipeline{lookupStage, matchStage, unwindStage, limitStage, skipStage, sortStage}

	// Aggregate
	cur, err := store.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	// Here's an array in which you can store the decoded documents
	var results []model.Notify

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

//Read set read flag true
func (store *DatabaseNotifyStore) Read(ctx context.Context, memberID string, ids ...string) error {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return errors.New("Error set read: " + err.Error())
	}

	var objectIDs []primitive.ObjectID
	for _, id := range ids {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			continue
		}

		objectIDs = append(objectIDs, objectID)
	}

	_, err = store.collection.UpdateOne(
		ctx,
		bson.M{
			model.NotifyMongoType().ID: bson.M{
				"$in": objectIDs,
			},
			model.NotifyMongoType().OwnerMemberID: objectMemberID,
		},
		bson.M{"$set": bson.M{
			model.NotifyMongoType().Read: true,
		}},
	)

	return err
}
