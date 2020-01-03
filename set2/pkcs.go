package set2

// PKCSPadding accounts for irregularly-sized messages is by padding, creating a plaintext that is an even multiple of the blocksiz
func PKCSPadding(input []byte, toLen int) []byte {
	pad := len(input) % toLen
	paddedText := make([]byte, len(input)+pad)
	if pad > 0 {
		n := copy(paddedText, input)
		for i := 0; i < pad; i++ {
			paddedText[n+i] = byte(pad)
		}
	}
	return paddedText
}
