package main

import (
	"log"
	"net"

	"github.com/Shulammite-Aso/filebox-email-service/pkg/config"
	"github.com/Shulammite-Aso/filebox-email-service/pkg/proto"
	"github.com/Shulammite-Aso/filebox-email-service/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	log.Println("email service listening on", c.Port)

	s := services.Server{}

	grpcServer := grpc.NewServer()

	proto.RegisterEmailServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
