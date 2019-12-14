package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// LoadFile returns []byte content of a file
func LoadFile(path string) []byte {
	ciphFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return ciphFile
}

// Detect function loads and goes through a txt file to find an XOR'ed string
func Detect() CipherScore {
	ciphFile := LoadFile("data.txt")
	ciphStrings := strings.Split(string(ciphFile), "\n")

	stringScores := make([]CipherScore, len(ciphStrings))

	for i, s := range ciphStrings {
		c, err := hex.DecodeString(s)
		if err != nil {
			fmt.Println(err.Error())
		}
		stringScores[i] = SingleXor(c)
	}

	sort.Slice(stringScores, func(i, j int) bool {
		return stringScores[i].score > stringScores[j].score
	})
	return stringScores[0]

}
