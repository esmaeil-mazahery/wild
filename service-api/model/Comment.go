package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NilComment is the nil value for Comment
var NilComment Comment

// Comment contains Comment's information
type Comment struct {
	ID           primitive.ObjectID   `bson:"_id"`
	Content      string               `bson:"content"`
	Image        string               `bson:"image"`
	RegisterDate time.Time            `bson:"register_date"`
	Likes        []primitive.ObjectID `bson:"likes"`
	CountLikes   int64                `bson:"count_likes"`
	PostID       primitive.ObjectID   `bson:"post_id"`
	MemberID     primitive.ObjectID   `bson:"member_id"`
	Member       *Member              `bson:"member,omitempty" json:"member,omitempty"`
}

//CommentMongoType ...
func CommentMongoType() struct {
	ID           string
	Content      string
	Image        string
	RegisterDate string
	Likes        string
	CountLikes   string
	PostID       string
	MemberID     string
	Member       string
} {
	return struct {
		ID           string
		Content      string
		Image        string
		RegisterDate string
		Likes        string
		CountLikes   string
		PostID       string
		MemberID     string
		Member       string
	}{
		"_id",
		"content",
		"image",
		"register_date",
		"likes",
		"count_likes",
		"post_id",
		"member_id",
		"member",
	}
}
