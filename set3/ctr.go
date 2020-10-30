package set3

import (
	"github.com/matasano/MatasanoCryptoChallenges/set1"
)

var (
	nonce     []byte
	blockSize int
)

func init() {
	blockSize = 16
	// TODO: Randomly generate nonce. A non-random nonce is used for now
	nonce = make([]byte, blockSize)
}

// CtrMode perform AES encryption in CTR mode. A counter is added to the nonce on each block.
func CtrMode(key, inputText []byte, nonceRule func(int, []byte) []byte) []byte {
	var messageBlocks int = len(inputText) / blockSize
	var lastBlockSize int = len(inputText) - (messageBlocks * blockSize)
	var response []byte // holds output

	feedCtr := func(msgBlock, msgBlockSize int, message []byte) []byte {
		inputCtr := nonceRule(msgBlock, nonce) // running nonce modified with block number
		aesPayload := set1.Ciph{
			Message:   inputCtr,
			CipherKey: key,
		}
		// encrypt nonce
		keyStream := aesPayload.Aesencrypt()
		// xor with messageBlock
		keyStream = keyStream[:msgBlockSize]
		set1.Xor(keyStream, message)
		return keyStream
	}

	for i := 1; i <= messageBlocks; i++ {
		endRange := i * 16
		// message block
		messageBlock := inputText[endRange-16 : endRange]

		keyStream := feedCtr(i-1, blockSize, messageBlock)
		response = append(response, keyStream...)
	}
	if lastBlockSize > 0 {
		keyStream := feedCtr(messageBlocks, lastBlockSize, inputText[messageBlocks*blockSize:])
		response = append(response, keyStream...)
	}
	return response
}
