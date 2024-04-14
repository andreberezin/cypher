package main

import "fmt"

func main() {

	// 1. Greet the user
	fmt.Println("Welcome to Awais', Pärtel's and Andre's Cypher tool!")

	// 2. Allow the user to select the operation (encrypt or decrypt)
	// 3. Allow the user to select the encryption type
	// 4. Allow the user to input the message to encrypt/decrypt
	toEncrypt, encoding, message, shift := getAndValidateInput()

	// 5. Output the result of the operation based on user input
	if toEncrypt == true {
		switch encoding {
		case "rot13", "1":
			fmt.Printf("This is the encrypted message:\n%v\n", Rot13_encrypt(message))
		case "reverse", "2":
			fmt.Printf("This is the encrypted message:\n%v\n", encryptreverse(message))
		case "custom", "3":
			fmt.Printf("This is the encrypted message:\n%v\n",encryptcaesar(message, shift))
		} 
	} else if toEncrypt == false { // to decrypt
		switch encoding {
		case "rot13", "1":
			fmt.Printf("This is the decrytped message:\n%v\n", decrypt_rot13(message))
		case "reverse", "2":
			fmt.Printf("This is the decrytped message:\n%v\n", encryptreverse(message))
		case "custom", "3":
			fmt.Printf("This is the decrytped message:\n%v\n",decryptcaesar(message, shift)) // decrypt with custom
		}
	}
}