package bencode

import (
	"fmt"
	"strconv"
	"unicode"
)

func DecodeString(s string, idx int) (string, int, error) {
	l := len(s)

	if ok := unicode.IsDigit(rune(s[idx])); !ok {
		return "", 0, fmt.Errorf("method DecodeString error: no bencoded string len detected")
	}

	colonIdx := -1
	for i := idx + 1; i < l; i++ {
		if s[i] == ':' {
			colonIdx = i
			break
		}
	}
	if colonIdx == -1 {
		return "", 0, fmt.Errorf("DecodeString error: missing ':' delimiter")
	}

	requestedLen, err := strconv.Atoi(s[idx:colonIdx])
	if err != nil {
		return "", 0, fmt.Errorf("DecodeString error: %v", err)
	}

	if colonIdx+1+requestedLen > l {
		return "", idx, fmt.Errorf("DecodeString error: insufficient length, requested %d, got %d", requestedLen, l-(colonIdx+1))
	}

	return s[colonIdx+1 : colonIdx+1+requestedLen], colonIdx + 1 + requestedLen, nil
}

func EncodeString(s string) string {
	return fmt.Sprintf("%d:%s", len(s), s)
}
