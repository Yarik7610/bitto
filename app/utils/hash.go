package utils

import (
	"crypto/sha1"
	"fmt"

	"github.com/codecrafters-io/bittorrent-starter-go/app/constants"
)

func GetHash(b []byte) ([constants.HASH_LENGTH]byte, error) {
	s := sha1.New()
	_, err := s.Write(b)
	if err != nil {
		return [constants.HASH_LENGTH]byte{}, fmt.Errorf("GetHash error: %v", err)
	}

	var hash [constants.HASH_LENGTH]byte
	copy(hash[:], s.Sum(nil))
	return hash, nil
}
