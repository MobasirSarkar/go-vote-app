package main

import (
	"fmt"
	"github.com/MobasirSarkar/go-vote-app/internal/server"
	"log"
)

func main() {
	server := server.NewServer()
	fmt.Println("Hello")

	err := server.ListenAndServe()
	if err != nil {
		log.Panicf("http server error: %s\n", err)
	}
}
