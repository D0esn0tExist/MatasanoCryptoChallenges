package set2

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

func TestPKCSPadding(t *testing.T) {
	byteText := []byte("YELLOW SUBMARINE")
	paddedString := PKCSPadding(byteText, 20)
	if string(paddedString) != "YELLOW SUBMARINE\x04\x04\x04\x04" {
		t.Errorf("PKCSPadding(). Want: %s. Expected: YELLOW SUBMARINE\x04\x04\x04\x04", string(paddedString))
	}
}

func TestPKCSUnpadding(t *testing.T) {
	paddedBytes := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	unpaddedText := string(PKCSUnpadding(paddedBytes))
	if string(unpaddedText) != "YELLOW SUBMARINE" {
		t.Errorf("PKCSUnpadding(). Want: %s. Expected: YELLOW SUBMARINE", unpaddedText)
	}
}

func TestCbc(t *testing.T) {
	cbcCiph := set1.LoadFile("cbcPlain.txt")
	cbcBytes := make([]byte, len(cbcCiph))
	base64.RawStdEncoding.Decode(cbcBytes, cbcCiph)
	byteText := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	c := CbcProp{iv, byteText}
	out, _ := c.Cbc(cbcBytes)
	out = PKCSUnpadding(out)

	outText := base64.StdEncoding.EncodeToString(out)
	fmt.Println(outText)
}
