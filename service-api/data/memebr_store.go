package data

import (
	"context"

	"github.com/EsmaeilMazahery/wild/model"
)

// MemberStore is an interface to store Service data
type MemberStore interface {
	//add member to database
	Add(ctx context.Context, member *model.Member) (string, error)
	//find a member by id
	Find(ctx context.Context, id string) (*model.Member, error)
	//find a member by username
	FindByUsername(ctx context.Context, username string) (*model.Member, error)
	//find a member by mobile
	FindByMobile(ctx context.Context, mobile string) (*model.Member, error)
	//find a member y email
	FindByEmail(ctx context.Context, email string) (*model.Member, error)
	//change password of member by id , password must be hashed
	ChangePassword(ctx context.Context, ID string, NewPassword string) error
	//change image profile of member
	ChangeImageProfile(ctx context.Context, id string, newImage string) error
	//change image header of member profile
	ChangeImageHeader(ctx context.Context, id string, newImage string) error
	//edit member information
	Edit(ctx context.Context, member *model.Member) error
	//list of account suggest for a member , its a random list of users that does not follow by member
	Suggestion(ctx context.Context, memberID string) (*[]model.Member, error)
	//list of followers
	Followers(ctx context.Context, memberID string) (*[]model.Member, error)
	//list of following
	Followings(ctx context.Context, memberID string) (*[]model.Member, error)
	//check a member is follow by other member
	CheckFollow(ctx context.Context, follower string, following string) (bool, error)
	//follow a member by other member
	Follow(ctx context.Context, follower string, following string, follow bool) (bool, error)
}
