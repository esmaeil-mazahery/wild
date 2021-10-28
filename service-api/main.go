package main

import (
	"context"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"

	_ "github.com/EsmaeilMazahery/wild/database/migrations" // database migrations
	"github.com/EsmaeilMazahery/wild/grpc/service"
	"github.com/EsmaeilMazahery/wild/infrastructure/auth"
	"github.com/EsmaeilMazahery/wild/infrastructure/constant"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_auth"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_comment"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_notify"
	"github.com/EsmaeilMazahery/wild/proto/pb/pb_post"
	"github.com/EsmaeilMazahery/wild/third-party/email"
	"github.com/EsmaeilMazahery/wild/third-party/sms"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	"github.com/EsmaeilMazahery/wild/data"
	"github.com/EsmaeilMazahery/wild/database"
	"github.com/EsmaeilMazahery/wild/gateway"
	"github.com/EsmaeilMazahery/wild/insecure"

	// Static files
	_ "github.com/EsmaeilMazahery/wild/statik"
)

func main() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	//get grpc port from environment
	port := os.Getenv("SERVE_PORT")

	//get jwt secure code from environment
	secretKey := os.Getenv("SERVICE_SECRET_KEY")

	tokenLoginDuration, err := strconv.Atoi(os.Getenv("SERVICE_TOKEN_LOGIN_DURATION"))
	if err != nil {
		log.Fatal("SERVICE_TOKEN_LOGIN_DURATION is incorrect", err)
	}

	//test connect to database
	c := database.GetClient()
	err = c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Info("Database Connected!")
	}

	db := c.Database(constant.Dbname())
	migrate.SetDatabase(db)
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		log.Fatal("Fail Run Migrations", err)
	}

	address := "0.0.0.0:" + port
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	jwtManager := auth.NewJWTManager(secretKey, time.Duration(tokenLoginDuration)*time.Second)
	interceptor := service.NewAuthInterceptor(jwtManager, auth.AccessibleRoles())

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
		grpc.MaxRecvMsgSize(1024*1024*100),
		grpc.MaxSendMsgSize(1024*1024*100),
	)

	smsProvider := sms.NewSmskavenegar()
	emailProvider := email.NewProvider(constant.EmailFrom(), constant.EmailPassword(), constant.EmailSmtphost(), constant.EmailSmtpport())
	cacheServer := data.NewRedisCacheStore()
	memberStore := data.NewDatabaseMemberStore(constant.Dbname())
	postStore := data.NewDatabasePostStore(constant.Dbname())
	notifyStore := data.NewDatabaseNotifyStore(constant.Dbname())
	commentStore := data.NewDatabaseCommentStore(constant.Dbname())

	authServer := service.NewAuthServer(memberStore, cacheServer, smsProvider, emailProvider, jwtManager, notifyStore)
	postServer := service.NewPostServer(memberStore, cacheServer, smsProvider, emailProvider, jwtManager, postStore, notifyStore)
	commentServer := service.NewCommentServer(memberStore, cacheServer, smsProvider, emailProvider, jwtManager, commentStore, notifyStore)
	notifyServer := service.NewNotifyServer(memberStore, cacheServer, smsProvider, emailProvider, jwtManager, notifyStore)

	pb_auth.RegisterAuthServiceServer(grpcServer, authServer)
	pb_post.RegisterPostServiceServer(grpcServer, postServer)
	pb_comment.RegisterCommentServiceServer(grpcServer, commentServer)
	pb_notify.RegisterNotifyServiceServer(grpcServer, notifyServer)

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", address)
	go func() {
		log.Fatal(grpcServer.Serve(listener))
	}()

	err = gateway.Run("dns:///" + address)
	log.Fatalln("gateway err : %s", err)
}
