package main

// Rot13_encrypt function takes a string and returns its ROT13 encryption
func Rot13_encrypt(s string) string {
	var result string = ""   // Initialize an empty string to store the encrypted result
	for _, char := range s { // Iterate over each character in the input string
		if char >= 'A' && char <= 'Z' { // Check if the character is an uppercase letter
			// Apply ROT13 encryption to uppercase letters
			// ROT13 shifts each letter by 13 positions in the alphabet
			// Modulus operation ensures that the result wraps around if it goes beyond Z
			result += string((char-'A'+13)%26 + 'A')
		} else if char >= 'a' && char <= 'z' { // Check if the character is a lowercase letter
			// Apply ROT13 encryption to lowercase letters
			// Same process as above for lowercase letters
			result += string((char-'a'+13)%26 + 'a')
		} else { // If the character is not an alphabet character
			result += string(char) // Leave non-alphabet characters unchanged
		}
	}
	return result // Return the encrypted string
}

// decrypt_rot13 function takes a string encrypted with ROT13 and returns its decrypted form.
func decrypt_rot13(s string) string {
	return encrypt_rot13(s) // Since ROT13 encryption and decryption are the same operation, simply return the result of applying ROT13 encryption to the input string.
}
