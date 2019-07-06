package main

import (
	"context"
	"net"

	"github.com/chrisfesler/grpc_fun/pkg/app"
	"github.com/chrisfesler/grpc_fun/pkg/echo"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *echo.EchoMsg) (*echo.EchoMsg, error) {
	app.Log.Infow("Echo!",
		"message", in.Msg,
	)

	return in, nil
}

func main() {
	app.Log.Infof("echo_server started at %v", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		app.Log.Fatalf("cannot listen at port %s: %v", port, err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		app.Log.Fatalf("Serve failed: %v", err)
	}
}
