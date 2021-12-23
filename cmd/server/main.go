package main

import (
	"context"
	"github.com/shaohsiung/memo/api/protobuf"
	"github.com/shaohsiung/memo/internal/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	conf, err := config.Load("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	memoServer, _ := InitMemoServer(conf.DSN)
	protobuf.RegisterMemoServer(grpcServer, memoServer)

	lis, err := net.Listen("tcp", conf.GRPC.Port)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Printf("grpc server start at %v\n", conf.GRPC.Port)
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	log.Println("shutting down the grpcServer")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	grpcServer.GracefulStop()
	<-ctx.Done()
	close(ch)
	log.Println("graceful shutdown end")
}
