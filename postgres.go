package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	pb "github.com/habib/biz/gen/grpc"
	_ "github.com/lib/pq"
)

const UserQuery string = `SELECT "id", "name", "family", "age", "sex", "createdat" FROM "User"`

func runPostgres() {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASS")
	dbname := os.Getenv("POSTGRES_DB")
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	postgresClient, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}
	err = postgresClient.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
}

func rowsToUsers(rows *sql.Rows) ([]*pb.User, error) {
	defer rows.Close()
	var users []*pb.User
	for rows.Next() {
		var user pb.User
		err := rows.Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex, &user.CreatedAt)
		if err != nil {
			log.Fatal(err)
			continue
		}
		users = append(users, &user)
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return users, nil

}
