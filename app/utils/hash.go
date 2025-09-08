package utils

import (
	"crypto/sha1"
	"fmt"
)

func GetHash(b []byte) ([]byte, error) {
	s := sha1.New()
	_, err := s.Write(b)
	if err != nil {
		return nil, fmt.Errorf("GetHash error: %v", err)
	}
	return s.Sum(nil), nil
}
