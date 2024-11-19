package main

import (
	"log"

	"github.com/soupaulodev/chat-server/server"
)

func main() {
	server := server.NewServer(":8080")
	log.Println("Server running on port 8080")
	server.Start()
}