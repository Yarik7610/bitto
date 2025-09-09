package client

import "net"

const (
	PEER_ID   = "lHavHuZBkaYWXsuvjGJh"
	PEER_PORT = 6881
)

type Client struct {
	PeerID   string
	PeerPort int
	Peers    map[string]net.Conn
}

func New() *Client {
	return &Client{
		PeerID:   PEER_ID,
		PeerPort: PEER_PORT,
		Peers:    make(map[string]net.Conn),
	}
}
