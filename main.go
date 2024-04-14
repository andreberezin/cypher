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
	//getOutput(toEncrypt, encoding, message, shift)

	//output := ""
	if toEncrypt == true {
		switch encoding {
		case "rot13":
			fmt.Println(Rot13_encrypt(message))
			//output = Rot13_encrypt(message)// encrypt with rot13
		case "reverse":
			fmt.Println(encryptreverse(message))
			//output = encryptreverse(message)// encrypt with reverse
		case "custom":
			fmt.Println(encryptcaesar(message, shift))
			//output = encryptcaesar(message, shift)// encrypt with custom
		} 
	} else if toEncrypt == false { // meaning to decrypt
		switch encoding {
		case "rot13":
			fmt.Println(decrypt_rot13(message))
			//output = decrypt_rot13(message)// decrypt with rot13
		case "reverse":
			fmt.Println(encryptreverse(message))
			//output = encryptreverse(message)// decrypt with reverse
		case "custom":
			fmt.Println(decryptcaesar(message, shift)) // decrypt with custom
		}
	// output := getOutput(toEncrypt, encoding, message, shift)

	//return output
}
}