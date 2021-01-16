package service

import (
	"context"

	"github.com/EsmaeilMazahery/wild/data"
	"github.com/EsmaeilMazahery/wild/infrastructure/auth"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_models"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_notify"
	"github.com/EsmaeilMazahery/wild/third-party/email"
	"github.com/EsmaeilMazahery/wild/third-party/sms"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NotifyServer is the server for authentication
type NotifyServer struct {
	BaseServer
}

// NewNotifyServer returns a new auth server
func NewNotifyServer(
	memberStore data.MemberStore,
	cacheServer data.CacheServer,
	smsProvider sms.IProvider,
	emailProvider email.IProvider,
	jwtManager *auth.JWTManager,
	notifyStore data.NotifyStore,
) *NotifyServer {
	return &NotifyServer{
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

// List Notify
func (server *NotifyServer) List(ctx context.Context, req *pb_notify.ListRequest) (*pb_notify.ListResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get Notify: %v", err)
	}

	list, err := server.notifyStore.List(ctx, memberID, req.GetPage())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get Notify: %v", err)
	}

	var Notifies []*pb_models.Notify
	ExistMore := false
	for _, notify := range *list {

		//we retrive one item more to specify more items than 20 exist
		if len(Notifies) >= 20 {
			ExistMore = true
			break
		}

		p := pb_models.Notify{
			ID:      notify.ID.Hex(),
			Content: notify.Content,
			Type:    pb_models.Notify_NotifyType(notify.Type),
			TargetMember: &pb_models.Member{
				ID:       notify.TargetMember.ID.Hex(),
				Username: notify.TargetMember.Username,
				Name:     notify.TargetMember.Name,
				Family:   notify.TargetMember.Family,
				Image:    notify.TargetMember.Image,
			},
		}

		p.RegisterDate, _ = ptypes.TimestampProto(notify.RegisterDate)

		Notifies = append(Notifies, &p)
	}

	res := &pb_notify.ListResponse{
		Notifies:  Notifies,
		ExistMore: ExistMore,
	}
	return res, nil
}

// Read set read flag Notify
func (server *NotifyServer) Read(ctx context.Context, req *pb_notify.ReadRequest) (*pb_notify.ReadResponse, error) {
	memberID, err := server.GetAuthMemberID(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot set read: %v", err)
	}

	err = server.notifyStore.Read(ctx, memberID, req.IDs...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot set read: %v", err)
	}

	res := &pb_notify.ReadResponse{}
	return res, nil
}
