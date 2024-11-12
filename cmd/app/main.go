package main

import (
	"log"
	"fmt"
	"github.com/MobasirSarkar/go-vote-app/internal/server"
)

func main() {
	server := server.NewServer()
	fmt.Println("Hello")

	err := server.ListenAndServe()
	if err != nil {
		log.Panicf("http server error: %s\n", err)
	}
}
