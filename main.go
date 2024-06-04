package main

import (
	"log"

	"github.com/ab01fazl1/go-microservice/internal/database"
	"github.com/ab01fazl1/go-microservice/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialize database client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}