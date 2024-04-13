package main

import "fmt"

func getInput() (toEncrypt bool, encoding string, message string) {
	// Allow the user to select the operation (encrypt or decrypt)
	fmt.Printf("Select operation: \n 1. Encrypt\n 2. Decrypt\n")
	fmt.Scan(&toEncrypt)


	return toEncrypt, encoding, message
}