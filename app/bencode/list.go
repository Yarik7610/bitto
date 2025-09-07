package bencode

import "fmt"

func DecodeList(s string, idx int) ([]any, int, error) {
	l := len(s)

	if idx >= l || s[idx] != 'l' {
		return nil, 0, fmt.Errorf("DecodeList error: no starting delimiter 'l' detected")
	}

	i := idx + 1
	res := make([]any, 0)
	for {
		if i >= l {
			return nil, 0, fmt.Errorf("DecodeList error: unterminated list, missing 'e'")
		}
		if s[i] == 'e' {
			return res, i + 1, nil
		}

		val, newIdx, err := decode(s, i)
		if err != nil {
			return nil, 0, err
		}
		res = append(res, val)
		i = newIdx
	}
}
