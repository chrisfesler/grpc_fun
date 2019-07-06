package main

import (
	"context"
	"github.com/chrisfesler/grpc_fun/pkg/app"
	"github.com/chrisfesler/grpc_fun/pkg/echo"
	"google.golang.org/grpc"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		app.Log.Fatalf("failed to connect to %v", address)
	}
	defer conn.Close()
	cli := echo.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := cli.Echo(ctx, &echo.EchoMsg{Msg: "Echo!"})
	if err != nil {
		app.Log.Fatalf("Shouting into the void: %v", err)
	}
	app.Log.Infof("Response: %v", r.Msg)
}
