package main

import (
	"context"
	"go.uber.org/zap"
	"time"

	"github.com/chrisfesler/grpc_fun/pkg/app"
	"github.com/chrisfesler/grpc_fun/pkg/echo"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

var log = app.LoggerWith(zap.String("some", "val"))

func main() {
	defer log.Sync()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect", zap.String("address", address), zap.Error(err))
	}
	defer conn.Close()
	cli := echo.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // go vet complains if we don't do this
	r, err := cli.Echo(ctx, &echo.EchoMsg{Msg: "Echo!"})
	if err != nil {
		log.Fatal("Shouting into the void", zap.Error(err))
	}
	log.Info("Heard an echo!", zap.String("message", r.Msg))
}
