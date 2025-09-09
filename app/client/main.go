package client

import (
	"net"

	"github.com/codecrafters-io/bittorrent-starter-go/app/constants"
)

type Client struct {
	PeerID   string
	PeerPort int
	Peers    map[string]net.Conn
}

func New() *Client {
	return &Client{
		PeerID:   constants.PEER_ID,
		PeerPort: constants.PEER_PORT,
		Peers:    make(map[string]net.Conn),
	}
}
