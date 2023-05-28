package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

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
