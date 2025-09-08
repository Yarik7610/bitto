package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected OK status, received: %d", resp.StatusCode)
	}

	var response bytes.Buffer
	_, err = io.Copy(&response, resp.Body)
	if err != nil {
		return nil, err
	}

	return response.Bytes(), err
}
