package service

import (
	"context"

	"github.com/EsmaeilMazahery/wild/data"
	"github.com/EsmaeilMazahery/wild/infrastructure/auth"
	"github.com/EsmaeilMazahery/wild/infrastructure/exception"
	"github.com/EsmaeilMazahery/wild/model"
	"github.com/EsmaeilMazahery/wild/third-party/email"
	"github.com/EsmaeilMazahery/wild/third-party/sms"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// BaseServer ...
type BaseServer struct {
	cacheServer   data.CacheServer
	smsProvider   sms.IProvider
	emailProvider email.IProvider
	jwtManager    *auth.JWTManager
	memberStore   data.MemberStore
	notifyStore   data.NotifyStore
}

// NewBaseServer returns a new auth server
func NewBaseServer(
	cacheServer data.CacheServer,
	smsProvider sms.IProvider,
	emailProvider email.IProvider,
	jwtManager *auth.JWTManager,
	memberStore data.MemberStore,
	notifyStore data.NotifyStore,
) *BaseServer {
	return &BaseServer{
		cacheServer,
		smsProvider,
		emailProvider,
		jwtManager,
		memberStore,
		notifyStore,
	}
}

//GetAuthMemberID ...
func (server *BaseServer) GetAuthMemberID(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := server.jwtManager.VerifyLogin(accessToken)
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return claims.ID, nil
}

//GetAuthMember ...
func (server *BaseServer) GetAuthMember(ctx context.Context) (*model.Member, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &model.NilMember, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return &model.NilMember, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := server.jwtManager.VerifyLogin(accessToken)
	if err != nil {
		return &model.NilMember, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	member, err := server.memberStore.Find(ctx, claims.ID)

	return member, nil
}

//GetIP ...
func (server *BaseServer) GetIP(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", &exception.AppError{
			Description: "err get metadata",
			Err:         nil,
		}
	}

	values := md["x-forwarded-for"]
	ip := values[0]

	return ip, nil
}
