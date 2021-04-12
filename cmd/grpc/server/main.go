package main

import (
	"fmt"
	"log"
	"net"

	otr "github.com/tengla/nsr-parser/trainroute"
	pb "github.com/tengla/nsr-parser/trainroute/protos"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.ServerOption
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090))
	if err != nil {
		log.Fatal(err.Error())
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterOperatingTrainRouteServiceServer(grpcServer, otr.NewOtrServer())
	log.Println("Starting server at :9090")
	grpcServer.Serve(lis)
}
