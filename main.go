package main

import (
	"database/sql"

	pb "github.com/habib/biz/grpc"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

var (
	postgresClient *sql.DB
)

func main() {
	loadEnv()
	runPostgres()
	startServer()
}
