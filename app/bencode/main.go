package bencode

import (
	"fmt"
	"unicode"
)

func Decode(s string) (any, error) {
	val, _, err := decode(s, 0)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func decode(s string, idx int) (any, int, error) {
	if idx >= len(s) {
		return nil, 0, fmt.Errorf("decode error: index %d out of range", idx)
	}

	if ok := unicode.IsDigit(rune(s[idx])); ok {
		return DecodeString(s, idx)
	} else if s[idx] == 'i' {
		return DecodeInt(s, idx)
	} else if s[idx] == 'l' {
		return DecodeList(s, idx)
	} else if s[idx] == 'd' {
		return DecodeDict(s, idx)
	} else {
		return nil, 0, fmt.Errorf("decode error: unsupported bencoded type detected")
	}
}
