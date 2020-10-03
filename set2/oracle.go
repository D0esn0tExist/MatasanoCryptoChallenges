package set2

import (
	"math/rand"
	"strings"
	"time"

	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

/*
Write a function to generate a random AES key; that's just 16 random bytes.
Write a function that encrypts data under an unknown key --- that is, a function that generates a random key and encrypts under it.
The function should look like:
encryption_oracle(your-input)
=> [MEANINGLESS JIBBER JABBER]
Under the hood, have the function append 5-10 bytes (count chosen randomly) before the plaintext and 5-10 bytes after the plaintext.
Now, have the function choose to encrypt under ECB 1/2 the time, and under CBC the other half (just use random IVs each time for CBC). Use rand(2) to decide which to use.
*/

func init() {
	rand.Seed(time.Now().UnixNano())
}

// EncryptionOracle is a blackbox- a CBC/AES encryption blackbox
func EncryptionOracle(input string) string {
	theInput := padString(input)
	encKey := KeyGen(16)
	// Do AES ECB
	if rand.Intn(2) == 1 {
		c := set1.Ciph{CipherText: nil, CipherKey: []byte(encKey), Message: []byte(theInput)}
		cipher := c.Aesencrypt()
		return string(cipher)
	}

	// Do AES CBC
	iv := make([]byte, 16)
	_, err := rand.Read(iv)
	if err != nil {
		panic(err)
	}
	cbc := CbcProp{
		IV:      iv,
		Key:     []byte(encKey),
		Message: set1.PKCSPadding([]byte(theInput), 16),
	}
	cipher := cbc.CbcEncrypt()
	return string(cipher)
}

// KeyGen function generates a 16-length alphanumeric key. Good time to learn about runes too :)
// TODO: Will be back to play with crypto/rand package
func KeyGen(length int) string {
	var alphaNumRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	key := make([]rune, length)
	for i := range key {
		key[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes))]
	}
	return string(key)
}

func padString(unpadded string) string {
	count := 0
	for count < 5 {
		count = rand.Intn(10)
	}
	pad := strings.Repeat("0", count)
	padded := pad + unpadded + pad
	return padded
}
