package main

import (
	"log"
	"my_first_api/internal/db"
	"my_first_api/internal/todo"
	"my_first_api/internal/transport"
)

func main() {
	d, err := db.New("postgres", "example", "postgres", "localhost", 5432)
	if err != nil {
		log.Fatal(err)
	}
	svc := todo.NewService(d)

	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
