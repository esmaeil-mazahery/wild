package data

import (
	"context"

	"github.com/EsmaeilMazahery/wild/model"
)

// MemberStore is an interface to store Service data
type MemberStore interface {
	Add(ctx context.Context, member *model.Member) (string, error)
	Find(ctx context.Context, id string) (*model.Member, error)
	FindByUsername(ctx context.Context, username string) (*model.Member, error)
	FindByMobile(ctx context.Context, mobile string) (*model.Member, error)
	FindByEmail(ctx context.Context, email string) (*model.Member, error)
	ChangePassword(ctx context.Context, ID string, NewPassword string) error
	ChangeImageProfile(ctx context.Context, id string, newImage string) error
	ChangeImageHeader(ctx context.Context, id string, newImage string) error
	Edit(ctx context.Context, member *model.Member) error
	Suggestion(ctx context.Context, memberID string) (*[]model.Member, error)
	Followers(ctx context.Context, memberID string) (*[]model.Member, error)
	Followings(ctx context.Context, memberID string) (*[]model.Member, error)

	CheckFollow(ctx context.Context, follower string, following string) (bool, error)
	Follow(ctx context.Context, follower string, following string, follow bool) (bool, error)
}
