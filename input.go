package cypher

import "fmt"

func getInput() (toEncrypt bool, encoding string, message string) {
	// Allow the user to select the operation (encrypt or decrypt)
	fmt.Printf("Select operation: \n 1. Encrypt\n 2. Decrypt\n")
	fmt.Scan(&toEncrypt)
	fmt.Printf("You chose: %v\n\n", toEncrypt)

	// Allow the user to select the encryption type
	fmt.Printf("Select operation: \n 1. ROT13\n 2. Reverse\n 3. Custom\n")
	fmt.Scan(&encoding)
	fmt.Printf("You chose: %v}\n\n", encoding)

	// Allow the user to input the message to encrypt/decrypt
	fmt.Printf("Enter the message:\n")
	fmt.Scan(&message)
	fmt.Printf("Your message: %v\n", message)

	// Output the result of the operation

	return toEncrypt, encoding, message
}