package main

import "fmt"

func main() {

	// 1. Greet the user
	fmt.Printf("\nWelcome to Awais', Pärtel's and Andre's Cypher tool!\n\n")

	// 2. Allow the user to select the operation (encrypt or decrypt)
	// 3. Allow the user to select the encryption type
	// 4. Allow the user to input the message to encrypt/decrypt
	toEncrypt, encoding, message, shift := getAndValidateInput()

	// 5. Output the result of the operation based on user input
	if toEncrypt {
		switch encoding {
		case "rot13", "1":
			fmt.Printf("\nEncrypted message using ROT13:\n%v\n\n", Rot13_encrypt(message))
		case "reverse", "2":
			fmt.Printf("\nEncrypted message using reverse:\n%v\n\n", encryptreverse(message))
		case "custom", "3":
			fmt.Printf("\nEncrypted message using custom encryption:\n%v\n\n",encryptcaesar(message, shift))
		} 
	} else if !toEncrypt { // to decrypt
		switch encoding {
		case "rot13", "1":
			fmt.Printf("\nDecrypted message using ROT13:\n%v\n\n", decrypt_rot13(message))
		case "reverse", "2":
			fmt.Printf("\nDecrypted message using reverse:\n%v\n\n", encryptreverse(message))
		case "custom", "3":
			fmt.Printf("\nDecrypted message using custom decryption:\n%v\n\n",decryptcaesar(message, shift)) // decrypt with custom
		}
	}
}