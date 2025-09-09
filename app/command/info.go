package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/bittorrent-starter-go/app/bencode"
	"github.com/codecrafters-io/bittorrent-starter-go/app/constants"
	"github.com/codecrafters-io/bittorrent-starter-go/app/utils"
)

type Torrent struct {
	TrackerURL  string
	Length      int64
	InfoHash    [constants.HASH_LENGTH]byte
	PieceLength int64
	Pieces      [][constants.HASH_LENGTH]byte
}

func (r *Torrent) String() string {
	var res strings.Builder

	res.WriteString(fmt.Sprintf("Tracker URL: %s\n", r.TrackerURL))
	res.WriteString(fmt.Sprintf("Length: %d\n", r.Length))
	res.WriteString(fmt.Sprintf("Info Hash: %x\n", r.InfoHash))
	res.WriteString(fmt.Sprintf("Piece Length: %d\n", r.PieceLength))
	res.WriteString("Piece Hashes:\n")
	for _, pieceHash := range r.Pieces {
		res.WriteString(fmt.Sprintf("%x\n", pieceHash))
	}

	return res.String()
}

func (c controller) Info(fileName string) (*Torrent, error) {
	rawBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	dict, _, err := bencode.DecodeDict(string(rawBytes), 0)
	if err != nil {
		return nil, err
	}

	tracker, ok := dict["announce"].(string)
	if !ok {
		return nil, fmt.Errorf("no announce field detected")
	}

	info, ok := dict["info"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("no info field detected")
	}

	length, ok := info["length"].(int64)
	if !ok {
		return nil, fmt.Errorf("no length field detected")
	}

	encodedInfo, err := bencode.Encode(info)
	if err != nil {
		return nil, err
	}

	hash, err := utils.GetHash([]byte(encodedInfo))
	if err != nil {
		return nil, err
	}

	pieceLength, ok := info["piece length"].(int64)
	if !ok {
		return nil, fmt.Errorf("no piece length field detected")
	}

	pieces, ok := info["pieces"].(string)
	if !ok {
		return nil, fmt.Errorf("no pieces field detected")
	}
	pieceHashes, err := splitPieceHashes([]byte(pieces))
	if err != nil {
		return nil, err
	}

	response := &Torrent{
		TrackerURL:  tracker,
		Length:      length,
		InfoHash:    hash,
		PieceLength: pieceLength,
		Pieces:      pieceHashes,
	}
	return response, nil
}

func splitPieceHashes(b []byte) ([][constants.HASH_LENGTH]byte, error) {
	l := len(b)
	if l%constants.HASH_LENGTH != 0 {
		return nil, fmt.Errorf("splitPieceHashes: slice length %d isn't a multiple of %d", l, constants.HASH_LENGTH)
	}

	piecesCount := l / constants.HASH_LENGTH
	res := make([][constants.HASH_LENGTH]byte, 0, piecesCount)
	for i := 0; i < l; i += constants.HASH_LENGTH {
		var hash [constants.HASH_LENGTH]byte
		copy(hash[:], b[i:i+constants.HASH_LENGTH])
		res = append(res, hash)
	}
	return res, nil
}
