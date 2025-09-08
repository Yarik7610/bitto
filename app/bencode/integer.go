package bencode

import (
	"fmt"
	"strconv"
)

func DecodeInt64(s string, idx int) (int64, int, error) {
	l := len(s)

	if s[idx] != 'i' {
		return 0, 0, fmt.Errorf("DecodeInt64 error: no starting delimiter 'i' detected")
	}

	numEndIdx := idx + 1
	for ; numEndIdx < l && s[numEndIdx] != 'e'; numEndIdx++ {
	}

	valStr := s[idx+1 : numEndIdx]
	if len(valStr) > 1 {
		if valStr[0] == '-' && valStr[1] == '0' {
			return 0, 0, fmt.Errorf("DecodeInt64 error: negative zero detected")
		}
		if valStr[0] == '0' {
			return 0, 0, fmt.Errorf("DecodeInt64 error: leading zero detected")
		}
	}

	val, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("DecodeInt64 error: %v", err)
	}
	return val, numEndIdx + 1, nil
}

func EncodeInt64(num int64) (string, error) {
	return fmt.Sprintf("i%de", num), nil
}
