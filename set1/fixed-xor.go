package set1

// Xor takes two equal-length buffers and produces their XOR combination.
func Xor(input, xoree []byte) {
	for i := range input {
		input[i] ^= xoree[i]
	}
}
