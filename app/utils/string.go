package utils

import "fmt"

func SplitEachNBytes(s []byte, n int) ([][]byte, error) {
	l := len(s)
	if l%n != 0 {
		return nil, fmt.Errorf("SplitEachNBytes: slice length %d isn't a multiple of %d", l, n)
	}

	res := make([][]byte, 0)
	for i := 0; i < len(s); i += n {
		res = append(res, s[i:i+n])
	}
	return res, nil
}
