package set2

// PKCSPadding accounts for irregularly-sized messages is by padding, creating a plaintext that is an even multiple of the blocksiz
func PKCSPadding(input []byte, toLen int) []byte {
	pad := toLen - len(input)%toLen
	paddedText := make([]byte, len(input)+pad)
	if pad > 0 {
		n := copy(paddedText, input)
		for i := 0; i < pad; i++ {
			paddedText[n+i] = byte(pad)
		}
	}
	return paddedText
}

// PKCSUnpadding removes padding on the provided byte array.
func PKCSUnpadding(input []byte) []byte {
	lenInput := len(input)
	padLen := int(input[lenInput-1])
	if padLen < 0 {
		return input
	}
	return input[:lenInput-padLen]
}
