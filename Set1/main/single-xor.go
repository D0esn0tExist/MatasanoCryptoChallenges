package main

import (
	"encoding/hex"
	"log"
	"math"
)

// CipherScore struct define decoded message score.
type CipherScore struct {
	key    string
	cipher string
	msg    string
	freq   map[string]int
	score  int
}

// SingleXor finds the key, decrypts the message, does a letter freq calculation on msg,
// returns highest scoring.
func SingleXor(cipher string) CipherScore {
	ciph := make([]CipherScore, 256)
	cicoded, err := hex.DecodeString(cipher)
	if err != nil {
		log.Fatal("Error")
	}
	for i := 0; i < math.MaxUint8; i++ {
		ciph[i].key, ciph[i].msg = decrypt(cicoded, byte(i))
		ciph[i].cipher = cipher
		ciph[i].freq = freqChar([]byte(ciph[i].msg))
		ciph[i].score = charFreqScore(ciph[i])
	}

	max := ciph[0]
	for _, v := range ciph {
		if v.score > max.score {
			max = v
		}
	}
	// fmt.Println("Key: "+max.key+". \n Decoded message: "+max.msg+". \n Freq: ", max.score)
	return max
}

// This function decrypts the cipher message. Returns key and message.
func decrypt(cipher []byte, key byte) (foundChar, msg string) {
	message := make([]byte, len(cipher))

	for i := range cipher {
		message[i] = cipher[i] ^ key
	}
	return string(key), string(message)
}

// This function counts frequency of each letter in decoded message.
func freqChar(decoded []byte) (freq map[string]int) {
	freq = make(map[string]int)
	for _, v := range decoded {
		freq[string(v)]++
	}
	return freq
}

func charFreqScore(c CipherScore) (score int) {
	score = c.freq["e"] + c.freq["t"] + c.freq["o"] + c.freq["a"] + c.freq["i"] + c.freq["n"]
	return score
}
