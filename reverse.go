package cypher

import "strings"

//This is the reverse encyrption/decription function

func encryptreverse(message string) string {
	var result strings.Builder
	for _, char := range message {
		if char >= 'a' && char <= 'z' {
			x := char - 'a'
			result.WriteRune('z' - x)
		} else if char >= 'A' && char <= 'Z' {
			y := char - 'A'
			result.WriteRune('Z' - y)
		} else {
			result.WriteRune(char)
		}

	}
	return result.String()

}
