package bencode

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/bittorrent-starter-go/app/utils"
)

func DecodeDict(s string, idx int) (map[string]any, int, error) {
	l := len(s)

	if idx >= l || s[idx] != 'd' {
		return nil, 0, fmt.Errorf("DecodeDict error: no starting delimiter 'd' detected")
	}

	i := idx + 1
	dict := make(map[string]any)
	for {
		if i >= l {
			return nil, 0, fmt.Errorf("DecodeDict error: unterminated list, missing 'e'")
		}
		if s[i] == 'e' {
			return dict, i + 1, nil
		}

		key, nextIdx, err := DecodeString(s, i)
		if err != nil {
			return nil, 0, fmt.Errorf("DecodeDict key parse error: %v", err)
		}
		val, newIdx, err := decode(s, nextIdx)
		if err != nil {
			return nil, 0, fmt.Errorf("DecodeDict key parse error: %v", err)
		}
		dict[key] = val
		i = newIdx
	}
}

func EncodeDict(dict map[string]any) (string, error) {
	var res strings.Builder
	res.WriteRune('d')

	sortedKeys := utils.SortMapKeys(dict)

	for _, key := range sortedKeys {
		encodedKey, err := EncodeString(key)
		if err != nil {
			return "", err
		}
		res.WriteString(encodedKey)

		encodedVal, err := encode(dict[key])
		if err != nil {
			return "", err
		}
		res.WriteString(encodedVal)
	}

	res.WriteRune('e')
	return res.String(), nil
}
