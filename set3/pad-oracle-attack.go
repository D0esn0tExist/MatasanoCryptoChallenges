package set3

import (
	"bytes"
	"log"
	"os"
)

// custom logger
var logger = log.New(os.Stdout, "", log.Lshortfile)

// FindPadLength function determines the pad length of the ciphertext
// replacerByte is used to replace each byte until a invalid padding is detected.
func FindPadLength(ciphertext, iv []byte) int {
	var (
		replacerByte byte = 0
		startOfPad   int  = len(ciphertext)
	)
	// prepend iv with ciphertext for ease of traversal
	replacerCiphertext := bytes.Join([][]byte{iv, ciphertext}, []byte(""))

	isValid := PadServerValidatePad(
		replacerCiphertext[16:],
		replacerCiphertext[:16],
	)
	if !isValid { // string not padded
		return 0
	}

	for i := 0; i < len(ciphertext); i++ {
		replacerCiphertext[i] = replacerByte
		isValid = PadServerValidatePad(
			replacerCiphertext[16:],
			replacerCiphertext[:16],
		)
		if !isValid {
			startOfPad = i + 16 // plus 16 because a change in one set of 16 affects the exact byte position in the next set
			break
		}
	}
	padLength := len(replacerCiphertext) - startOfPad
	logger.Printf("The length of pad: %v", padLength)
	return padLength
}

// AttackOracle function performs a brute force on the ciphertext to attempt to decrypt
// to plaintext using pad oracle attack.
func AttackOracle(ciphertext, iv []byte, padLength int) []byte {
	var (
		numberOfBlocks = len(ciphertext) / 16
		unknownPlain   = make([]byte, len(ciphertext)-padLength)
	)

	// loop through the blocks
	for blockNumber := numberOfBlocks; blockNumber > 0; blockNumber-- {
		ciphertext = ciphertext[:blockNumber*16]
		// prepend iv with ciphertext for ease of traversal
		ivCipherBlock := bytes.Join([][]byte{iv, ciphertext}, []byte(""))

		bytePosition := 0
		// account for possible on the last block
		if padLength > 0 {
			bytePosition = padLength
			padLength = 0
		}

		for bytePosition < 16 {
			padByte := bytePosition + 1

			// replace bytes with next expected pad
			for i := 1; i <= bytePosition; i++ {
				modifyPad := ivCipherBlock[len(ivCipherBlock)-i-16]
				ivCipherBlock[len(ivCipherBlock)-i-16] = byte(padByte) ^ byte(bytePosition) ^ modifyPad
			}

			targetByteIdx := len(ivCipherBlock) - padByte - 16 // the byte to modify
			targetByte := ivCipherBlock[targetByteIdx]

			// loop until you find byte that would produce valid padding
			for correctByte := 0; correctByte <= 255; correctByte++ {
				ivCipherBlock[targetByteIdx] = byte(correctByte)
				isValid := PadServerValidatePad(ivCipherBlock[16:], ivCipherBlock[:16])
				if isValid {
					logger.Printf("viable byte found: %v", correctByte)
					// calculate xor to recover plain byte.
					// targetByte, modifyByte, padByte.. needing plainByte
					plainByte := byte(correctByte) ^ byte(padByte) ^ byte(targetByte)
					unknownPlain[(blockNumber*16)-1-bytePosition] = plainByte
					break
				}
			}
			bytePosition++
		}
	}
	logger.Printf("Decrypted bytes: %v", unknownPlain)
	return unknownPlain
}
