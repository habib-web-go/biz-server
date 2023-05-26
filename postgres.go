package main

import (
	"database/sql"
	"fmt"
	"os"
)

func runPostgres() {
	url := os.Getenv("POSTGRES_URL")
	postgresClient, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}
	defer postgresClient.Close()
	err = postgresClient.Ping()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
}
