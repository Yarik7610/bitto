package main

import (
	"log"
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/app/client"
	"github.com/codecrafters-io/bittorrent-starter-go/app/command"
)

func main() {
	commandController := command.NewController(client.New())
	if len(os.Args) < 3 {
		log.Fatal("Not enough arguments for commands work")
	}
	commandController.HandleCommand(os.Args[1], os.Args[2:])
}
