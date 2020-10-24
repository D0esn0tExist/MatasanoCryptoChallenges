package set1

import "log"

// PKCSPadding accounts for irregularly-sized messages is by padding, creating a plaintext that is an even multiple of the blocksize.
func PKCSPadding(input []byte, toLen int) []byte {
	diff := len(input) % toLen
	pad := 0
	if diff > 0 {
		pad = toLen - diff
	}
	paddedText := make([]byte, len(input)+pad)
	n := copy(paddedText, input)
	for i := 0; i < pad; i++ {
		paddedText[n+i] = byte(pad)
	}
	log.Printf("Pad length: %v", pad)
	return paddedText
}

// PKCSUnpadding removes padding on the provided byte array.
func PKCSUnpadding(input []byte) []byte {
	isValid := PKCSValidation(input)
	if !isValid {
		return input // if not padded, just return original byte array.
	}
	padByte := int(input[len(input)-1])
	return input[:len(input)-padByte]
}

// PKCSValidation function checks if a plaintext has PKCS padding.
func PKCSValidation(input []byte) (isPadded bool) {
	isPadded = false
	validationByte := int(input[len(input)-1])
	if validationByte > len(input) || validationByte == 0 {
		return
	}
	for i := 0; i < validationByte; i++ {
		if int(input[len(input)-1-i]) != validationByte {
			return
		}
	}
	isPadded = true
	return
}
