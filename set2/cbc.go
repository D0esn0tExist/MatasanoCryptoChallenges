package set2

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

// CbcProp struct has the encryption properties
type CbcProp struct {
	IV  []byte
	key []byte
}

// Cbc implements AES in ECB mode.
func (c CbcProp) Cbc(path string) ([]byte, error) {
	cbcCiph := set1.LoadFile(path)
	cbcBytes := make([]byte, len(cbcCiph))
	base64.RawStdEncoding.Decode(cbcBytes, cbcCiph)
	fmt.Println(len(cbcBytes))

	paddedBytes := PKCSPadding(cbcBytes, aes.BlockSize)
	fmt.Println(len(paddedBytes))

	cipher := make([]byte, aes.BlockSize)
	finalCipher := make([]byte, len(paddedBytes))
	ciphBlock := c.IV

	for idx := 0; idx < len(paddedBytes); idx += aes.BlockSize {
		lim := idx + aes.BlockSize

		ciphBlock = set1.Xor(paddedBytes[idx:lim], ciphBlock)
		cipher = set1.EncryptAes(ciphBlock, c.key)

		copy(finalCipher[idx:], cipher)
		finalCipher = append(finalCipher, cipher...)
	}
	return finalCipher, nil
}
