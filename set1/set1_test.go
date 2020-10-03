package set1

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

// TestBreakRepeat is a test case for this function
func TestSingleXor(t *testing.T) {
	cicoded, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	c := SingleXor(cicoded)
	if c.msg != "Cooking MC's like a pound of bacon" {
		t.Errorf("SingleXor = %s. Want: Cooking MC's like a pound of bacon.", c.msg)
	}
}

func TestXor(t *testing.T) {
	input, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	xoree, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	xor := hex.EncodeToString(Xor(input, xoree))
	if xor != "746865206b696420646f6e277420706c6179" {
		t.Errorf("Xor(input,xoree) = %s. Want: 746865206b696420646f6e277420706c6179", xor)
	}
}

func TestHextobase64(t *testing.T) {
	ans := Hextobase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if ans != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		t.Errorf("Hextobase64() = %s. Want: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t", ans)
	}
}

func TestDetect(t *testing.T) {
	ans := Detect()
	if ans.msg != "Now that the party is jumping\n" {
		t.Errorf("Detect() = %s. Want: Now that the party is jumping", ans.msg)
	}
}

func TestRepeatingXor(t *testing.T) {
	ans := RepeatingXor("Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal", "ICE")
	if ans != "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20690a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f" {
		t.Errorf("RepeatingXor with key ICE = %s. Want: 0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20690a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f", ans)
	}
}

func TestBreakRepeat(t *testing.T) {
	// TODO: Write this test
}

// TODO: Write benchmarks for the slower functions. :)

func TestAesdecrypt(t *testing.T) {
	loaded := LoadFile("aes.txt")

	c := Ciph{}
	c.CipherKey = []byte("YELLOW SUBMARINE")
	c.CipherText = make([]byte, len(loaded))
	base64.RawStdEncoding.Decode(c.CipherText, loaded)

	decoded := c.Aesdecrypt()
	if decoded[:10] != "I'm back a" {
		t.Errorf("Aesdecrypt() = %s. Want: In ecstasy", decoded[:10])
	}
}

func TestDetectAes(t *testing.T) {
	found := DetectAes("detect_aes.txt")
	if len(found) != 1 {
		t.Errorf("Excepting 1 aes cipher detected.")
	}
	if found[0][:10] != "d880619740" {
		t.Errorf("DetectAes() = %s. Want: ds", found[0][:10])
	}
}

func TestPKCSPadding(t *testing.T) {
	byteText := []byte("YELLOW SUBMARINE")
	paddedString := PKCSPadding(byteText, 20)
	if string(paddedString) != "YELLOW SUBMARINE\x04\x04\x04\x04" {
		t.Errorf("PKCSPadding(). Want: %s. Expected: YELLOW SUBMARINE\x04\x04\x04\x04", string(paddedString))
	}
}

func TestPKCSUnpadding(t *testing.T) {
	// test success case
	paddedBytes := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	unpaddedText := string(PKCSUnpadding(paddedBytes))
	expected := "YELLOW SUBMARINE"
	if string(unpaddedText) != expected {
		t.Errorf("PKCSUnpadding(). Want: %s. Expected: %s", unpaddedText, expected)
	}
	// test unpadded string
	paddedBytes = []byte("ICE ICE BABY'\x01\x02\x03\x04")
	unpaddedText = string(PKCSUnpadding(paddedBytes))
	expected = "ICE ICE BABY'\x01\x02\x03\x04"
	if string(unpaddedText) != expected {
		t.Errorf("PKCSUnpadding(). Want: %s. Expected: %s", unpaddedText, expected)
	}
}

func TestPKCSValidation(t *testing.T) {
	// success test
	testText := []byte("ICE ICE BABY\x04\x04\x04\x04")
	isValid := PKCSValidation(testText)
	if !isValid {
		t.Errorf("PKCSValidation(). Got:%v. Expected: true", isValid)
	}
	// test for wrong pad value
	testText = []byte("ICE ICE BABY\x05\x05\x05\x05")
	isValid = PKCSValidation(testText)
	if isValid {
		t.Errorf("PKCSValidation(). Got:%v. Expected: false", !isValid)
	}
	// test for wrong sequence
	testText = []byte("ICE ICE BABY\x01\x02\x03\x04")
	isValid = PKCSValidation(testText)
	if isValid {
		t.Errorf("PKCSValidation(). Got:%v. Expected: false", !isValid)
	}
}
