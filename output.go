package main



// func getOutput(toEncrypt bool, encoding, message string, shift int) (output string) {
	
// 	if toEncrypt == true {
// 		switch encoding {
// 		case "rot13":
// 			output = Rot13_encrypt(message)// encrypt with rot13
// 		case "reverse":
// 			output = encryptreverse(message)// encrypt with reverse
// 		case "custom":
// 			output = encryptcaesar(message, shift)// encrypt with custom
// 		} 
// 	} else if toEncrypt == false { // meaning to decrypt
// 		switch encoding {
// 		case "rot13":
// 			output = decrypt_rot13(message)// decrypt with rot13
// 		case "reverse":
// 			output = encryptreverse(message)// decrypt with reverse
// 		case "custom":
// 			output = decryptcaesar(message, shift)// decrypt with custom
// 		}

// 		// When encrypting or decrypting, ensure that any non-alphabet characters in the message are left unchanged.
// 	}
// 	return output
// }