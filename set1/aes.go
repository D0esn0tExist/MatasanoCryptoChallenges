package set1

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

// Ciph struct holds cipher properties
type Ciph struct {
	CipherText []byte
	CipherKey  []byte
	Message    []byte
}

// Aesencrypt function handles enc and decryption of a plaintext in ECB mode
func (c *Ciph) Aesencrypt() []byte {
	ciph, err := aes.NewCipher(c.CipherKey)
	if err != nil {
		fmt.Println(err.Error())
	}
	paddedInput := PKCSPadding(c.Message, aes.BlockSize)

	block := make([]byte, aes.BlockSize)
	cipher := make([]byte, len(paddedInput))

	for idx := 0; idx < len(paddedInput); idx += aes.BlockSize {
		lim := idx + aes.BlockSize
		if lim > len(paddedInput) {
			lim = len(paddedInput)
		}
		block = paddedInput[idx:lim]
		ciph.Encrypt(cipher[idx:lim], block)
	}
	fmt.Println("Cip: ", len(paddedInput))

	return cipher
}

// Aesdecrypt function decrypts a file encrypted with AES in ECB mode.
func (c *Ciph) Aesdecrypt() string {
	c.CipherKey = []byte("YELLOW SUBMARINE")

	loaded := LoadFile("aes.txt")
	c.CipherText = make([]byte, len(loaded))
	base64.RawStdEncoding.Decode(c.CipherText, loaded)

	paddedCipher := PKCSPadding(c.CipherText, 16)
	cipher, err := aes.NewCipher(c.CipherKey)
	if err != nil {
		fmt.Println(err.Error())
	}

	msgSize := len(paddedCipher)
	print(msgSize)

	block := make([]byte, len(c.CipherKey))
	decoded := make([]byte, msgSize)

	for idx := 0; idx < msgSize; idx += len(c.CipherKey) {
		lim := idx + len(c.CipherKey)
		if lim > msgSize {
			lim = msgSize
		}
		block = paddedCipher[idx:lim]
		cipher.Decrypt(decoded[idx:lim], block)
	}
	decoded = PKCSUnpadding(decoded)

	return string(decoded)
}
