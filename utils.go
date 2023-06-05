package main

import (
	"log"
	"net"
	"os"

	pb "github.com/habib-web-go/biz-server/gen/grpc"
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
	port := os.Getenv("BIZ_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSQLServiceServer(s, &server{})

	log.Printf("gRPC server listening on port %s\n", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
