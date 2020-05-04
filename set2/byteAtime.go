package set2

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

var encKey string

func init() {
	rand.Seed(time.Now().UnixNano())
	encKey = KeyGen(16)
}

// AES128ECBSuffixoracle encrypts buffers under ECB mode using a consistent but unknown key
func AES128ECBSuffixoracle(b bytes.Buffer) string {
	pad := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"
	decodedPad, _ := base64.RawStdEncoding.DecodeString(pad)
	_, err := b.Write(decodedPad)
	if err != nil {
		panic(err)
	}
	// Do AES ECB
	c := &set1.Ciph{CipherText: nil, CipherKey: []byte(encKey), Message: b.Bytes()}
	cipher := c.Aesencrypt()
	return base64.RawStdEncoding.EncodeToString(cipher)

}

// FindUnknownString function attempts to break byte-at-a-time oracle to find the unknown suffix string padded.
func FindUnknownString(unknownStringSize, blockSize int) string {
	unknown := []byte{}
	for i := 0; i < unknownStringSize; i++ {
		// input text to determine byte of the unknown string padding
		fix := strings.Repeat("A", blockSize-1-(len(unknown)%blockSize))
		b := bytes.Buffer{}
		b.Write([]byte(fix))
		check, _ := base64.RawStdEncoding.DecodeString(AES128ECBSuffixoracle(b))
		start := 16 * (len(unknown) / 16)
		check = check[start : start+16]
		fmt.Println("-----------------------\n For check: ", check)

		// Prepare brute msg to input
		bruteMsg := bytes.Repeat([]byte{'A'}, blockSize)
		bruteMsg = append(bruteMsg, unknown...)
		bruteMsg = append(bruteMsg, 'X')
		bruteMsg = bruteMsg[len(bruteMsg)-blockSize:]

		// Loop through possible character bytes
		for l := 0; l < 256; l++ {
			bruteMsg[len(bruteMsg)-1] = byte(l)
			loopvar := bytes.Buffer{}
			loopvar.Write(bruteMsg)
			cipherText, _ := base64.RawStdEncoding.DecodeString(AES128ECBSuffixoracle(loopvar))
			if bytes.Equal(check, cipherText[:blockSize]) {
				unknown = append(unknown, byte(l))
				break
			}
		}
	}
	return string(unknown)

}

// AES128ECBPrefixoracle function generates a random count of random bytes and prepend this string to AES128ECBSuffixoracle suffix pad.
func AES128ECBPrefixoracle(input []byte) string {
	prefix := make([]byte, rand.Intn(10))
	_, err := rand.Read(prefix)
	if err != nil {
		fmt.Println(fmt.Errorf(err.Error()))
	}
	b := bytes.Buffer{}
	b.Write(prefix)
	b.Write(input)
	return AES128ECBSuffixoracle(b)
}
