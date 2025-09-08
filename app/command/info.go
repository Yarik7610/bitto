package command

import (
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/app/bencode"
)

func (c controller) Info(fileName string) (map[string]any, error) {
	rawBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	dict, _, err := bencode.DecodeDict(string(rawBytes), 0)
	if err != nil {
		return nil, err
	}
	return dict, nil
}
