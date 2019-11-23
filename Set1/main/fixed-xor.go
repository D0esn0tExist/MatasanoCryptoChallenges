package main

import (
	"encoding/hex"
	"fmt"
)

// Xor takes two equal-length buffers and produces their XOR combination.
func Xor(input []byte) string {
	xoree, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	fmt.Println(xoree)
	res := make([]byte, len(input))

	for i := range input {
		res[i] = input[i] ^ xoree[i]
	}

	return hex.EncodeToString(res)
}
