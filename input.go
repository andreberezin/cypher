package main

import (
	"fmt"
	"strings"
)

func getAndValidateInput() (toEncrypt bool, encoding string, message string) {
	// Allow the user to select the operation (encrypt or decrypt)
	toEncrypt_string := ""
	TOENCRYPT: fmt.Printf("Select operation: \n 1. Encrypt\n 2. Decrypt\n")
	fmt.Scan(&toEncrypt_string)

	// remove leading and trailing whitespace and make string to lowercase
	toEncrypt_string = strings.ToLower(strings.TrimSpace(toEncrypt_string)) 

	// check if input is valid, else ask for input again
	if toEncrypt_string != "encrypt" && toEncrypt_string != "decrypt" {
		fmt.Println("Invalid input, enter again.")
		goto TOENCRYPT
	}

	if toEncrypt_string == "Encrypt" {
		toEncrypt = true
	}

	// fmt.Printf("You chose: %v\n\n", toEncrypt) // for checking output

	// Allow the user to select the encryption type
	ENCODING: fmt.Printf("Select operation: \n 1. ROT13\n 2. Reverse\n 3. Custom\n")
	fmt.Scan(&encoding)

	// remove leading and trailing whitespace and make string to lowercase
	encoding = strings.ToLower(strings.TrimSpace(encoding)) 

	// check if input is valid, else ask for input again
	if encoding != "rot13" && encoding != "reverse" && encoding != "custom" {
		fmt.Println("Invalid input, enter again.")
		goto ENCODING
	}

	// fmt.Printf("You chose: %v\n\n", encoding) // for checking output

	// Allow the user to input the message to encrypt/decrypt
	fmt.Printf("Enter the message:\n")
	fmt.Scan(&message)

	// remove leading and trailing whitespace
	message = strings.TrimSpace(message) 

	// fmt.Printf("Your message: %v\n", message) // for checking output

	// Output the result of the operation

	return toEncrypt, encoding, message
}


// func validateInput(toEncrypt bool, encoding string, message string) (bool, bool, bool) {
// 	isValidtoEncrypt := toEncrypt == true || toEncrypt == false
// 	isValidEncoding:= (encoding == "rot13" || encoding == "reverse" || encoding =="custom")
// 	isValidMessage := strings.Contains(message, "") // mida peab message sisaldama?

// 	fmt.Println(isValidtoEncrypt, isValidEncoding, isValidMessage) // checking output

// 	return isValidtoEncrypt, isValidEncoding, isValidMessage
// }