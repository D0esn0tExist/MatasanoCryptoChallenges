package set2

import "testing"

func TestPKCSPadding(t *testing.T) {
	byteText := []byte("YELLOW SUBMARINE")
	paddedString := PKCSPadding(byteText, 20)
	if string(paddedString) != "YELLOW SUBMARINE\x04\x04\x04\x04" {
		t.Errorf("PKCSPadding(). Want: %s. Expected: YELLOW SUBMARINE\x04\x04\x04\x04", string(paddedString))
	}
}

func TestCbc(t *testing.T) {
	// byteText := []byte("YELLOW SUBMARINE")
	// iv := make([]byte, 16)
	// c := CbcProp{iv, byteText}
}
