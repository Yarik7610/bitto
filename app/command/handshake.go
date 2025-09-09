package command

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/codecrafters-io/bittorrent-starter-go/app/constants"
	"github.com/codecrafters-io/bittorrent-starter-go/app/utils"
)

type HandshakeMessage struct {
	ProtocolString string
	PeerID         string
	*Torrent
}

func (m *HandshakeMessage) Bytes() []byte {
	var res bytes.Buffer

	res.WriteByte(byte(len(m.ProtocolString)))
	res.WriteString(m.ProtocolString)
	res.Write(make([]byte, 8))
	res.Write(m.InfoHash[:])
	res.WriteString(m.PeerID)

	return res.Bytes()
}

func (c controller) Handshake(fileName string, peerSocketString string) (*HandshakeMessage, error) {
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
		ProtocolString: constants.PROTOCOL_STRING,
		PeerID:         c.client.PeerID,
		Torrent:        torrent,
	}

	_, err = conn.Write(message.Bytes())
	if err != nil {
		return nil, fmt.Errorf("write to peer error: %v", err)
	}

	b := make([]byte, 1024)
	n, err := conn.Read(b)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("read from peer error: %v", err)
	}

	response, err := parsePeerHandhakeResponse(torrent, b[:n])
	if err != nil {
		return nil, err
	}
	return response, err
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

	conn, err := net.DialTimeout("tcp", peerSocketString, time.Second*constants.CONN_OPERATIONS_TIME_SECONDS)
	if err != nil {
		return nil, fmt.Errorf("dialPeer error: %v", err)
	}
	conn.SetDeadline(time.Now().Add(time.Second * constants.CONN_OPERATIONS_TIME_SECONDS))

	return conn, nil
}

func parsePeerHandhakeResponse(torrent *Torrent, b []byte) (*HandshakeMessage, error) {
	protocolLen := b[0]
	protocol := (b[1:protocolLen])
	b = b[1+protocolLen:]

	b = b[8:]

	var infoHash [constants.HASH_LENGTH]byte
	copy(infoHash[:], b[:20])
	if !bytes.Equal(infoHash[:], torrent.InfoHash[:]) {
		return nil, fmt.Errorf("wrong info hashes detected while parsing another peer handshake response")
	}
	b = b[20:]

	peerID := b[:20]

	response := &HandshakeMessage{
		ProtocolString: string(protocol),
		PeerID:         string(peerID),
		Torrent: &Torrent{
			InfoHash: infoHash,
		},
	}
	return response, nil
}
