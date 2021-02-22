package main

import (
	"encoding/hex"
	"fmt"

	"github.com/apokryptein/cryptopals/cryptanalysis"
	"github.com/apokryptein/cryptopals/encoding"
)

func main() {
	// Original hex-encoded string
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// Decode hex string to byte array
	byteString, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	// float64 variable to track phrase with highest score
	highScore := 0.0

	// rune variable to track key value associated with the highest score
	var decryptKey int

	// Iterate through alphaList and generate a XOR result for reach run
	// Take result and score
	for val := 0; val < 256; val++ {
		xorVal := encoding.SingleByteXor(byteString, val)
		currentScore := cryptanalysis.ScoreResult(xorVal)

		// checks for highest score
		if currentScore > highScore {
			highScore = currentScore
			decryptKey = val
		}
	}

	// Generate and display result
	fmt.Printf("Key: %c\n", decryptKey)
	fmt.Println(string(encoding.SingleByteXor(byteString, decryptKey)))
}
