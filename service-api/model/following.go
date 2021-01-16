package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NilFollowing is the nil value for Following
var NilFollowing Member

// Following contains member's Following
type Following struct {
	ID primitive.ObjectID `bson:"_id"`

	FollowerID   primitive.ObjectID `bson:"follower_id"`
	FollowingID  primitive.ObjectID `bson:"following_id"`
	RegisterDate time.Time          `bson:"register_date"`
}

//FollowingMongoType ...
func FollowingMongoType() struct {
	ID           string
	FollowerID   string
	FollowingID  string
	RegisterDate string
} {
	return struct {
		ID           string
		FollowerID   string
		FollowingID  string
		RegisterDate string
	}{
		"_id",
		"follower_id",
		"following_id",
		"register_date",
	}
}
