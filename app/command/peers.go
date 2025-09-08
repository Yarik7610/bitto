package command

import (
	"encoding/binary"
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/codecrafters-io/bittorrent-starter-go/app/bencode"
	"github.com/codecrafters-io/bittorrent-starter-go/app/utils"
)

const (
	PEER_ID      = "lHavHuZBkaYWXsuvjGJh"
	PEER_PORT    = 6881
	COMPACT_MODE = 1
)

type PeerSocket struct {
	IP   net.IP
	Port uint16
}

func (s PeerSocket) String() string {
	return fmt.Sprintf("%s:%d", s.IP, s.Port)
}

func (c controller) Peers(fileName string) ([]PeerSocket, error) {
	info, err := c.Info(fileName)
	if err != nil {
		return nil, fmt.Errorf("info command error: %v", err)
	}

	baseURL, err := url.Parse(info.TrackerURL)
	if err != nil {
		return nil, err
	}
	queryParams := createPeersQueryParams(info)
	baseURL.RawQuery = queryParams.Encode()

	b, err := utils.Get(baseURL.String())
	if err != nil {
		return nil, err
	}

	decoded, err := bencode.Decode(string(b))
	if err != nil {
		return nil, err
	}

	response, ok := decoded.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("wrong response type, need dict bencode type")
	}

	peersStr, ok := response["peers"].(string)
	if !ok {
		return nil, fmt.Errorf("no peers field found in response")
	}

	peers, err := readPeers([]byte(peersStr))
	if err != nil {
		return nil, err
	}

	return peers, nil
}

func createPeersQueryParams(info *InfoResponse) *url.Values {
	params := url.Values{}

	params.Add("info_hash", string(info.InfoHash))
	params.Add("peer_id", PEER_ID)
	params.Add("port", strconv.Itoa(PEER_PORT))

	params.Add("uploaded", strconv.Itoa(0))
	params.Add("downloaded", strconv.Itoa(0))
	params.Add("left", strconv.FormatInt(info.Length, 10))

	params.Add("compact", strconv.Itoa(COMPACT_MODE))

	return &params
}

func readPeers(b []byte) ([]PeerSocket, error) {
	l := len(b)

	const IP_ADDRESS_BYTES_COUNT = 4
	const PORT_BYTES_COUNT = 2
	const SOCKET_BYTES_COUNT = IP_ADDRESS_BYTES_COUNT + PORT_BYTES_COUNT

	if l%(SOCKET_BYTES_COUNT) != 0 {
		return nil, fmt.Errorf("readPeers: slice length %d isn't a multiple of %d", l, SOCKET_BYTES_COUNT)
	}

	peers := make([]PeerSocket, 0)
	for i := 0; i < len(b); i += SOCKET_BYTES_COUNT {
		IP := b[i : i+IP_ADDRESS_BYTES_COUNT]
		port := binary.BigEndian.Uint16(b[i+IP_ADDRESS_BYTES_COUNT : i+SOCKET_BYTES_COUNT])
		peers = append(peers, PeerSocket{IP: IP, Port: port})
	}
	return peers, nil
}
