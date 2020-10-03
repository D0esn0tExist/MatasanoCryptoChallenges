package set2

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

func TestCbcEncrypt(t *testing.T) {
	encryptionKey := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	c := CbcProp{
		Key: encryptionKey,
		IV:  iv}

	// Test text encrytion
	messageBytes := []byte("Hello World!")
	c.Message = messageBytes
	encryptedBytes := c.CbcEncrypt()
	encryptedContent := base64.StdEncoding.EncodeToString(encryptedBytes)
	expected := "CznrvEzPURGVlXa5MgPj3g=="
	if encryptedContent != expected {
		t.Errorf("CbcEncrypt(). Got: %v Expected: %v", encryptedContent, expected)
	}
}

func TestCbcDecrypt(t *testing.T) {
	decryptionKey := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	c := CbcProp{
		Key: decryptionKey,
		IV:  iv}
	// Test decryption of file.
	encryptedContent := set1.LoadFile("testCBCDec.txt")
	encryptedBytes := make([]byte, len(encryptedContent))
	_, err := base64.RawStdEncoding.Decode(encryptedBytes, encryptedContent)
	if err != nil {
		t.Fatal(err)
	}
	c.CipherText = encryptedBytes
	decryptedBytes := c.CbcDecrypt()
	decryptedContent := string(decryptedBytes)
	if !strings.Contains(decryptedContent, "ringin' the bell") {
		t.Errorf("CbcDecrypt() on file failed. Expected decrypted content to contain: ringin' the bell")
	}

	// Test text decrytion
	encryptedString := "CznrvEzPURGVlXa5MgPj3g=="
	encryptedBytes, err = base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		t.Fatal(err)
	}
	c.CipherText = encryptedBytes
	decryptedBytes = c.CbcDecrypt()
	decryptedContent = string(decryptedBytes)
	expected := "Hello World!"
	if decryptedContent != expected {
		t.Errorf("CbcDecrypt(). Got: %v Expected: %v", decryptedContent, expected)
	}
}
func TestEnryptionOracle(t *testing.T) {
	cipher := EncryptionOracle("Hello World!")
	fmt.Println(cipher)
}

func TestByteAtime(t *testing.T) {
	b := bytes.Buffer{}
	blocksize, unknownStringSize := BreakSuffixOracleLength(b)
	/*
		TODO: The FindUnknownSuffixPad function for now, because it specifically tries to
		break the AES128ECBSuffixoracle.
		Modify that function for generic finding padded e.g. suffix etc.
	*/
	unknown := FindUnknownSuffixPad(nil, unknownStringSize, blocksize)
	fmt.Println("Pad: ", unknown)
}

func TestBreakPrefixOracle(t *testing.T) {
	// Break prefix; find prefix length
	prefixSize := BreakPrefixOracleLength(16)
	fmt.Println("Size of prefix: ", prefixSize)
	// Constant pad of input by the prefix length; find suffix length
	b := bytes.Buffer{}
	prefixBytes := []byte(strings.Repeat("A", prefixSize))
	b.Write(prefixBytes)
	blocksize, suffixSize := BreakSuffixOracleLength(b)
	// find target: suffix pad
	unknown := FindUnknownSuffixPad(prefixBytes, suffixSize, blocksize)
	fmt.Println(unknown)
	if !strings.Contains(unknown, "Rollin' in my 5.0") {
		t.Errorf("Wrong pad. Pad contains: ")
	}
}

func TestPriv(t *testing.T) {
	// Generate random key.
	randKey := KeyGen(16)
	c := set1.Ciph{}
	c.CipherKey = []byte(randKey)

	cookie := ProfileFor("foo@bar.com")
	if cookie != "email=foo@bar.com&uid=10&role=user" {
		t.Errorf("Excepted email=foo@bar.com&uid=10&role=user. Found %s", cookie)
	}

	// Encrypt cookie.
	cookieEncrypter := func(email string) []byte {
		cookie := ProfileFor(email)
		c.Message = []byte(cookie)
		c.CipherText = c.Aesencrypt()
		return c.CipherText
	}

	synthCookie := PrivEsc(cookieEncrypter)
	c.CipherText = synthCookie
	newCookie := c.Aesdecrypt()
	profile := Parser(newCookie)
	if profile["role"] != "admin" {
		t.Errorf("Excepted admin. Found %s", profile["role"])
	}

}
