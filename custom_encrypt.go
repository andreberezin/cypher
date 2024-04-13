package cypher

import "strings"

func encryptcaesar(s string, shift int) string {
    var result strings.Builder
    // To handle negative shifts properly
    shift = (shift%26 + 26) % 26
    for , char := range s {
        if char >= 'a' && char <= 'z' {
            // Shift lowercase letters
            result.WriteByte(byte('a' + (char-'a'+rune(shift))%26))
        } else if char >= 'A' && char <= 'Z' {
            // Shift uppercase letters
            result.WriteByte(byte('A' + (char-'A'+rune(shift))%26))
        } else {
            // Append non-alphabetic characters unchanged
            result.WriteByte(byte(char))
        }
    }
    return result.String()
}

// Custom decryption for Simple Caesar Cipher
func decrypt_caesar(s string, shift int) string {
    // To decrypt, we shift in the opposite direction, so we negate the shift
    return encrypt_caesar(s, -shift)
}