package data

import (
	"context"

	"github.com/EsmaeilMazahery/wild/model"
)

// CommentStore is an interface to store Service data
type CommentStore interface {
	Add(ctx context.Context, comment *model.Comment) (string, error)
	List(ctx context.Context, postID string, page int64) (*[]model.Comment, error)
	MyComments(ctx context.Context, memberID string, page int64) (*[]model.Comment, error)
	GetLike(ctx context.Context, memberID string, commentID string) (bool, int, error)
	EditLike(ctx context.Context, memberID string, commentID string, like bool) (bool, error)
}
