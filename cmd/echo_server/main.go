package main

import (
	"context"
	"go.uber.org/zap"
	"net"

	"github.com/chrisfesler/grpc_fun/pkg/app"
	"github.com/chrisfesler/grpc_fun/pkg/echo"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var log *zap.Logger

type server struct{}

func init() {
	log = app.LoggerWith(zap.String("some", "val"))
}

func (*server) Echo(ctx context.Context, in *echo.EchoMsg) (*echo.EchoMsg, error) {
	log.Info("Echo!", zap.String("message", in.Msg))

	return in, nil
}

func main() {
	log.Info("echo_server started at %v", zap.String("port", port))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen", zap.String("port", port), zap.Error(err))
	}
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("Serve failed", zap.Error(err))
	}
}
