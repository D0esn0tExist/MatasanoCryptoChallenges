package set2

import (
	"crypto/aes"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

// CbcProp struct has the encryption properties
type CbcProp struct {
	Message    []byte
	CipherText []byte
	IV         []byte
	Key        []byte
}

// CbcEncrypt function handles encryption of a plaintext in CBC mode
func (c *CbcProp) CbcEncrypt() []byte {
	ciph, err := aes.NewCipher(c.Key)
	if err != nil {
		panic(err)
	}
	paddedBytes := set1.PKCSPadding(c.Message, aes.BlockSize)

	cipher := make([]byte, len(paddedBytes))
	ciphBlock := c.IV

	for idx := 0; idx < len(paddedBytes); idx += aes.BlockSize {
		lim := idx + aes.BlockSize
		set1.Xor(paddedBytes[idx:lim], ciphBlock)
		ciph.Encrypt(cipher[idx:lim], paddedBytes[idx:lim])
		ciphBlock = cipher[idx:lim]
	}
	return cipher
}

// CbcDecrypt function handles decryption of a plaintext in CBC mode
func (c *CbcProp) CbcDecrypt() []byte {
	ciph, err := aes.NewCipher(c.Key)
	if err != nil {
		panic(err)
	}
	cipherContent := c.CipherText
	plainBytes := make([]byte, len(c.CipherText))
	ciphBlock := c.IV

	for idx := 0; idx < len(plainBytes); idx += aes.BlockSize {
		lim := idx + aes.BlockSize

		ciph.Decrypt(plainBytes[idx:lim], cipherContent[idx:lim])
		set1.Xor(plainBytes[idx:lim], ciphBlock)
		ciphBlock = cipherContent[idx:lim]
	}
	plainBytes = set1.PKCSUnpadding(plainBytes)
	return plainBytes
}

// BitFlip attempts to do a byte flip attack on CBC ciphertext.
// The function takes in the ciphertext to modify, the original
// plaintext and a list of positions to modify and another list
// of the corresponding bytes to insert at the positions.
// NOTE: Match by index is important for these two lists.
func BitFlip(ciphertext []byte, originalPlain []byte, positionsToModify []int, bytesToInsert []byte) {
	if len(positionsToModify) != len(bytesToInsert) {
		panic("lengths must match")
	}
	for i, position := range positionsToModify {
		modifPosition := 0
		// TODO: If the position to be changed is in the first set, modify IV. Add this functionality.
		if position > 16 {
			modifPosition = position - 16
		}
		ciphertext[modifPosition] = ciphertext[modifPosition] ^ originalPlain[position] ^ bytesToInsert[i]
	}
}
