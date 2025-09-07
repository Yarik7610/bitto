package bencode

import "fmt"

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
