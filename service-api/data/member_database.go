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
)

//DatabaseMemberStore store member data in memory
type DatabaseMemberStore struct {
	dbDatabase          *mongo.Database
	collectionMember    *mongo.Collection
	collectionFollowing *mongo.Collection
	timeout             time.Duration
}

//NewDatabaseMemberStore return a new DatabaseMemberStore
func NewDatabaseMemberStore(dnname string) *DatabaseMemberStore {

	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT_DB"))
	if err != nil {
		log.Fatalln("TIMEOUT_DB is incorrect:", err)
	}

	db := database.GetClient().Database(dnname)

	return &DatabaseMemberStore{
		dbDatabase:          db,
		collectionMember:    db.Collection("members"),
		collectionFollowing: db.Collection("following"),
		timeout:             time.Duration(timeout) * time.Millisecond,
	}
}

//Add save Member data to the store
func (store *DatabaseMemberStore) Add(ctx context.Context, member *model.Member) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	member.ID = primitive.NewObjectID()
	insertResult, err := store.collectionMember.InsertOne(ctx, member)
	if err != nil {
		return "", errors.New("Error inserting new Member: " + err.Error())
	}

	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

// Find finds a member by ID
func (store *DatabaseMemberStore) Find(ctx context.Context, id string) (*model.Member, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id is incorrect: " + err.Error())
	}

	filter := bson.M{model.MemberMongoType().ID: objectID}

	member := model.Member{}
	err = store.collectionMember.FindOne(ctx, filter).Decode(&member)
	if err != nil {
		return nil, err
	}

	return &member, nil
}

//FindByUsername ...
func (store *DatabaseMemberStore) FindByUsername(ctx context.Context, username string) (*model.Member, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	filter := bson.M{
		model.MemberMongoType().Username: username,
	}

	member := model.Member{}
	err := store.collectionMember.FindOne(ctx, filter).Decode(&member)
	if err != nil {
		return nil, err
	}

	return &member, nil
}

//FindByMobile ...
func (store *DatabaseMemberStore) FindByMobile(ctx context.Context, mobile string) (*model.Member, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()
	filter := bson.D{{Key: model.MemberMongoType().Mobile, Value: mobile}}

	member := model.Member{}
	err := store.collectionMember.FindOne(ctx, filter).Decode(&member)
	if err != nil {
		return nil, err
	}

	return &member, nil
}

//FindByEmail ...
func (store *DatabaseMemberStore) FindByEmail(ctx context.Context, email string) (*model.Member, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()
	filter := bson.M{model.MemberMongoType().Email: email}

	member := model.Member{}
	err := store.collectionMember.FindOne(ctx, filter).Decode(&member)
	if err != nil {
		return nil, err
	}

	return &member, nil
}

//ChangePassword change password of member new password must be hashed
func (store *DatabaseMemberStore) ChangePassword(ctx context.Context, id string, NewPassword string) error {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()
	ID, _ := primitive.ObjectIDFromHex(id)

	_, err := store.collectionMember.UpdateOne(
		ctx,
		bson.M{model.MemberMongoType().ID: ID},
		bson.D{{
			Key:   "$set",
			Value: bson.M{model.MemberMongoType().Password: NewPassword},
		}},
	)

	return err
}

//ChangeImageProfile change image of member
func (store *DatabaseMemberStore) ChangeImageProfile(ctx context.Context, id string, newImage string) error {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()
	oid, _ := primitive.ObjectIDFromHex(id)

	_, err := store.collectionMember.UpdateOne(
		ctx,
		bson.M{model.MemberMongoType().ID: oid},
		bson.M{"$set": bson.M{model.MemberMongoType().Image: newImage}},
	)

	return err
}

//ChangeImageHeader change image header of member
func (store *DatabaseMemberStore) ChangeImageHeader(ctx context.Context, id string, newImage string) error {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()
	oid, _ := primitive.ObjectIDFromHex(id)

	_, err := store.collectionMember.UpdateOne(
		ctx,
		bson.M{model.MemberMongoType().ID: oid},
		bson.M{"$set": bson.M{model.MemberMongoType().ImageHeader: newImage}},
	)

	return err
}

//Edit change token firebase of member
func (store *DatabaseMemberStore) Edit(ctx context.Context, member *model.Member) error {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	_, err := store.collectionMember.UpdateOne(
		ctx,
		bson.M{model.MemberMongoType().ID: member.ID},
		bson.M{"$set": bson.M{
			model.MemberMongoType().Username: member.Username,
			model.MemberMongoType().Name:     member.Name,
			model.MemberMongoType().Family:   member.Family,
			model.MemberMongoType().Email:    member.Email,
			model.MemberMongoType().Image:    member.Image,
			model.MemberMongoType().Mobile:   member.Mobile,
			model.MemberMongoType().Password: member.Password,
		}},
	)

	return err
}

func (store *DatabaseMemberStore) Suggestion(ctx context.Context, memberID string) (*[]model.Member, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return nil, errors.New("Error get comment list: " + err.Error())
	}

	sampleStage := bson.D{{
		"$sample", bson.D{{"size", 10}},
	}}

	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "following"},
			{"localField", model.MemberMongoType().ID},
			{"foreignField", model.FollowingMongoType().FollowingID},
			{"as", "followinglist"},
		},
	}}

	matchStage := bson.D{{
		"$match", bson.D{
			{"followinglist", bson.D{{
				"$not", bson.D{{
					"$elemMatch", bson.D{{
						model.FollowingMongoType().FollowerID, bson.D{{
							"$eq", objectMemberID,
						}},
					}},
				}},
			}},
			},
			{model.MemberMongoType().ID, bson.M{
				"$ne": objectMemberID,
			}},
		},
	}}

	projectStage := bson.D{{
		"$project", bson.D{
			{model.MemberMongoType().ID, 1},
			{model.MemberMongoType().Name, 1},
			{model.MemberMongoType().Family, 1},
			{model.MemberMongoType().Image, 1},
			{model.MemberMongoType().Username, 1},
		},
	}}

	pipeline := mongo.Pipeline{sampleStage, lookupStage, matchStage, projectStage}

	// Aggregate
	cur, err := store.collectionMember.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	// Here's an array in which you can store the decoded documents
	var results []model.Member

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

func (store *DatabaseMemberStore) Followers(ctx context.Context, memberID string) (*[]model.Member, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return nil, errors.New("Error get comment list: " + err.Error())
	}

	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "following"},
			{"localField", model.MemberMongoType().ID},
			{"foreignField", model.FollowingMongoType().FollowerID},
			{"as", "followerlist"},
		},
	}}

	matchStage := bson.D{{
		"$match", bson.D{
			{"followerlist", bson.D{{
				"$elemMatch", bson.D{{
					model.FollowingMongoType().FollowingID, bson.D{{
						"$eq", objectMemberID,
					}},
				}},
			}},
			},
			{model.MemberMongoType().ID, bson.M{
				"$ne": objectMemberID,
			}},
		},
	}}

	projectStage := bson.D{{
		"$project", bson.D{
			{model.MemberMongoType().ID, 1},
			{model.MemberMongoType().Name, 1},
			{model.MemberMongoType().Family, 1},
			{model.MemberMongoType().Image, 1},
			{model.MemberMongoType().Username, 1},
		},
	}}

	pipeline := mongo.Pipeline{lookupStage, matchStage, projectStage}

	// Aggregate
	cur, err := store.collectionMember.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	// Here's an array in which you can store the decoded documents
	var results []model.Member

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

func (store *DatabaseMemberStore) Followings(ctx context.Context, memberID string) (*[]model.Member, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return nil, errors.New("Error get comment list: " + err.Error())
	}

	lookupStage := bson.D{{
		"$lookup", bson.D{
			{"from", "following"},
			{"localField", model.MemberMongoType().ID},
			{"foreignField", model.FollowingMongoType().FollowingID},
			{"as", "followinglist"},
		},
	}}

	matchStage := bson.D{{
		"$match", bson.D{
			{"followinglist", bson.D{{
				"$elemMatch", bson.D{{
					model.FollowingMongoType().FollowerID, bson.D{{
						"$eq", objectMemberID,
					}},
				}},
			}},
			},
			{model.MemberMongoType().ID, bson.M{
				"$ne": objectMemberID,
			}},
		},
	}}

	projectStage := bson.D{{
		"$project", bson.D{
			{model.MemberMongoType().ID, 1},
			{model.MemberMongoType().Name, 1},
			{model.MemberMongoType().Family, 1},
			{model.MemberMongoType().Image, 1},
			{model.MemberMongoType().Username, 1},
		},
	}}

	pipeline := mongo.Pipeline{lookupStage, matchStage, projectStage}

	// Aggregate
	cur, err := store.collectionMember.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	// Here's an array in which you can store the decoded documents
	var results []model.Member

	if err = cur.All(ctx, &results); err != nil {
		return nil, err
	}

	return &results, nil
}

//CheckFollow ...
func (store *DatabaseMemberStore) CheckFollow(ctx context.Context, follower string, following string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	objectIDFollower, err := primitive.ObjectIDFromHex(follower)
	if err != nil {
		return false, err
	}

	objectIDFollowing, err := primitive.ObjectIDFromHex(following)
	if err != nil {
		return false, err
	}

	filter := bson.M{
		model.FollowingMongoType().FollowerID:  objectIDFollower,
		model.FollowingMongoType().FollowingID: objectIDFollowing,
	}

	var result *model.Following

	err = store.collectionFollowing.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return false, nil
	}

	return true, nil
}

//Follow ...
func (store *DatabaseMemberStore) Follow(ctx context.Context, follower string, following string, follow bool) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, store.timeout)
	defer cancel()

	result, err := store.CheckFollow(ctx, follower, following)
	if err != nil {
		return false, err
	}
	if result == follow {
		return follow, nil
	}

	objectIDFollower, err := primitive.ObjectIDFromHex(follower)
	if err != nil {
		return false, err
	}

	objectIDFollowing, err := primitive.ObjectIDFromHex(following)
	if err != nil {
		return false, err
	}

	if follow {

		following := model.Following{
			FollowerID:   objectIDFollower,
			FollowingID:  objectIDFollowing,
			RegisterDate: time.Now(),
		}

		following.ID = primitive.NewObjectID()

		_, err := store.collectionFollowing.InsertOne(ctx, following)
		if err != nil {
			return false, errors.New("Error inserting new post: " + err.Error())
		}

		return true, nil

	} else {

		// Passing filter to matches specify result
		filter := bson.M{
			model.FollowingMongoType().FollowerID:  objectIDFollower,
			model.FollowingMongoType().FollowingID: objectIDFollowing,
		}

		_, err = store.collectionFollowing.DeleteMany(ctx, filter)
		if err != nil {
			return false, err
		}

		return false, nil
	}
}
