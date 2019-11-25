package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	//First challenge, Set 1.
	fmt.Println("\n------Hex to Base64------")
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println("\nDecoded base64 string: ", Hextobase64(hexString))

	// Second challenge, Set 1.
	fmt.Println("\n------XOR!------")
	input, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	xoree, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	fmt.Println("input: ", input)
	fmt.Println("Xor output:", Xor(input, xoree))

	// //Third challenge, Set 1.
	fmt.Println("\n------Single-byte XOR cipher------")
	cicoded, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	SingleXor(cicoded)

	// Fourth challenge, Set 1.
	fmt.Println("\n------Detect single-character XOR------")
	Detect()

	// Fifth challenge, Set 1.
	fmt.Println("\n------Implement repeating-key XOR------")
	RepeatingXor("Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal", "ICE")

	// Sixth challenge, Set 1.
	fmt.Println("\n------Break repeating-key XOR------")
	BreakRepeat("base64data.txt")
}
