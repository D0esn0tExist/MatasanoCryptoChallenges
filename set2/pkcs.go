package set2

// PKCSPadding accounts for irregularly-sized messages is by padding, creating a plaintext that is an even multiple of the blocksiz
func PKCSPadding(text string, toLen int) string {
	byteText := []byte(text)
	pad := toLen - len(byteText)
	paddedText := make([]byte, len(byteText)+pad)
	if pad > 0 {
		n := copy(paddedText, byteText)
		for i := 0; i < pad; i++ {
			paddedText[n+i] = byte(pad)
		}
	}
	return string(paddedText)
}
