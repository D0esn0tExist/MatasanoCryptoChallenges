package set2

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

var encKey string

func init() {
	encKey = KeyGen(16)
}

// AES128ECBoracle encrypts buffers under ECB mode using a consistent but unknown key
func AES128ECBoracle(b bytes.Buffer) string {
	pad := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	decodedPad, _ := base64.RawStdEncoding.DecodeString(pad)

	_, err := b.Write(decodedPad)
	if err != nil {
		panic(err)
	}

	fmt.Println("Padded: ", b.Len())
	// Do AES ECB
	c := &set1.Ciph{CipherText: nil, CipherKey: []byte(encKey), Message: b.Bytes()}
	cipher := c.Aesencrypt()
	return base64.RawStdEncoding.EncodeToString(cipher)

}
