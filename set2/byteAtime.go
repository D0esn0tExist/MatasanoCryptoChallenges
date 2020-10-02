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
var prefix []byte

func init() {
	// generate random key
	rand.Seed(time.Now().UnixNano())
	encKey = KeyGen(16)
	// generate random prefix
	prefix = make([]byte, rand.Intn(10))
	_, err := rand.Read(prefix)
	if err != nil {
		panic(err)
	}
	fmt.Printf("---------Prefix length: %d\n", len(prefix))
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

// BreakSuffixOracleLength function attempts to break AES128Suffix oracle to find the length of the suffix pad.
func BreakSuffixOracleLength(b bytes.Buffer) (blockSize, suffixLen int) {
	cipher, _ := base64.RawStdEncoding.DecodeString(AES128ECBSuffixoracle(b))
	cipherLen := len(cipher)
	initLen := cipherLen
	for cipherLen == initLen {
		b.WriteString("A")
		cipher, _ := base64.RawStdEncoding.DecodeString(AES128ECBSuffixoracle(b))
		cipherLen = len(cipher)
		fmt.Printf("With %s, len: %d\n", string(b.Bytes()), cipherLen)
	}
	blockSize = cipherLen - initLen
	suffixLen = initLen - (b.Len() - 1)
	return
}

// FindUnknownSuffixPad function attempts to break byte-at-a-time oracle to find the unknown suffix string padded.
func FindUnknownSuffixPad(prefix []byte, unknownStringSize, blockSize int) string {
	unknown := []byte{}
	for i := 0; i < unknownStringSize+len(prefix); i++ { // TODO: Dirty solution. Fix this.
		// input text to determine byte of the unknown string padding
		fix := strings.Repeat("A", blockSize-1-(len(unknown)%blockSize))
		b := bytes.Buffer{}
		// append prefix bytes if present.
		if prefix != nil {
			b.Write(prefix)
		}
		b.Write([]byte(fix))
		check, _ := base64.RawStdEncoding.DecodeString(AES128ECBSuffixoracle(b))
		start := 16 * (len(unknown) / 16)
		check = check[start : start+16]

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
	return string(unknown[len(prefix):])

}

// AES128ECBPrefixoracle function generates a random count of random bytes and prepend this string to AES128ECBSuffixoracle suffix pad.
func AES128ECBPrefixoracle(input []byte) string {
	b := bytes.Buffer{}
	b.Write(prefix)
	b.Write(input)
	return AES128ECBSuffixoracle(b)
}

// BreakPrefixOracleLength function attempts to break AES128PrefixLength oracle to find the length of the prefix pad.
func BreakPrefixOracleLength(blockSize int) int {
	// to figure out the size of the prefix oracle, we have to add a padding before the input of
	// atleast 3 blocks(16 * 3) of repeating chars.
	padInput := bytes.Repeat([]byte{'A'}, blockSize*3)
	prefixSize := findPrefixSize(padInput, blockSize)
	return prefixSize
}

// findPrefixSize function determines how long the prefix string is.
func findPrefixSize(input []byte, blockSize int) int {
	for i := 0; i < blockSize; i++ {
		padding := make([]byte, i)
		plainText := append(padding, input...)

		// - keep prepending a padding byte until we find a n block long cipher text
		cipherText := AES128ECBPrefixoracle(plainText)
		cipher, _ := base64.RawStdEncoding.DecodeString(cipherText)

		blockText, location := findRepeatingBlock(cipher, blockSize, len(input)/blockSize)
		if location != -1 {
			// - change the attack text content (but not the prefix padding) and confirm that
			//   we get a different n block long cipher text AT THE SAME LOCATION
			cipher = bytes.Repeat([]byte{'B'}, len(cipher))
			plainText := append(padding, cipher...)
			cipherText := AES128ECBPrefixoracle(plainText)
			cipher, _ := base64.RawStdEncoding.DecodeString(cipherText)
			newBlock, newLocation := findRepeatingBlock(cipher, blockSize, len(input)/blockSize)
			if newLocation == location && !bytes.Equal(blockText, newBlock) {
				return location*blockSize - i
			}
		}

		// - repeat until we know the prefix size
	}
	return 0
}

// findRepeatingBlock returns the first repeating block of count blocks
// of blockSize, or -1 if there isn't one.
func findRepeatingBlock(buf []byte, blockSize int, count int) (content []byte, location int) {
	if len(buf)%blockSize != 0 {
		panic("Need multiple of block size")
	}

	location = -1
	totalBlocks := len(buf) / blockSize

	var previous []byte
	seen := 0

	for i := 0; i < totalBlocks; i++ {
		start := i * blockSize
		end := start + blockSize
		chunk := buf[start:end]
		if bytes.Equal(previous, chunk) {
			seen++
			if seen == count {
				content = chunk
				location = i + 1 - seen
				break
			}
		} else {
			seen = 1
		}
		previous = chunk
	}
	return
}
