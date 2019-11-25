package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

// Hextobase64 converts hex string to base64 string
func Hextobase64(hexString string) string {
	decoded, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal("Error")
	}
	fmt.Println("\n------Hex bytes:------\n ", decoded)
	base64String := b64.StdEncoding.EncodeToString(decoded)
	return base64String
}
