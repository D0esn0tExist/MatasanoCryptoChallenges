package main

import (
	"fmt"
)

func main() {
	//First challenge, Set 1.
	//fmt.Println("\n------Hex to Base64------")
	// hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	// fmt.Println("\nDecoded base64 string: ", Hextobase64(hexString))

	// Second challenge, Set 1.
	// fmt.Println("\n------XOR!------")
	// input, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	// fmt.Println("input: ", input)
	// fmt.Println("Xor output:", Xor(input))

	// SingleXor()

	//Third challenge, Set 1.
	fmt.Println("\n------Single-byte XOR cipher------")
	SingleXor()

}
