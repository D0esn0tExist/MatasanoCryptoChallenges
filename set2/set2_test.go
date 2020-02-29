package set2

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

func TestCbc(t *testing.T) {
	cbcCiph := set1.LoadFile("cbcPlain.txt")
	cbcBytes := make([]byte, len(cbcCiph))
	base64.RawStdEncoding.Decode(cbcBytes, cbcCiph)
	byteText := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	c := CbcProp{iv, byteText}
	out, _ := c.Cbc(cbcBytes)

	outText := base64.StdEncoding.EncodeToString(out)
	fmt.Println(outText)
}

func TestEnryptionOracle(t *testing.T) {
	EncryptionOracle("Hello World!")
}

func TestByteAtime(t *testing.T) {
	b := bytes.Buffer{}
	cipher, _ := base64.RawStdEncoding.DecodeString(AES128ECBSuffixoracle(b))
	cipherLen := len(cipher)
	initLen := cipherLen
	fmt.Println("Initlen: ", initLen)

	for cipherLen == initLen {
		b.WriteString("A")
		cipher, _ := base64.RawStdEncoding.DecodeString(AES128ECBSuffixoracle(b))
		cipherLen = len(cipher)
		fmt.Printf("With %s, len: %d\n", string(b.Bytes()), cipherLen)
	}
	blocksize := cipherLen - initLen
	inputSize := b.Len()
	fmt.Printf("Blocksize: %d \n InputSize: %d", blocksize, inputSize)

	unknown := FindUnknownString(blocksize)
	fmt.Println("Pad: ", unknown)
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
