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
		return DecodeInt64(s, idx)
	} else if s[idx] == 'l' {
		return DecodeList(s, idx)
	} else if s[idx] == 'd' {
		return DecodeDict(s, idx)
	} else {
		return nil, 0, fmt.Errorf("decode error: unsupported bencode type detected")
	}
}

func Encode(v any) (string, error) {
	return encode(v)
}

func encode(v any) (string, error) {
	switch v := v.(type) {
	case string:
		return EncodeString(v)
	case int64:
		return EncodeInt64(v)
	case []any:
		return EncodeList(v)
	case map[string]any:
		return EncodeDict(v)
	default:
		return "", fmt.Errorf("encode error: unsupported bencode type detected")
	}
}
