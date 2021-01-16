package model

import (
	"time"

	"github.com/EsmaeilMazahery/wild/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NilNotify is the nil value for Notify
var NilNotify Notify

// Notify contains Comment's information
type Notify struct {
	ID             primitive.ObjectID `bson:"_id"`
	Content        string             `bson:"content"`
	Type           enums.NotifyType   `bson:"type"`
	RegisterDate   time.Time          `bson:"register_date"`
	Read           bool               `bson:"read"`
	PostID         string             `bson:"post_id"`
	OwnerMemberID  primitive.ObjectID `bson:"owner_member_id"`
	TargetMemberID primitive.ObjectID `bson:"target_member_id"`
	OwnerMember    *Member            `bson:"owner_member,omitempty" json:"owner_member,omitempty"`
	TargetMember   *Member            `bson:"target_member,omitempty" json:"target_member,omitempty"`
}

//NotifyMongoType ...
func NotifyMongoType() struct {
	ID             string
	Content        string
	Type           string
	RegisterDate   string
	Read           string
	PostID         string
	OwnerMemberID  string
	TargetMemberID string
	OwnerMember    string
	TargetMember   string
} {
	return struct {
		ID             string
		Content        string
		Type           string
		RegisterDate   string
		Read           string
		PostID         string
		OwnerMemberID  string
		TargetMemberID string
		OwnerMember    string
		TargetMember   string
	}{
		"_id",
		"content",
		"type",
		"register_date",
		"read",
		"post_id",
		"owner_member_id",
		"target_member_id",
		"owner_member",
		"target_member",
	}
}
