package command

import (
	"fmt"
	"log"
)

type Controller interface {
	HandleCommand(cmd string, args []string)
	Decode(encoded string) ([]byte, error)
	Info(fileName string) (map[string]any, error)
}

type controller struct{}

func NewController() Controller {
	return controller{}
}

func (c controller) HandleCommand(cmd string, args []string) {
	switch cmd {
	case "decode":
		if len(args) < 1 {
			log.Fatalf("decode command error: no args detected")
		}

		b, err := c.Decode(args[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	case "info":
		if len(args) < 1 {
			log.Fatalf("info command error: no args detected")
		}

		dict, err := c.Info(args[0])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Tracker URL: %s\n", dict["announce"])

		infoSection, ok := dict["info"].(map[string]any)
		if !ok {
			log.Fatalf("info command error: no info section detected")
		}
		fmt.Printf("Length: %d\n", infoSection["length"])

	default:
		log.Fatalf("Unknown command: %s", cmd)
	}
}
