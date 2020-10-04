package set2

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

func TestPad(t *testing.T) {
	testInput := []byte("World")
	testPrefix := "Hello "
	testSuffix := " FROM ME"
	prefExpect := "Hello World"
	suffixExpect := "Hello World FROM ME"
	// test prepend prefix
	result := PrependPrefix(testPrefix, testInput)
	if string(result) != prefExpect {
		t.Errorf("PrependPrefix() fail. Got: %v Expected: %v", string(result), prefExpect)
	}
	// test append suffix
	result = AppendSuffix(testSuffix, result)
	if string(result) != suffixExpect {
		t.Errorf("AppendSuffix() fail. Got %v Expected %v", string(result), suffixExpect)
	}
	// test sanitization
	// sanitizeRule
	removeSpecial := func(text string) string {
		r := strings.NewReplacer("=", "", ";", "")
		r.Replace(text)
		return text
	}
	testPrefix = "Hel=lo; "
	padded := PadInput(removeSpecial, testPrefix, testSuffix, "World")
	if padded != suffixExpect {
		t.Errorf("PadInput() fail. Got: %v Expected %v", padded, suffixExpect)
	}
}

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

func TestCbcBitFlip(t *testing.T) {
	// random AES key
	encKey := KeyGen(16)
	// pad input and sanitize. Allign such that all the input is in one set
	input := "adminadminatrue"
	prefix := "comment1=cooking%20MCs;userdata="
	suffix := ";comment2=%20like%20a%20pound%20of%20bacon"

	// sanitizeRule
	removeSpecial := func(text string) string {
		text = strings.Replace(strings.Replace(text, "=", "", -1), ";", "", -1)
		return text
	}

	// pad and sanitize input
	paddedYsanitizedInput := PadInput(removeSpecial, prefix, suffix, input)
	// Encrypt
	c := CbcProp{
		Message: []byte(paddedYsanitizedInput),
		Key:     []byte(encKey),
		IV:      make([]byte, 16),
	}
	ciphertext := c.CbcEncrypt()
	fmt.Println(ciphertext)

	// identify characters to modify
	plainSize := len(paddedYsanitizedInput)
	for i := 0; i < plainSize/aes.BlockSize; i++ {
		fmt.Println(paddedYsanitizedInput[i*aes.BlockSize : (i+1)*aes.BlockSize])
	}
	/*
			comment1cooking%
			20MCsuserdataadm
			inadminatruecomm
			ent2%20like%20a%
			20pound%20of%20b

		Attempt a byte flip. Modify:
		* n : position 33 -> ;
		* a : position 39 -> =
		* c : position 44 -> ;
	*/
	toInsert1 := []byte(";")[0]
	toInsert2 := []byte("=")[0]
	positionsToModify := []int{33, 39, 44}
	bytesToInsert := []byte{toInsert1, toInsert2, toInsert1}
	BitFlip(ciphertext, []byte(paddedYsanitizedInput), positionsToModify, bytesToInsert)
	fmt.Println(ciphertext)
	var decrypted []byte

	// check existence of ";admin=true;" in plaintext
	inspectPlain := func(ciphertext []byte) bool {
		stringToCheck := ";admin=true;"
		exists := false
		c.CipherText = ciphertext
		decrypted = c.CbcDecrypt()
		if strings.Contains(string(decrypted), stringToCheck) {
			exists = true
		}
		return exists
	}
	if !inspectPlain(ciphertext) {
		t.Errorf("CbcBitFlip fail. Expected plaintext to contain ;admin=true; Got: %v.", string(decrypted))
	}
}
