package main

import "strings"

func Reverse_encrypt(s string) string {
	var result strings.Builder

	// Iterate over each character in the input string
	for _, char := range s {
		// Check if the character is a lowercase letter
		if char >= 'a' && char <= 'z' {
			// Reverse the letter's position in the alphabet and append to the result
			result.WriteByte(byte('z' - (char - 'a')))
		} else if char >= 'A' && char <= 'Z' { // Check if the character is an uppercase letter
			// Reverse the letter's position in the alphabet and append to the result
			result.WriteByte(byte('Z' - (char - 'A')))
		} else { // If the character is not a letter
			// Append the character unchanged to the result
			result.WriteByte(byte(char))
		}
	}

	// Convert the result builder to a string and return
	return result.String()
}
