package main

import (
	"database/sql"
	"fmt"
	"log"

	pb "github.com/habib/biz/grpc"
	_ "github.com/lib/pq"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"127.0.0.1", 5432, "user", "password", "database")
	postgresClient, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}
	err = postgresClient.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	query := "SELECT * FROM \"User\" WHERE \"id\" = $1;"
	userId := 5
	rows, err := postgresClient.Query(query, userId)
	log.Println(rows.Next())
	var user pb.User
	err = rows.Scan(&user.Id, &user.Name, &user.Family, &user.Age, &user.Sex, &user.CreatedAt)

	

	log.Println(rows.Next())
	log.Println(rows)
}
