package set1

import (
	"encoding/hex"
)

// Xor takes two equal-length buffers and produces their XOR combination.
func Xor(input, xoree []byte) string {
	res := make([]byte, len(input))

	for i := range input {
		res[i] = input[i] ^ xoree[i]
	}

	return hex.EncodeToString(res)
}
