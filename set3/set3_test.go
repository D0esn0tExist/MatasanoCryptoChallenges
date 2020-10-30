package set3

import (
	"encoding/base64"
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

func TestCtrMode(t *testing.T) {
	testCipher := "L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ=="
	testCipherBytes, err := base64.StdEncoding.DecodeString(testCipher)
	if err != nil {
		panic("Error decoding string")
	}
	keyBytes := []byte("YELLOW SUBMARINE")

	// rule on how to update nonce on each block. For this test, update the 9th byte incrementally from 0.
	nonceRule := func(blockNumber int, nonce []byte) []byte {
		nonce[8] = byte(blockNumber)
		return nonce
	}

	expectedMessage := "Yo, VIP Let's kick it Ice, Ice, baby Ice, Ice, baby "
	decryptedBytes := CtrMode(keyBytes, testCipherBytes, nonceRule)
	if string(decryptedBytes) != expectedMessage {
		t.Errorf("Decrytion fail. Expected: %v. Got: %v", expectedMessage, string(decryptedBytes))
	}

	encryptedBytes := CtrMode(keyBytes, decryptedBytes, nonceRule)
	testCipherResult := base64.StdEncoding.EncodeToString(encryptedBytes)

	logger.Printf("Encrypted message: %v", testCipherResult)
	if testCipherResult != testCipher {
		t.Errorf("CtrMode() fail. Expected: %v. Found: %v", testCipher, testCipherResult)
	}
}
