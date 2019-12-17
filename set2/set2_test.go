package set2

import "testing"

func TestPKCSPadding(t *testing.T) {
	paddedString := PKCSPadding("YELLOW SUBMARINE", 20)
	if paddedString != "YELLOW SUBMARINE\x04\x04\x04\x04" {
		t.Errorf("PKCSPadding(). Want: %s. Expected: YELLOW SUBMARINE\x04\x04\x04\x04", paddedString)
	}
}
