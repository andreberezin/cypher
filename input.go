package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"strconv"
)

func getAndValidateInput() (toEncrypt bool, encoding string, message string, shift int) {
	// Allow the user to select the operation (encrypt or decrypt)
	toEncrypt_string := ""
	TOENCRYPT: fmt.Printf("Select operation: \n 1. Encrypt\n 2. Decrypt\n")
	fmt.Scan(&toEncrypt_string)

	// remove leading and trailing whitespace and make string to lowercase
	toEncrypt_string = strings.ToLower(strings.TrimSpace(toEncrypt_string)) 

	// check if input is valid, else ask for input again
	if toEncrypt_string != "encrypt" && toEncrypt_string != "1"  && toEncrypt_string != "decrypt" && toEncrypt_string != "2" {
		fmt.Println("Invalid input, enter again.")
		goto TOENCRYPT
	}

	if toEncrypt_string == "Encrypt" || toEncrypt_string == "1" {
		toEncrypt = true
	}

	// Allow the user to select the encryption type
	ENCODING: fmt.Printf("Select operation: \n 1. ROT13\n 2. Reverse\n 3. Custom\n")
	fmt.Scan(&encoding)

	// remove leading and trailing whitespace and make string to lowercase to check validity
	encoding = strings.ToLower(strings.TrimSpace(encoding)) 

	// check if input is valid, else ask for input again
	if encoding != "rot13" && encoding != "1" && encoding != "reverse" && encoding != "2" && encoding != "custom" && encoding != "3"{
		fmt.Println("Invalid input, enter again.")
		goto ENCODING
	}

	if encoding == "custom" || encoding == "3" {
		fmt.Println("Enter shift number:")
		fmt.Scan(&shift)
	}

	// Allow the user to input the message to encrypt/decrypt
	reader := bufio.NewReader(os.Stdin) // get user input with bufio to include whitespaces
	fmt.Printf("Enter the message:\n")
	shiftInput, _ := reader.ReadString('\n')
	message = strings.TrimSpace(shiftInput)  // remove leading and trailing whitespace

	// return the result of the operation
	return toEncrypt, encoding, message, shift
}