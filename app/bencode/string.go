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

	colonIdx := -1
	for i := 1; i < len(s); i++ {
		if s[i] == ':' {
			colonIdx = i
			break
		}
	}
	if colonIdx == -1 {
		return "", fmt.Errorf("DecodeString error: missing ':' delimiter")
	}

	requestedLen, err := strconv.Atoi(s[:colonIdx])
	if err != nil {
		return "", fmt.Errorf("DecodeString error: %v", err)
	}

	realLen := len(s) - colonIdx - 1
	if requestedLen != realLen {
		return "", fmt.Errorf("DecodeString error: wrong requested length of string, requsted: %d, got %d", requestedLen, realLen)
	}

	return s[colonIdx+1 : colonIdx+1+requestedLen], nil
}

func EncodeString(s string) string {
	return fmt.Sprintf("%d:%s", len(s), s)
}
