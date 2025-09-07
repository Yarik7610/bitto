package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/codecrafters-io/bittorrent-starter-go/app/bencode"
)

func main() {
	command := os.Args[1]

	if command == "decode" {
		bencodedValue := os.Args[2]

		decoded, err := bencode.DecodeString(bencodedValue)
		if err != nil {
			log.Println(err)
			return
		}

		jsonOutput, _ := json.Marshal(decoded)
		fmt.Println(string(jsonOutput))
	} else {
		log.Fatalf("Unknown command: %s", command)
	}
}
