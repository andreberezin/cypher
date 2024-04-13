package main

import "strings"

func Reverse_encrypt(s string) string {
	var result strings.Builder

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			result.WriteByte(byte('z' - (char - 'a')))
		} else if char >= 'A' && char <= 'Z' {
			result.WriteByte(byte('Z' - (char - 'A')))
		} else {
			result.WriteByte(byte(char))
		}
	}

	return result.String()
}
