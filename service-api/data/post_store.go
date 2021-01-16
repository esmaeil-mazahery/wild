package data

import (
	"context"

	"github.com/EsmaeilMazahery/wild/model"
)

// PostStore is an interface to store Service data
type PostStore interface {
	Register(ctx context.Context, post *model.Post) (string, error)
	List(ctx context.Context, memberID string, page int64) (*[]model.Post, error)
	Search(ctx context.Context, term string, page int64) (*[]model.Post, error)
	MyPosts(ctx context.Context, memberID string, page int64) (*[]model.Post, error)
	GetLike(ctx context.Context, memberID string, postID string) (bool, int, error)
	EditLike(ctx context.Context, memberID string, postID string, like bool) (bool, error)
}
