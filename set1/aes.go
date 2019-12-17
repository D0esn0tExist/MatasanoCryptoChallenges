package set1

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

// Ciph struct holds cipher properties
type Ciph struct {
	cipherText []byte
	isPadded   bool
	padLen     int
	cipherKey  []byte
	message    []byte
}

// Aesdecrypt function decrypts a file encrypted with AES in ECB mode.
func (c *Ciph) Aesdecrypt() string {
	c.cipherKey = []byte("YELLOW SUBMARINE")

	loaded := LoadFile("aes.txt")
	c.cipherText = make([]byte, len(loaded))
	base64.RawStdEncoding.Decode(c.cipherText, loaded)

	pad := len(c.cipherKey) - len(c.cipherText)%len(c.cipherKey)
	paddedCipher := make([]byte, len(c.cipherText)+pad)

	if pad > 0 {
		c.isPadded = true
		c.padLen = pad
		n := copy(paddedCipher, c.cipherText)
		for i := 0; i < pad; i++ {
			paddedCipher[n+i] = byte(pad)
		}
	}

	msgSize := len(paddedCipher)

	block := make([]byte, len(c.cipherKey))
	decoded := make([]byte, msgSize)

	for idx := 0; idx < msgSize; idx += len(c.cipherKey) {
		lim := idx + len(c.cipherKey)
		if lim > msgSize {
			lim = msgSize
		}
		block = paddedCipher[idx:lim]
		msg := DecryptAes(block, c.cipherKey)
		msg = append(msg, msg...)
	}
	if c.isPadded {
		decoded = decoded[:len(decoded)-c.padLen]
	}
	return string(decoded)
}

// DecryptAes function decrypts a block in ECB fashion
func DecryptAes(cipher, key []byte) []byte {
	msgBytes := make([]byte, len(key))
	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Decrypt(msgBytes, cipher)
	return msgBytes
}
