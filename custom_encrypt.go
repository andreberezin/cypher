package main

import "strings"

// Custom encryption using Simple Caesar Cipher
func encryptSimpleCaesar(s string, shift int) string {
	var result strings.Builder
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result.WriteByte((char-'a'+byte(shift))%26 + 'a')
		} else if char >= 'A' && char <= 'Z' {
			result.WriteByte((char-'A'+byte(shift))%26 + 'A')
		} else {
			result.WriteByte(byte(char))
		}
	}
	return result.String()
}

// Custom decryption for Simple Caesar Cipher
func decryptSimpleCaesar(s string, shift int) string {
	// To decrypt, we shift in the opposite direction, so we negate the shift
	return encryptSimpleCaesar(s, -shift)
}
