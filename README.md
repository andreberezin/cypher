1.  **Explain what the tool does**
    The tool either encrypts or decrypts a user inputted message and outputs the message.
    The user can choose for the tool either to encrypt or decrypt and can choose also the type of encryption/decrytion done from 3 possible variants.

2.  **Explains the usage of the tool, with example**
    - ```bash 
      go build
    - ```bash 
      go run .
    
Example:
```bash
    Welcome to Awais', Pärtel's and Andre's Cypher tool!

    Select operation:

    1.Encrypt
    2.Decrypt
    1

    Select cypher:

    1.ROT13
    2.Reverse
    3.Custom
    3

    Enter shift number:
    13

    Enter the message:
    Hei

    This is the encrytped message using custom encryption:
    Urv
```

3.  **Explain the cyphers used**

    _Custom (Ceasar)_
    For each alphabetic character, it calculates its shifted position in the alphabet by adding the shift value.
    If the shift causes the position to go beyond the range of the alphabet (more than 25 positions), it wraps around by taking the remainder when divided by 26 (the number of letters in the alphabet). This ensures that the shifted position stays within the alphabet range.
    Finally, it converts the shifted position back to ASCII representation by adding the ASCII value of 'a' or 'A', depending on the case, to get the correct ASCII value of the letter.
