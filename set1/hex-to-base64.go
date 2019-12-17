package set1

import (
	b64 "encoding/base64"
	"encoding/hex"
	"log"
)

// Hextobase64 converts hex string to base64 string
func Hextobase64(hexString string) string {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal("Error")
	}
	base64String := b64.StdEncoding.EncodeToString(decoded)
	return base64String
}
