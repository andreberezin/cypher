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
	switch toEncrypt_string {
		case "encrypt", "1":
			toEncrypt = true
		case "decrypt", "2":
			toEncrypt = false
		default:
			fmt.Printf("\nInvalid input, enter again.\n")
			goto TOENCRYPT
	}

	// Allow the user to select the encryption type
	ENCODING: fmt.Printf("\nSelect operation: \n 1. ROT13\n 2. Reverse\n 3. Custom\n")
	fmt.Scan(&encoding)
	// remove leading and trailing whitespace and make string to lowercase to check validity
	encoding = strings.ToLower(strings.TrimSpace(encoding)) 

	// check if input is valid, else ask for input again
	switch encoding {
	case "rot13", "1", "reverse", "2":
		break
	case "custom", "3": // if user chooses custom then ask for shift number
		fmt.Printf("\nEnter shift number:\n")
		fmt.Scan(&shift)
	default:
		fmt.Printf("\nInvalid input, enter again.\n")
		goto ENCODING
}

	// Allow the user to input the message to encrypt/decrypt
	reader := bufio.NewReader(os.Stdin) // get user input with bufio to include whitespaces
	fmt.Printf("\nEnter the message:\n")
	shiftInput, _ := reader.ReadString('\n')
	message = strings.TrimSpace(shiftInput)  // remove leading and trailing whitespace

	// return the result of the operation
	return toEncrypt, encoding, message, shift
}