package main

import (
	"log"
	"net"

	pb "github.com/habib/biz/grpc"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
}

func startServer() {

	lis, err := net.Listen("tcp", ":5052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSQLServiceServer(s, &server{})

	log.Println("gRPC server listening on port 5052")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
