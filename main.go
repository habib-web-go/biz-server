package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	pb "github.com/habib/biz/grpc"
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
	if userId != nil {
		query = "SELECT * FROM \"User\" WHERE id = $1;"
	} else {
		query = "SELECT * FROM \"User\" LIMIT 100;"
	}
	rows, err := postgresClient.Query(query, userId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []*pb.User
	for rows.Next() {
		var user *pb.User
		err := rows.Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
			continue
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// todo set message Id how?
	// todo check and update message id
	return &pb.GetUsersResponse{Users: users, MessageId: req.GetMessageId() + 1}, nil

}
func (s *server) getUsersWithSqlInject(ctx context.Context, req *pb.GetUsersWithSqlInjectRequest) (*pb.GetUsersResponse, error) {
	userId := req.UserId
	var query string
	if userId != nil {
		query = fmt.Sprintf("SELECT * FROM \"User\" WHERE id = %s;", *userId)
	} else {
		query = "SELECT * FROM \"User\" LIMIT 100;"
	}
	rows, err := postgresClient.Query(query, userId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []*pb.User
	for rows.Next() {
		var user *pb.User
		err := rows.Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
			continue
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
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
