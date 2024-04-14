package main



func getOutput(toEncrypt bool, encoding, message string) (output string) {
	
	if toEncrypt == true {
		switch encoding {
		case "rot13":
			output = ""// encrypt with rot13
		case "reverse":
			output = ""// encrypt with reverse
		case "custom":
			output = ""// encrypt with custom
		} 
	} else if toEncrypt == false { // meaning to decrypt
		switch encoding {
		case "rot13":
			output = ""// decrypt with rot13
		case "reverse":
			output = ""// decrypt with reverse
		case "custom":
			output = ""// decrypt with custom
		}

		// When encrypting or decrypting, ensure that any non-alphabet characters in the message are left unchanged.
	}
	return output
}