package main

import (
	"context"
	"log"
	"net"

	"github.com/thaidzai285/dzai-mp3-crawler-service/internal/pkg/dcrawl"
	"github.com/thaidzai285/dzai-mp3-protobuf/pkg/pb"
	"google.golang.org/grpc"
)

var (
	ctxCancel context.CancelFunc
	ctx context.Context
)

func main() {
	ctx, ctxCancel = context.WithCancel(context.Background())

	go func() {
		err := serverGRPC()
		if err != nil {
			log.Fatal("GRPC server error: ", err)
			return
		}
	}()

	<-ctx.Done()
}

func serverGRPC() error {
	defer ctxCancel()
	listen, err := net.Listen("tcp", ":4001")
	if err != nil {
		return err
	}

	crawlerService := dcrawl.NewCrawlerService()
	s := grpc.NewServer()
	pb.RegisterCrawlerServer(s, crawlerService)

	log.Println("GRPC listen on 4001")
	return s.Serve(listen)
}
