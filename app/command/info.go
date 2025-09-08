package command

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/app/bencode"
	"github.com/codecrafters-io/bittorrent-starter-go/app/utils"
)

type InfoResponse struct {
	Tracker string
	Length  int64
	Hash    []byte
}

func (c controller) Info(fileName string) (*InfoResponse, error) {
	rawBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	dict, _, err := bencode.DecodeDict(string(rawBytes), 0)
	if err != nil {
		return nil, err
	}

	tracker, ok := dict["announce"].(string)
	if !ok {
		return nil, fmt.Errorf("no announce field detected")
	}

	info, ok := dict["info"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("no info field detected")
	}

	length, ok := info["length"].(int64)
	if !ok {
		return nil, fmt.Errorf("no length field detected")
	}

	encodedInfo, err := bencode.Encode(info)
	if err != nil {
		return nil, err
	}

	hash, err := utils.GetHash([]byte(encodedInfo))
	if err != nil {
		return nil, err
	}

	response := &InfoResponse{
		Tracker: tracker,
		Length:  length,
		Hash:    hash,
	}
	return response, nil
}
