package set1

import (
	"bytes"
	"encoding/hex"
)

// RepeatingXor function encodes a string with key using repating XOR.
func RepeatingXor(message, key string) string {
	buf := &bytes.Buffer{}
	msgBytes := []byte(message)
	keyBytes := []byte(key)

	mult := len(msgBytes) / len(keyBytes)
	pad := len(msgBytes) % len(keyBytes)

	keyBytes = bytes.Repeat(keyBytes, mult)
	buf.Write(keyBytes)

	if pad > 0 {
		for i := 0; i < pad; i++ {
			buf.WriteByte(keyBytes[i])
		}
	}

	Xor(msgBytes, buf.Bytes())
	return hex.EncodeToString(msgBytes)
}
