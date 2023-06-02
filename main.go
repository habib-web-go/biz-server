package main

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/habib/biz/gen/grpc"
)

type server struct {
	pb.UnimplementedSQLServiceServer
}

var (
	postgresClient *sql.DB
)

func (s *server) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	userId := req.UserId
	var query string
	var err error
	var rows *sql.Rows

	if userId != nil {
		query = UserQuery + ` WHERE "id" = $1;`
		rows, err = postgresClient.Query(query, userId)
	} else {
		query = UserQuery + " LIMIT 100;"
		rows, err = postgresClient.Query(query)
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	users, err := rowsToUsers(rows)
	if err != nil {
		return nil, err
	}
	// todo set message Id how?
	// todo check and update message id
	return &pb.GetUsersResponse{Users: users, MessageId: req.GetMessageId() + 1}, nil
}
func (s *server) GetUsersWithSqlInject(ctx context.Context, req *pb.GetUsersWithSqlInjectRequest) (*pb.GetUsersResponse, error) {
	userId := req.UserId
	var query string
	var err error
	var rows *sql.Rows

	if userId != nil {
		query = UserQuery + ` WHERE "id" = ` + *userId + ";"
		rows, err = postgresClient.Query(query)
	} else {
		query = UserQuery + " LIMIT 100;"
		rows, err = postgresClient.Query(query)
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	users, err := rowsToUsers(rows)
	if err != nil {
		return nil, err
	}
	// todo set message Id how?
	// todo check and update message id
	return &pb.GetUsersResponse{Users: users, MessageId: req.GetMessageId() + 1}, nil
}
func main() {
	loadEnv()
	runPostgres()
	startServer()
}
