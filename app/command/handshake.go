package command

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/codecrafters-io/bittorrent-starter-go/app/utils"
)

const PROTOCOL_STRING = "BitTorrent protocol"

type HandshakeMessage struct {
	ProtocolString string
	PeerID         string
	*Torrent
}

func (m *HandshakeMessage) Bytes() []byte {
	var res bytes.Buffer

	res.WriteString(fmt.Sprintf("%d", len(m.ProtocolString)))
	res.WriteString(m.ProtocolString)
	res.WriteString(strings.Repeat("0", 8))
	res.WriteString(string(m.InfoHash))
	res.WriteString(m.PeerID)

	return res.Bytes()
}

func (c controller) Handshake(fileName string, peerSocketString string) (net.Conn, error) {
	torrent, err := c.Info(fileName)
	if err != nil {
		return nil, fmt.Errorf("info command error: %v", err)
	}
	torrentPeers, err := c.Peers(fileName)
	if err != nil {
		return nil, fmt.Errorf("peers command error: %v", err)
	}

	err = checkPeerSocketValid(torrentPeers, peerSocketString)
	if err != nil {
		return nil, err
	}

	conn, err := dialPeer(peerSocketString)
	if err != nil {
		return nil, err
	}

	c.client.Peers[utils.GetRemoteAddrString(conn)] = conn

	message := &HandshakeMessage{
		ProtocolString: PROTOCOL_STRING,
		PeerID:         c.client.PeerID,
		Torrent:        torrent,
	}

	fmt.Println(string(message.Bytes()))
	_, err = conn.Write(message.Bytes())
	if err != nil {
		return nil, fmt.Errorf("write to peer error: %v", err)
	}

	response := make([]byte, 0)
	_, err = conn.Read(response)
	if err != nil {
		return nil, fmt.Errorf("read from peer error: %v", err)
	}
	return conn, err
}

func checkPeerSocketValid(torrentPeers []PeerSocket, peerSocketString string) error {
	splitted := strings.Split(peerSocketString, ":")
	if len(splitted) != 2 {
		return fmt.Errorf("wrong peer socket string you want to handshake with")
	}
	port, err := strconv.ParseUint(splitted[1], 10, 64)
	if err != nil {
		return err
	}

	desiredPeer := PeerSocket{IP: net.ParseIP(splitted[0]), Port: uint16(port)}
	hasDesiredPeer := false
	for _, peer := range torrentPeers {
		if desiredPeer.Equal(&peer) {
			hasDesiredPeer = true
			break
		}
	}
	if !hasDesiredPeer {
		return fmt.Errorf("torrent file doesn't have a desired peer in peers list")
	}
	return nil
}

func dialPeer(peerSocketString string) (net.Conn, error) {
	conn, err := net.Dial("tcp", peerSocketString)
	if err != nil {
		return nil, fmt.Errorf("dialPeer error: %v", err)
	}
	return conn, nil
}

func parsePeerHandhakeResponse(b []byte) (*HandshakeMessage, error) {
	return nil, nil
}
