package set3

import (
	"log"
	"os"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
	"github.com/matasano/MatasanoCryptoChallenges/set2"
)

var (
	encKey string
)

func init() {
	encKey = set2.KeyGen(16) // generates random encryption key
	// custom logger
	var logger = log.New(os.Stdout, "", log.Lshortfile)
	logger.Printf("key: %v", encKey)
}

// // PadServerEncryptionResponse is a response from the padServer; a pair: ciphertext and iv
// type PadServerEncryptionResponse struct {
// 	ciphertext []byte
// 	iv         []byte
// }

// PadServerEncryption function encrypts message bytes via CBC with a random key and spits out padServerEncryptionResponse.
func PadServerEncryption(message []byte) ([]byte, []byte) {
	logger.Printf("--------PadServerEncryption--------\nmessage: %v", string(message))
	logger.Printf("message length: %v", len(message))
	// iv
	iv := make([]byte, 16)
	logger.Printf("iv used: %v", iv)
	properties := set2.CbcProp{
		IV:      iv,
		Message: message,
		Key:     []byte(encKey),
	}
	cipher := properties.CbcEncrypt()
	logger.Printf("ciphertext: %v", cipher)
	return cipher, iv
}

// PadServerValidatePad function consume the ciphertext produced by the first function,
// decrypt it, check its padding, and return true or false depending on whether the padding is valid
func PadServerValidatePad(ciphertext []byte, iv []byte) bool {
	properties := set2.CbcProp{
		IV:         iv,
		CipherText: ciphertext,
		Key:        []byte(encKey),
	}
	message := properties.CbcDecrypt()
	logger.Printf("The message: %v", message)
	isValid := set1.PKCSValidation(message)
	logger.Printf("isPadValid? : %v", isValid)
	return isValid
}
