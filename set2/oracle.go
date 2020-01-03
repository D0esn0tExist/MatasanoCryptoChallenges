package set2

/*
Write a function to generate a random AES key; that's just 16 random bytes.
Write a function that encrypts data under an unknown key --- that is, a function that generates a random key and encrypts under it.
The function should look like:
encryption_oracle(your-input)
=> [MEANINGLESS JIBBER JABBER]
Under the hood, have the function append 5-10 bytes (count chosen randomly) before the plaintext and 5-10 bytes after the plaintext.
Now, have the function choose to encrypt under ECB 1/2 the time, and under CBC the other half (just use random IVs each time for CBC). Use rand(2) to decide which to use.
*/

// EncryptionOracle is a blackbox- a CBC/AES encryption blackbox
func EncryptionOracle(input []byte) {

}
