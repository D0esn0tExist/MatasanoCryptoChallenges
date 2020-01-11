package set1

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"strings"
)

// DetectAes function goes through a bunch of hex encoded ciphertexts in a text file
// and finds the one encrypted with ECB
func DetectAes(path string) []string {
	aesFile := string(LoadFile(path))
	aesStrings := strings.Split(aesFile, "\n")
	found := make([]string, 0)
	for i := 0; i < len(aesStrings); i++ {
		byteCipher, _ := hex.DecodeString(aesStrings[i])
		if DetectAesCipher(byteCipher) == true {
			fmt.Printf("Found: %s\n", aesStrings[i])
			found = append(found, aesStrings[i])
		}
	}
	return found
}

// DetectAesCipher checks if a byte array is an ECB encrypted cipher
func DetectAesCipher(cipher []byte) bool {
	if len(cipher)%aes.BlockSize != 0 {
		fmt.Println("nOT A vALID ciPHER!\n Length not multiple of blocksize.")
		return false
	}
	dups := make(map[string]int, len(cipher)/aes.BlockSize)
	for i := 0; i < len(cipher); i += aes.BlockSize {
		if count, ok := dups[string(cipher[i:i+aes.BlockSize])]; ok {
			count++
			return true
		}
		dups[string(cipher[i:i+aes.BlockSize])] = 0
	}
	return false
}
