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
	cipher := EncryptionOracle("Hello World!")
	fmt.Println(cipher)
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
	// Initial ciphertext length, initlen - (inputsize-1) gives length of the unknown string.
	unknownStringSize := initLen - (inputSize - 1)
	/*
		The FindUnknownString function for now, because it specifically tries to
		break the AES128ECBSuffixoracle.
		Modify that function for generic finding padded e.g. suffix etc.
	*/
	fmt.Printf("Blocksize: %d \nInputSize: %d\nLength of unknown string: %d\n", blocksize, inputSize, unknownStringSize)

	unknown := FindUnknownString(unknownStringSize, blocksize)
	fmt.Println("Pad: ", unknown)
}

func TestAES128ECBPrefixoracle(t *testing.T) {
	for i := 0; i <= 10; i++ {
		input := []byte("")
		randomCipher := AES128ECBPrefixoracle(input)
		fmt.Println("Prefixed and suffixed: ", randomCipher)
		fmt.Println("Length of cipher: ", len(randomCipher))
	}
	// input := []byte("AA")

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
