package bencode

import (
	"fmt"
	"strconv"
)

const ZERO = "i0e"

func DecodeInt(s string) (int64, error) {
	if s == ZERO {
		return 0, nil
	}

	l := len(s)
	if l < 3 {
		return 0, fmt.Errorf("DecodeInt error: too small string provided, at least 3 chars required")
	}
	if s[0] != 'i' {
		return 0, fmt.Errorf("DecodeInt error: no starting delimiter 'i' detected")
	}
	if s[l-1] != 'e' {
		return 0, fmt.Errorf("DecodeInt error: no ending delimiter 'e' detected")
	}

	valStr := s[1 : l-1]
	if len(valStr) > 1 {
		if valStr[0] == '-' && valStr[1] == '0' {
			return 0, fmt.Errorf("DecodeInt error: negative zero detected")
		}
		if valStr[0] == '0' {
			return 0, fmt.Errorf("DecodeInt error: leading zero detected")
		}
	}

	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("DecodeInt error: %v", err)
	}
	return val, nil
}

func EncodeInt(val int64) string {
	return fmt.Sprintf("i%de", val)
}
