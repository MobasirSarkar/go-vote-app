package main

import (
	"log"

	"github.com/MobasirSarkar/go-vote-app/internal/server"
)

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		log.Panicf("http server error: %s\n", err)
	}
}
