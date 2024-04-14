package main

import "fmt"

func main() {

	// 1. Greet the user
	fmt.Println("Welcome to Awais', Pärtel's and Andre's Cypher tool!")

	// 2. Allow the user to select the operation (encrypt or decrypt)
	// 3. Allow the user to select the encryption type
	// 4. Allow the user to input the message to encrypt/decrypt
	getAndValidateInput()
	toEncrypt, encoding, message := getAndValidateInput()

	// 5. Output the result of the operation based on user input
	getOutput(toEncrypt, encoding, message)

	//var output string

	// if toEncrypt == true {
	// 	switch encoding {
	// 	case "rot13":
	// 		// encrypt with rot13
	// 	case "reverse":
	// 		// encrypt with reverse
	// 	case "custom":
	// 		// encrypt with custom
	// 	} 
	// } else if toEncrypt == false { // meaning to decrypt
	// 	switch encoding {
	// 	case "rot13":
	// 		// decrypt with rot13
	// 	case "reverse":
	// 		// decrypt with reverse
	// 	case "custom":
	// 		// decrypt with custom
	// 	}

}
