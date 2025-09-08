package main

import (
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/app/client"
)

func main() {
	cl := client.New()
	cl.HandleCommand(os.Args[1], os.Args[2:])
}
