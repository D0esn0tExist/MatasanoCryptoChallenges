package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sort"
)

type keyProp struct {
	key     string
	keysize int
	dist    float64
}

// BreakRepeat function attempts to get repeating key used to XOR encrypt resulting cipher.
func BreakRepeat(path string) {
	ciphFile := string(LoadFile(path))
	ciphBytes, err := base64.StdEncoding.DecodeString(ciphFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	possibleKeys := possibleKeySize(ciphBytes)

	sort.Slice(possibleKeys, func(i, j int) bool {
		return possibleKeys[i].dist < possibleKeys[j].dist
	})

	// Chunk up cipher to blocks of keysize length
	for i := 0; i < 4; i++ {
		blockindex := 0
		blocks := make([][]byte, len(ciphBytes)/possibleKeys[i].keysize+1)

		for idx := 0; idx < len(ciphBytes); idx += possibleKeys[i].keysize {
			block := ciphBytes[idx : idx+possibleKeys[i].keysize]
			blocks[blockindex] = block
			blockindex++
		}

		// Transpose blocks
		transposed := make([][]byte, possibleKeys[i].keysize)
		for i := range transposed {
			transposed[i] = make([]byte, len(blocks))
		}
		for k := 0; k < possibleKeys[i].keysize; k++ {
			for j := 0; j < len(blocks); j++ {
				transposed[k][j] = blocks[j][k]
			}
		}
		key := ""
		for _, v := range transposed {
			breakScore := SingleXor(v)
			key += breakScore.key
		}
		fmt.Println(key)
	}
}

// HemmingDistance function calculates difference between two strings
func HemmingDistance(s1, s2 []byte) (float64, error) {
	if len(s1) != len(s2) {
		return -1, errors.New("Two strings not the same length")
	}

	hdist := 0.0
	for i := 0; i < len(s1); i++ {
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (s1[i] & mask) != (s2[i] & mask) {
				hdist++
			}
		}
	}
	return hdist, nil
}

// PossibleKeySize function ranks possible key sizes blah!
func possibleKeySize(cipher []byte) []keyProp {
	keysize := 3
	keyndex := 0
	possibleKeys := make([]keyProp, 37)
	for i := keysize; i < 40; i++ {
		possibleKeys[keyndex].keysize = i
		dist, _ := HemmingDistance(cipher[0:i], cipher[i:2*i])
		dist1, _ := HemmingDistance(cipher[2*i:3*i], cipher[3*i:4*i])
		dist2, _ := HemmingDistance(cipher[3*i:4*i], cipher[5*i:6*i])
		dist = dist / float64(i)
		dist1 = dist1 / float64(i)
		dist2 = dist2 / float64(i)
		dist = (dist + dist1 + dist2) / float64(2)
		possibleKeys[keyndex].dist = dist
		keyndex++
	}
	return possibleKeys
}
