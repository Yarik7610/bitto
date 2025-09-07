package bencode

import (
	"fmt"
	"unicode"
)

func Decode(s string) (any, error) {
	l := len(s)

	if ok := unicode.IsDigit(rune(s[0])); ok {
		return DecodeString(s)
	} else if s[0] == 'i' && s[l-1] == 'e' {
		return DecodeInt(s)
	} else {
		return nil, fmt.Errorf("Decode error: unsupported bencoded type detected")
	}
}
