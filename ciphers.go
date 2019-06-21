package bsencryption

// !! IMPORTANT !! Do not change this interface as the change will impact available implementations

// BSCipher - Cipher interface, implement all the methods to attach new cypher to bykovstorage
type BSCipher interface {
	CleanAndInit() // Full clean and initialization of the entity

	// SetPassword or SetKey shoudld be called before using ecnrytption
	SetPassword(string) error    // Password should be set before starting encryption/decryption
	SetPasswordKey([]byte) error // Setting key from bytes array
	GetPasswordKey() []byte      // Get password current key
	IsKeyGenerated() bool        // Check if key is available and encryption/decryption is possible

	Encrypt(string) (string, error) // Encrypt string, return encrypted base64 string
	Decrypt(string) (string, error) // Decrypt string, encrypted base64 string should be provided as input

	EncryptBLOB(string) ([]byte, error) // Encrypt string, return encrypted bytes
	DecryptBLOB([]byte) (string, error) // Decrypt encrypted bytes, return string

	EncryptBIN([]byte) ([]byte, error) // Encrypt bytes array, return encrypted bytes
	DecryptBIN([]byte) ([]byte, error) // Decrypt encrypted bytes, return bytes array

	GetGryptID() string    // Unique ID of cipher implementation, use any alphanumeric symbols, 8 chars
	GetCipherName() string // Human readable name of the implemented cypher
}

// Ciphers - Add newly implemented cyphers here with `Cyphers append(Cyphers, new(cypherImplementationHere))`
var Ciphers = []BSCipher{
	new(cipherAES256), // AES256 internal implementation
	new(cypherNONE),   // No encryption
}