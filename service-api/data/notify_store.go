package data

import (
	"context"

	"github.com/EsmaeilMazahery/wild/model"
)

// NotifyStore is an interface to store Service data
type NotifyStore interface {
	Add(ctx context.Context, notify *model.Notify) (string, error)
	AddLikePost(ctx context.Context, postID string, memberID string) error
	AddLikeComment(ctx context.Context, commentID string, memberID string) error
	AddFollow(ctx context.Context, followerID string, followingID string) error
	AddComment(ctx context.Context, postID string, memberID string, comment *model.Comment) error

	List(ctx context.Context, memberID string, page int64) (*[]model.Notify, error)
	Read(ctx context.Context, memberID string, id ...string) error
}
