package main

import (
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/app/client"
	"github.com/codecrafters-io/bittorrent-starter-go/app/command"
)

func main() {
	commandController := command.NewController(client.New())
	commandController.HandleCommand(os.Args[1], os.Args[2:])
}
