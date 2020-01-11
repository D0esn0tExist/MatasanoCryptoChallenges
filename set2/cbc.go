package set2

import (
	"crypto/aes"
	"fmt"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

// CbcProp struct has the encryption properties
type CbcProp struct {
	IV  []byte
	key []byte
}

// Cbc implements AES in ECB mode.
func (c CbcProp) Cbc(cbcBytes []byte) ([]byte, error) {
	ciph, err := aes.NewCipher(c.key)
	if err != nil {
		fmt.Println(err.Error())
	}
	paddedBytes := set1.PKCSPadding(cbcBytes, aes.BlockSize)

	cipher := make([]byte, len(paddedBytes))
	ciphBlock := c.IV

	for idx := 0; idx < len(paddedBytes); idx += aes.BlockSize {
		lim := idx + aes.BlockSize

		ciphBlock = set1.Xor(paddedBytes[idx:lim], ciphBlock)
		ciph.Encrypt(cipher[idx:lim], ciphBlock)
	}
	finalCipher := set1.PKCSUnpadding(cipher)
	return finalCipher, nil
}
