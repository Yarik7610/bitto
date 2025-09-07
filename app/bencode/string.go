package bencode

import (
	"fmt"
	"strconv"
	"unicode"
)

func DecodeString(s string) (string, error) {
	if ok := unicode.IsDigit(rune(s[0])); !ok {
		return "", fmt.Errorf("method DecodeString error: no bencoded string len detected")
	}

	var colonIdx int
	for i := 1; i < len(s); i++ {
		if s[i] == ':' {
			colonIdx = i
			break
		}
	}

	requestedLen, err := strconv.Atoi(s[:colonIdx])
	if err != nil {
		return "", fmt.Errorf("method DecodeString error: %v", err)
	}

	realLen := len(s) - colonIdx - 1
	if requestedLen > realLen {
		return "", fmt.Errorf("method DecodeString error: wrong requested length of string, requsted: %d, got %d", requestedLen, realLen)
	}

	return s[colonIdx+1 : colonIdx+1+requestedLen], nil
}
