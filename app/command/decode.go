package command

import (
	"encoding/json"

	"github.com/codecrafters-io/bittorrent-starter-go/app/bencode"
)

func (c controller) Decode(encoded string) ([]byte, error) {
	decoded, err := bencode.Decode(encoded)
	if err != nil {
		return nil, err
	}

	jsonOutput, err := json.Marshal(decoded)
	if err != nil {
		return nil, err
	}
	return jsonOutput, nil
}
