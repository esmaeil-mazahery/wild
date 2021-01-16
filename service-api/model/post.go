package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NilPost is the nil value for Post
var NilPost Post

// Post contains Post's information
type Post struct {
	ID           primitive.ObjectID   `bson:"_id"`
	Content      string               `bson:"content"`
	Image        string               `bson:"image"`
	RegisterDate time.Time            `bson:"register_date"`
	Tags         []string             `bson:"tags"`
	Likes        []primitive.ObjectID `bson:"likes"`
	CountLikes   int64                `bson:"count_likes"`
	CommentIDs   []primitive.ObjectID `bson:"comment_ids"`
	Comments     *[]Comment           `bson:"comments,omitempty" json:"comments,omitempty"`
	MemberID     primitive.ObjectID   `bson:"member_id"`
	Member       *Member              `bson:"member,omitempty" json:"member,omitempty"`
}

//PostMongoType ...
func PostMongoType() struct {
	ID           string
	Content      string
	Image        string
	RegisterDate string
	Tags         string
	Likes        string
	CountLikes   string
	CommentIDs   string
	Comments     string
	MemberID     string
	Member       string
} {
	return struct {
		ID           string
		Content      string
		Image        string
		RegisterDate string
		Tags         string
		Likes        string
		CountLikes   string
		CommentIDs   string
		Comments     string
		MemberID     string
		Member       string
	}{
		"_id",
		"content",
		"image",
		"register_date",
		"tags",
		"likes",
		"count_likes",
		"comment_ids",
		"comments",
		"member_id",
		"member",
	}
}
