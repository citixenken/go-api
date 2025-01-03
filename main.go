package main

import (
	"github.com/citixenken/go-api.git/internal/db"
	"github.com/citixenken/go-api.git/internal/todo"
	"github.com/citixenken/go-api.git/internal/transport"
	"log"
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
