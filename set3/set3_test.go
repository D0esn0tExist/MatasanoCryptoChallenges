package set3

import (
	"log"
	"math/rand"
	"testing"
)

// TestPadOracleAttack simulates an attacker attempting a pad oracle attack on CBC
func TestPadOracleAttack(t *testing.T) {

	messages := []string{
		"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
		"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
		"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
		"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
		"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
		"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
		"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
		"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
		"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
		"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
	}
	// select a message at random and send to server
	message := messages[rand.Intn(len(messages)-1)]

	cipher, iv := PadServerEncryption([]byte(message))
	log.Printf("Cipher: %v", cipher)

	// TODO: Test FindPadLength().
	padLength := FindPadLength(cipher, iv)

	decryptedMsg := string(AttackOracle(cipher, iv, padLength))
	log.Printf("Decrypted bytes: %v", decryptedMsg)
	if decryptedMsg != message {
		t.Errorf("AttackOracle() fail. Expected: %v. Found: %v", message, decryptedMsg)
	}
}
