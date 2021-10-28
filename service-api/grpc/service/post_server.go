package service

import (
	"context"
	"regexp"
	"time"

	"github.com/EsmaeilMazahery/wild/data"
	"github.com/EsmaeilMazahery/wild/infrastructure/auth"
	"github.com/EsmaeilMazahery/wild/model"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_models"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_post"
	"github.com/EsmaeilMazahery/wild/third-party/email"
	"github.com/EsmaeilMazahery/wild/third-party/sms"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PostServer is the server for authentication
type PostServer struct {
	BaseServer
	postStore data.PostStore
}

// NewPostServer returns a new auth server
func NewPostServer(
	memberStore data.MemberStore,
	cacheServer data.CacheServer,
	smsProvider sms.IProvider,
	emailProvider email.IProvider,
	jwtManager *auth.JWTManager,
	postStore data.PostStore,
	notifyStore data.NotifyStore,
) *PostServer {
	return &PostServer{
		BaseServer: *NewBaseServer(
			cacheServer,
			smsProvider,
			emailProvider,
			jwtManager,
			memberStore,
			notifyStore,
		),
		postStore: postStore,
	}
}

// Register add a post
func (server *PostServer) Register(ctx context.Context, req *pb_post.RegisterRequest) (*pb_post.RegisterResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot add post: %v", err)
	}

	objectMemberID, err := primitive.ObjectIDFromHex(memberID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot add post: %v", err)
	}

	r, _ := regexp.Compile("/#\\w+/g")
	tags := r.FindAllString(req.Post.GetContent(), -1)

	post := model.Post{
		MemberID:     objectMemberID,
		Content:      req.Post.GetContent(),
		Image:        req.Post.GetImage(),
		RegisterDate: time.Now(),
		Tags:         tags,
	}

	id, err := server.postStore.Register(ctx, &post)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot add post: %v", err)
	}

	res := &pb_post.RegisterResponse{
		ID: id,
	}
	return res, nil
}

// List retrive posts that must show to login user
func (server *PostServer) List(ctx context.Context, req *pb_post.ListRequest) (*pb_post.ListResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get posts: %v", err)
	}

	list, err := server.postStore.List(ctx, memberID, req.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get posts: %v", err)
	}

	var Posts []*pb_models.Post
	ExistMore := false
	for _, post := range *list {

		//we retrive one item more to specify more items than 20 exist
		if len(Posts) >= 20 {
			ExistMore = true
			break
		}

		p := pb_models.Post{
			ID:         post.ID.Hex(),
			Content:    post.Content,
			Image:      post.Image,
			CountLikes: post.CountLikes,
			MemberID:   post.MemberID.Hex(),
			Tags:       post.Tags,
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

		p.RegisterDate, _ = ptypes.TimestampProto(post.RegisterDate)

		// var Comments []*pb_models.Comment
		// for _, comment := range post.Comments {
		// 	c := pb_models.Comment{
		// 		Content:    comment.Content,
		// 		Image:      comment.Image,
		// 		CountLikes: comment.CountLikes,
		// 		MemberID:   comment.MemberID.Hex(),
		// 	}

		// 	Comments = append(Comments, &c)
		// }

		for _, l := range post.Likes {
			if l.Hex() == memberID {
				p.MemberLike = true
			}
		}

		Posts = append(Posts, &p)
	}

	res := &pb_post.ListResponse{
		Posts:     Posts,
		ExistMore: ExistMore,
	}
	return res, nil
}

// Like like a post
func (server *PostServer) Like(ctx context.Context, req *pb_post.LikeRequest) (*pb_post.LikeResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get posts: %v", err)
	}

	newLike, err := server.postStore.EditLike(ctx, memberID, req.GetID(), req.GetLike())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get posts: %v", err)
	}

	if req.GetLike() {
		err = server.notifyStore.AddLikePost(ctx, req.GetID(), memberID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot set read: %v", err)
		}
	}

	res := &pb_post.LikeResponse{
		Result: newLike,
	}
	return res, nil
}

// MyPosts list of my post
func (server *PostServer) MyPosts(ctx context.Context, req *pb_post.MyPostsRequest) (*pb_post.MyPostsResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get posts: %v", err)
	}

	list, err := server.postStore.MyPosts(ctx, memberID, req.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get posts: %v", err)
	}

	var Posts []*pb_models.Post
	ExistMore := false
	for _, post := range *list {

		//we retrive one item more to specify more items than 20 exist
		if len(Posts) >= 20 {
			ExistMore = true
			break
		}

		p := pb_models.Post{
			ID:         post.ID.Hex(),
			Content:    post.Content,
			Image:      post.Image,
			CountLikes: post.CountLikes,
			MemberID:   post.MemberID.Hex(),
			Tags:       post.Tags,
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

		p.RegisterDate, _ = ptypes.TimestampProto(post.RegisterDate)

		// var Comments []*pb_models.Comment
		// for _, comment := range post.Comments {
		// 	c := pb_models.Comment{
		// 		Content:    comment.Content,
		// 		Image:      comment.Image,
		// 		CountLikes: comment.CountLikes,
		// 		MemberID:   comment.MemberID.Hex(),
		// 	}

		// 	Comments = append(Comments, &c)
		// }

		for _, l := range post.Likes {
			if l.Hex() == memberID {
				p.MemberLike = true
			}
		}

		Posts = append(Posts, &p)
	}

	res := &pb_post.MyPostsResponse{
		Posts:     Posts,
		ExistMore: ExistMore,
	}
	return res, nil
}

// Search in all post term
func (server *PostServer) Search(ctx context.Context, req *pb_post.SearchRequest) (*pb_post.SearchResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get posts: %v", err)
	}

	list, err := server.postStore.Search(ctx, req.GetTerm(), req.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get posts: %v", err)
	}

	var Posts []*pb_models.Post
	ExistMore := false
	for _, post := range *list {
		//we retrive one item more to specify more items than 20 exist
		if len(Posts) >= 20 {
			ExistMore = true
			break
		}

		p := pb_models.Post{
			ID:         post.ID.Hex(),
			Content:    post.Content,
			Image:      post.Image,
			CountLikes: post.CountLikes,
			MemberID:   post.MemberID.Hex(),
			Tags:       post.Tags,
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

		p.RegisterDate, _ = ptypes.TimestampProto(post.RegisterDate)

		// var Comments []*pb_models.Comment
		// for _, comment := range post.Comments {
		// 	c := pb_models.Comment{
		// 		Content:    comment.Content,
		// 		Image:      comment.Image,
		// 		CountLikes: comment.CountLikes,
		// 		MemberID:   comment.MemberID.Hex(),
		// 	}

		// 	Comments = append(Comments, &c)
		// }

		for _, l := range post.Likes {
			if l.Hex() == memberID {
				p.MemberLike = true
			}
		}

		Posts = append(Posts, &p)
	}

	res := &pb_post.SearchResponse{
		Posts:     Posts,
		ExistMore: ExistMore,
	}
	return res, nil
}
