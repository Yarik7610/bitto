package command

import (
	"fmt"
	"log"
)

type Controller interface {
	HandleCommand(cmd string, args []string)
	Decode(encoded string) (string, error)
	Info(fileName string) (*InfoResponse, error)
}

type controller struct{}

func NewController() Controller {
	return controller{}
}

func (c controller) HandleCommand(cmd string, args []string) {
	switch cmd {
	case "decode":
		if len(args) < 1 {
			log.Fatal("decode command error: no args detected")
		}

		decoded, err := c.Decode(args[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(decoded)
	case "info":
		if len(args) < 1 {
			log.Fatalf("info command error: no args detected")
		}

		response, err := c.Info(args[0])
		if err != nil {
			log.Fatalf("info command error: %v", err)
		}
		fmt.Print(response)
	default:
		log.Fatalf("Unknown command: %s", cmd)
	}
}
