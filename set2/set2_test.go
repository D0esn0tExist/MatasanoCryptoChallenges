package set2

import (
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
