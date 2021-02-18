package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Phrase to XOR encode using  a given key
	inPhrase := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	// Key used to encode inPhrase
	key := "ICE"

	// Determine length of inPhrase
	phraseLength := len(inPhrase)

	// String variable to hold repeated key
	var repeatedKey string

	// Byte array to store XOR values
	xorVal := make([]byte, phraseLength)

	// Answer to check against for accuracy
	answerCheck := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	// Calculate number of times entire key goes into phrase length
	// This is done in order to calculate how many iterations of the key needs to be repeated in order to match inPhrase length
	keyLength := phraseLength / len(key)

	// If the length of key does not divide evenly into the length of the phrase, find by how much, and append a slice of key onto repeated key
	// If it does divide evenly, repeat key to match length of inPhrase
	if (keyLength * len(key)) < phraseLength {
		diff := phraseLength - (keyLength * len(key))
		repeatedKey = strings.Repeat(key, keyLength) + key[:diff]
	} else {
		repeatedKey = strings.Repeat(key, keyLength)
	}

	// Ensure key and phrase lengths match.  If lengths match, XOR and store in xorVal
	if len(repeatedKey) == phraseLength {
		for i := 0; i < phraseLength; i++ {
			xorVal[i] = byte(inPhrase[i]) ^ byte(repeatedKey[i])
		}
	} else {
		fmt.Println("Lengths don't match. Exiting...")
		os.Exit(3)
	}

	// Encode xorVal to hex and store in result
	result := hex.EncodeToString(xorVal)

	// Output initial phrase and result
	fmt.Println(inPhrase)
	fmt.Println(result)

	// Check validity of result against answerCheck
	if result == answerCheck {
		fmt.Println("Answer checks out.")
	} else {
		fmt.Println(answerCheck)
	}
}
