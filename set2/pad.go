package set2

// PrependPrefix function prepends the given prefix string to the input bytes given and returns a bytearray.
func PrependPrefix(prefix string, input []byte) []byte {
	if len(prefix) == 0 {
		return input
	}
	inputBytes := []byte(prefix)
	inputBytes = append(inputBytes, input...)
	return inputBytes
}

// AppendSuffix function appends the given suffix string to the input bytes given and returns a bytearray.
func AppendSuffix(suffix string, input []byte) []byte {
	if len(suffix) == 0 {
		return input
	}
	input = append(input, []byte(suffix)...)
	return input
}

// PadInput function takes an input string and prepends and appends prefix and suffix if provided then sanitises based on provided rule.
func PadInput(sanitizeRule func(string) string, prefix, suffix, input string) string {
	paddedInput := AppendSuffix(suffix, PrependPrefix(prefix, []byte(input)))
	sanitizedInput := sanitizeRule(string(paddedInput))
	return sanitizedInput
}
