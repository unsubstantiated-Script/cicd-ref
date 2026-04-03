package main

import (
	"CICDRef/internal/server"
	"log"
)

func main() {
	srv := server.NewServer()

	log.Println("Server started")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
