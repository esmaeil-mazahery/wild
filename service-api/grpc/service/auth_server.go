package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/EsmaeilMazahery/wild/data"
	"github.com/EsmaeilMazahery/wild/infrastructure/auth"
	"github.com/EsmaeilMazahery/wild/infrastructure/random"
	"github.com/EsmaeilMazahery/wild/model"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_auth"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_models"
	"github.com/EsmaeilMazahery/wild/third-party/email"
	"github.com/EsmaeilMazahery/wild/third-party/sms"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is the server for authentication
type AuthServer struct {
	BaseServer
}

// NewAuthServer returns a new auth server
func NewAuthServer(
	memberStore data.MemberStore,
	cacheServer data.CacheServer,
	smsProvider sms.IProvider,
	emailProvider email.IProvider,
	jwtManager *auth.JWTManager,
	notifyStore data.NotifyStore,
) *AuthServer {
	return &AuthServer{
		BaseServer: *NewBaseServer(
			cacheServer,
			smsProvider,
			emailProvider,
			jwtManager,
			memberStore,
			notifyStore,
		),
	}
}

// Login is a unary RPC to login member
func (server *AuthServer) Login(ctx context.Context, req *pb_auth.LoginRequest) (*pb_auth.LoginResponse, error) {

	req.Username = strings.TrimLeft(req.GetUsername(), "0")

	member, err := server.memberStore.FindByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "incorrect username/password: %v", err)
	}

	if member == nil || !member.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.PermissionDenied, "incorrect username/password")
	}

	token, err := server.jwtManager.GenerateLogin(member)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	memberResult := &pb_models.Member{
		Token:    token,
		Username: member.Username,
		Name:     member.Name,
		Email:    member.Email,
		Family:   member.Family,
		Image:    member.Image,
		Mobile:   member.Mobile,
		Enable:   member.Enable,
	}

	res := &pb_auth.LoginResponse{
		Member: memberResult,
	}
	return res, nil
}

// ForgetPassword is a unary RPC to login member
func (server *AuthServer) ForgetPassword(ctx context.Context, req *pb_auth.ForgetPasswordRequest) (*pb_auth.ForgetPasswordResponse, error) {

	member, err := server.memberStore.FindByUsername(ctx, req.GetUsername())
	if err != nil || member == nil {
		return nil, status.Errorf(codes.NotFound, "cannot find member: %v", err)
	}

	verifyCode := strconv.FormatInt(random.Int(1111, 9999), 10)
	message := fmt.Sprintf("کد بازیابی :  %s", verifyCode)

	// err = server.smsProvider.Send(ctx, message, member.Mobile)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal err : %v", err)
	// }

	err = server.emailProvider.Send(ctx, message, []string{member.Email})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal err : %v", err)
	}

	err = server.cacheServer.Set(ctx, "ForgetPassword:"+req.GetUsername(), verifyCode, 5*time.Minute)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal err : %v", err)
	}

	res := &pb_auth.ForgetPasswordResponse{}
	return res, nil
}

// ForgetPasswordChange is a unary RPC to login member
func (server *AuthServer) ForgetPasswordChange(ctx context.Context, req *pb_auth.ForgetPasswordChangeRequest) (*pb_auth.ForgetPasswordChangeResponse, error) {

	member, err := server.memberStore.FindByUsername(ctx, req.GetUsername())
	if err != nil || member == nil {
		return nil, status.Errorf(codes.NotFound, "cannot find member: %v", err)
	}

	verifyCode, err := server.cacheServer.Get(ctx, "ForgetPassword:"+req.GetUsername())
	if err == nil && verifyCode == req.GetVerifyCode() {

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot hash password")
		}

		err = server.memberStore.ChangePassword(ctx, member.ID.Hex(), string(hashedPassword))
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Internal err : %v", err)
		}

		// err = server.smsProvider.Send(ctx, "رمزعبور شما با موفقیت تغییرکرد", member.Mobile)
		// if err != nil {
		// 	return nil, status.Errorf(codes.Internal, "Internal err : %v", err)
		// }

		token, err := server.jwtManager.GenerateLogin(member)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate access token")
		}

		memberResult := &pb_models.Member{
			Token:  token,
			Mobile: member.Mobile,
			Email:  member.Email,
			Image:  member.Image,
			Name:   member.Name,
			Family: member.Family,
			Enable: member.Enable,
		}

		res := &pb_auth.ForgetPasswordChangeResponse{
			Member: memberResult,
		}
		return res, nil
	}

	return nil, status.Errorf(codes.Internal, "کد تایید اشتباه است")
}

// Register is a unary RPC to login member
func (server *AuthServer) Register(ctx context.Context, req *pb_auth.RegisterRequest) (*pb_auth.RegisterResponse, error) {

	req.Member.Mobile = strings.TrimLeft(req.Member.GetMobile(), "0")

	member, err := server.memberStore.FindByUsername(ctx, req.Member.Username)
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "username is taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Member.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot hash password")
	}

	member = &model.Member{
		Username:     req.Member.Username,
		Name:         req.Member.Name,
		Family:       req.Member.Family,
		Image:        req.Member.Image,
		Password:     string(hashedPassword),
		Mobile:       req.Member.Mobile,
		Email:        req.Member.Email,
		Enable:       req.Member.Enable,
		RegisterDate: time.Now(),
	}

	id, err := server.memberStore.Add(ctx, member)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot add member: %v", err)
	}

	token, err := server.jwtManager.GenerateLogin(member)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	req.Member.Token = token
	req.Member.ID = id

	res := &pb_auth.RegisterResponse{
		Member: req.Member,
	}
	return res, nil
}

//ChangeImageProfile ...
func (server *AuthServer) ChangeImageProfile(ctx context.Context, req *pb_auth.ChangeImageProfileRequest) (*pb_auth.ChangeImageProfileResponse, error) {
	member, err := server.GetAuthMember(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	err = server.memberStore.ChangeImageProfile(ctx, member.ID.Hex(), req.GetImageURL())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	res := &pb_auth.ChangeImageProfileResponse{}
	return res, nil
}

//ChangeImageHeader ...
func (server *AuthServer) ChangeImageHeader(ctx context.Context, req *pb_auth.ChangeImageHeaderRequest) (*pb_auth.ChangeImageHeaderResponse, error) {
	member, err := server.GetAuthMember(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	err = server.memberStore.ChangeImageHeader(ctx, member.ID.Hex(), req.GetImageURL())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	res := &pb_auth.ChangeImageHeaderResponse{}
	return res, nil
}

// ChangePassword is a unary RPC to change passwotd member
func (server *AuthServer) ChangePassword(ctx context.Context, req *pb_auth.ChangePasswordRequest) (*pb_auth.ChangePasswordResponse, error) {
	req.Username = strings.TrimLeft(req.GetUsername(), "0")

	member, err := server.memberStore.FindByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "incorrect username/password: %v", err)
	}

	if member == nil || !member.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.PermissionDenied, "incorrect username/password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot hash password")
	}

	err = server.memberStore.ChangePassword(ctx, member.ID.Hex(), string(hashedPassword))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal err : %v", err)
	}

	member.Password = string(hashedPassword)
	token, err := server.jwtManager.GenerateLogin(member)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb_auth.ChangePasswordResponse{
		Member: &pb_models.Member{
			Token:  token,
			Name:   member.Name,
			Email:  member.Email,
			Family: member.Family,
			Image:  member.Image,
			Mobile: member.Mobile,
			Enable: member.Enable,
		},
	}
	return res, nil
}

// Suggestion ..
func (server *AuthServer) Suggestion(ctx context.Context, req *pb_auth.SuggestionRequest) (*pb_auth.SuggestionResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get comments: %v", err)
	}

	list, err := server.memberStore.Suggestion(ctx, memberID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal err : %v", err)
	}

	var Members []*pb_models.Member
	for _, member := range *list {
		m := pb_models.Member{
			ID:       member.ID.Hex(),
			Username: member.Username,
			Name:     member.Name,
			Family:   member.Family,
			Image:    member.Image,
		}

		Members = append(Members, &m)
	}

	res := &pb_auth.SuggestionResponse{
		Members: Members,
	}
	return res, nil
}

// Followers ..
func (server *AuthServer) Followers(ctx context.Context, req *pb_auth.FollowersRequest) (*pb_auth.FollowersResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get comments: %v", err)
	}

	list, err := server.memberStore.Followers(ctx, memberID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal err : %v", err)
	}

	var Members []*pb_models.Member
	for _, member := range *list {
		m := pb_models.Member{
			ID:           member.ID.Hex(),
			Username:     member.Username,
			Name:         member.Name,
			Family:       member.Family,
			Image:        member.Image,
			MemberFollow: true,
		}

		Members = append(Members, &m)
	}

	res := &pb_auth.FollowersResponse{
		Members: Members,
	}
	return res, nil
}

// Followings ..
func (server *AuthServer) Followings(ctx context.Context, req *pb_auth.FollowingsRequest) (*pb_auth.FollowingsResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get comments: %v", err)
	}

	list, err := server.memberStore.Followings(ctx, memberID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal err : %v", err)
	}

	var Members []*pb_models.Member
	for _, member := range *list {
		m := pb_models.Member{
			ID:           member.ID.Hex(),
			Username:     member.Username,
			Name:         member.Name,
			Family:       member.Family,
			Image:        member.Image,
			MemberFollow: member.MemberFollow,
		}

		Members = append(Members, &m)
	}

	res := &pb_auth.FollowingsResponse{
		Members: Members,
	}
	return res, nil
}

// Follow ..
func (server *AuthServer) Follow(ctx context.Context, req *pb_auth.FollowRequest) (*pb_auth.FollowResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get comments: %v", err)
	}

	result, err := server.memberStore.Follow(ctx, memberID, req.GetMemberID(), req.GetFollow())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal err : %v", err)
	}
	if req.GetFollow() {
		err = server.notifyStore.AddFollow(ctx, memberID, req.GetMemberID())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot set read: %v", err)
		}
	}

	res := &pb_auth.FollowResponse{
		Result: result,
	}
	return res, nil
}
