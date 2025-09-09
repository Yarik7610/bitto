package command

import (
	"fmt"
	"log"

	"github.com/codecrafters-io/bittorrent-starter-go/app/client"
)

type Controller interface {
	HandleCommand(cmd string, args []string)
	Decode(encoded string) (string, error)
	Info(fileName string) (*Torrent, error)
	Peers(fileName string) ([]PeerSocket, error)
	Handshake(fileName string, peerSocketString string) (*HandshakeMessage, error)
}

type controller struct {
	client *client.Client
}

func NewController(client *client.Client) Controller {
	return controller{
		client: client,
	}
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
	case "peers":
		if len(args) < 1 {
			log.Fatalf("peers command error: no args detected")
		}

		peers, err := c.Peers(args[0])
		if err != nil {
			log.Fatalf("peers command error: %v", err)
		}

		for _, peer := range peers {
			fmt.Println(peer)
		}
	case "handshake":
		if len(args) != 2 {
			log.Fatalf("handshake command error: wrong args count detected")
		}

		peerResponse, err := c.Handshake(args[0], args[1])
		if err != nil {
			log.Fatalf("hahdshake command error: %v", err)
		}

		fmt.Printf("Peer ID: %x\n", peerResponse.PeerID)
	default:
		log.Fatalf("Unknown command: %s", cmd)
	}
}
