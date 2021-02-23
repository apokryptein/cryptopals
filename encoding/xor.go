package encoding

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// takes byte string and rune containing xor key and returns xor'd output
func SingleByteXor(inBytes []byte, let int) []byte {
	var outBytes []byte
	for i := 0; i < len(inBytes); i++ {
		outBytes = append(outBytes, inBytes[i]^byte(let))
	}
	return outBytes
}

func RepeatKeyXor(inPhrase string, key string) string {
	// Determine length of inPhrase
	phraseLength := len(inPhrase)

	// String variable to hold repeated key
	var repeatedKey string

	// Byte array to store XOR values
	xorVal := make([]byte, phraseLength)

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

	return hex.EncodeToString(xorVal)
}
