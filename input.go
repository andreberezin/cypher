package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
			fmt.Printf("\nInvalid input, enter again.\n\n")
			goto TOENCRYPT
	}

	// Allow the user to select the encryption type
	ENCODING: fmt.Printf("\nSelect cypher: \n 1. ROT13\n 2. Reverse\n 3. Custom\n")
	fmt.Scan(&encoding)
	// remove leading and trailing whitespace and make string to lowercase to check validity
	encoding = strings.ToLower(strings.TrimSpace(encoding)) 
	

	var shiftInput string
	// check if input is valid, else ask for input again
	switch encoding {
	case "rot13", "1", "reverse", "2":
		break
	case "custom", "3": // if user chooses custom then ask for shift number and validate the input
		for {
            fmt.Printf("\nEnter shift number: ")
            reader := bufio.NewReader(os.Stdin)
            shiftInput, _ = reader.ReadString('\n')
            shiftInput = strings.TrimSpace(shiftInput)
            if !isNumeric(shiftInput) {
                fmt.Printf("\nInvalid input, enter again.\n")
				continue
            }
            shift, _ = strconv.Atoi(shiftInput) // Convert shiftInput to iteger
            break
        }

	default:
		fmt.Printf("\nInvalid input, enter again.\n")
		goto ENCODING
    }


	// Allow the user to input the message to encrypt/decrypt
	reader := bufio.NewReader(os.Stdin) // get user input with bufio to include whitespaces
	fmt.Printf("\nEnter the message:\n")
	message, _ = reader.ReadString('\n')
	message = strings.TrimSpace(message)  // remove leading and trailing whitespace

	// return the result of the operation
	return toEncrypt, encoding, message, shift
}

func isNumeric(s string) bool {
    _, err := strconv.Atoi(s)
    return err == nil
}
