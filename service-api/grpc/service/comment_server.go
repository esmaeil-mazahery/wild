package service

import (
	"context"
	"time"

	"github.com/EsmaeilMazahery/wild/data"
	"github.com/EsmaeilMazahery/wild/infrastructure/auth"
	"github.com/EsmaeilMazahery/wild/model"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_comment"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_models"
	"github.com/EsmaeilMazahery/wild/third-party/email"
	"github.com/EsmaeilMazahery/wild/third-party/sms"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CommentServer is the server for authentication
type CommentServer struct {
	BaseServer
	commentStore data.CommentStore
}

// NewCommentServer returns a new auth server
func NewCommentServer(
	memberStore data.MemberStore,
	cacheServer data.CacheServer,
	smsProvider sms.IProvider,
	emailProvider email.IProvider,
	jwtManager *auth.JWTManager,
	commentStore data.CommentStore,
	notifyStore data.NotifyStore,
) *CommentServer {
	return &CommentServer{
		BaseServer: *NewBaseServer(
			cacheServer,
			smsProvider,
			emailProvider,
			jwtManager,
			memberStore,
			notifyStore,
		),
		commentStore: commentStore,
	}
}

// Add comment on a post
func (server *CommentServer) Add(ctx context.Context, req *pb_comment.AddRequest) (*pb_comment.AddResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot add comment: %v", err)
	}

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot add comment: %v", err)
	}

	objectPostID, err := primitive.ObjectIDFromHex(req.Comment.GetPostID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot add comment: %v", err)
	}

	comment := model.Comment{
		MemberID:     objectMemberID,
		Content:      req.Comment.GetContent(),
		Image:        req.Comment.GetImage(),
		PostID:       objectPostID,
		RegisterDate: time.Now(),
	}

	id, err := server.commentStore.Add(ctx, &comment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot add comment: %v", err)
	}

	err = server.notifyStore.AddComment(ctx, req.Comment.GetPostID(), memberID, &comment)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot set read: %v", err)
	}

	res := &pb_comment.AddResponse{
		ID: id,
	}
	return res, nil
}

// List retrive posts that must show to login user
func (server *CommentServer) List(ctx context.Context, req *pb_comment.ListRequest) (*pb_comment.ListResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get comments: %v", err)
	}

	list, err := server.commentStore.List(ctx, req.GetPostID(), req.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get comments: %v", err)
	}

	var Comments []*pb_models.Comment
	ExistMore := false
	for _, comment := range *list {

		//we retrive one item more to specify more items than 20 exist
		if len(Comments) >= 20 {
			ExistMore = true
			break
		}

		c := pb_models.Comment{
			ID:         comment.ID.Hex(),
			Content:    comment.Content,
			Image:      comment.Image,
			CountLikes: comment.CountLikes,
			MemberID:   comment.MemberID.Hex(),
			Member: &pb_models.Member{
				ID:       comment.Member.ID.Hex(),
				Username: comment.Member.Username,
				Name:     comment.Member.Name,
				Family:   comment.Member.Family,
				Image:    comment.Member.Image,
			},
			MemberLike: false,
			// Likes: []string,
		}

		c.RegisterDate, _ = ptypes.TimestampProto(comment.RegisterDate)

		for _, l := range comment.Likes {
			if l.Hex() == memberID {
				c.MemberLike = true
			}
		}

		Comments = append(Comments, &c)
	}

	res := &pb_comment.ListResponse{
		Comments:  Comments,
		ExistMore: ExistMore,
	}
	return res, nil
}

// Like like a post
func (server *CommentServer) Like(ctx context.Context, req *pb_comment.LikeRequest) (*pb_comment.LikeResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get comments: %v", err)
	}

	newLike, err := server.commentStore.EditLike(ctx, memberID, req.GetID(), req.GetLike())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get comments: %v", err)
	}

	if req.GetLike() {
		err = server.notifyStore.AddLikeComment(ctx, req.GetID(), memberID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot set read: %v", err)
		}
	}

	res := &pb_comment.LikeResponse{
		Result: newLike,
	}
	return res, nil
}

// MyComments list of my comments
func (server *CommentServer) MyComments(ctx context.Context, req *pb_comment.MyCommentsRequest) (*pb_comment.MyCommentsResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get comments: %v", err)
	}

	list, err := server.commentStore.MyComments(ctx, memberID, req.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get comments: %v", err)
	}

	var Comments []*pb_models.Comment
	ExistMore := false
	for _, post := range *list {

		//we retrive one item more to specify more items than 20 exist
		if len(Comments) >= 20 {
			ExistMore = true
			break
		}

		c := pb_models.Comment{
			ID:         post.ID.Hex(),
			Content:    post.Content,
			Image:      post.Image,
			CountLikes: post.CountLikes,
			MemberID:   post.MemberID.Hex(),
			Member: &pb_models.Member{
				ID:       post.Member.ID.Hex(),
				Username: post.Member.Username,
				Name:     post.Member.Name,
				Family:   post.Member.Family,
				Image:    post.Member.Image,
			},
			MemberLike: false,
			// Likes: []string,
		}

		c.RegisterDate, _ = ptypes.TimestampProto(post.RegisterDate)

		for _, l := range post.Likes {
			if l.Hex() == memberID {
				c.MemberLike = true
			}
		}

		Comments = append(Comments, &c)
	}

	res := &pb_comment.MyCommentsResponse{
		Comments:  Comments,
		ExistMore: ExistMore,
	}
	return res, nil
}
