package main

import (
	"encoding/hex"
	"fmt"
	"unicode"
)

func main() {
	// Rune array of upper and lowercase letters
	// Will be used to as a XOR value for supplied hexString
	alphaList := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

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
	var decryptKey rune

	// Iterate through alphaList and generate a XOR result for reach run
	// Take result and score
	// Add score to array
	for _, val := range alphaList {
		xorVal := xorBytes(byteString, val)
		currentScore := scoreResult(xorVal)

		// checks for highest score
		if currentScore > highScore {
			highScore = currentScore
			decryptKey = val
		}
	}

	// Generate and display result
	fmt.Printf("Key: %c\n", decryptKey)
	fmt.Println(string(xorBytes(byteString, decryptKey)))

}

// takes bye string and rune containing xor key and returns xor'd output
func xorBytes(inBytes []byte, xorRune rune) []byte {
	var outBytes []byte
	for i := 0; i < len(inBytes); i++ {
		outBytes = append(outBytes, inBytes[i]^byte(xorRune))
	}
	return outBytes
}

// takes xor'd byte array and tests and scores text against English letter frequency
// returns phraseScore
func scoreResult(inBytes []byte) float64 {
	// [rune]flat64 map
	// Statistical letter frequencies taken from https://cs.wellesley.edu/~fturbak/codman/letterfreq.html
	// Value for space was estimated to 0.20
	var letterFreq = make(map[rune]float64)
	letterFreq['E'] = 0.124167
	letterFreq['T'] = 0.0969225
	letterFreq['A'] = 0.0820011
	letterFreq['O'] = 0.0714095
	letterFreq['I'] = 0.0768052
	letterFreq['N'] = 0.0764055
	letterFreq['S'] = 0.0706768
	letterFreq['R'] = 0.0668132
	letterFreq['H'] = 0.0350386
	letterFreq['D'] = 0.0363709
	letterFreq['L'] = 0.0448308
	letterFreq['U'] = 0.028777
	letterFreq['C'] = 0.0344391
	letterFreq['M'] = 0.0281775
	letterFreq['F'] = 0.0235145
	letterFreq['Y'] = 0.0189182
	letterFreq['W'] = 0.0135225
	letterFreq['G'] = 0.0181188
	letterFreq['P'] = 0.0203171
	letterFreq['B'] = 0.0106581
	letterFreq['V'] = 0.0124567
	letterFreq['K'] = 0.00393019
	letterFreq['X'] = 0.00219824
	letterFreq['Q'] = 0.0009325
	letterFreq['J'] = 0.0019984
	letterFreq['Z'] = 0.000599
	letterFreq[' '] = 0.20

	// variable containing score
	var phraseScore float64 = 0

	// iterates through byte string, looks for rune value in frequency table.
	// If in table, it adds the associated value to the phraseScore variable
	for _, let := range inBytes {
		val, ok := letterFreq[unicode.ToUpper(rune(let))]

		if ok {
			phraseScore += val
		}

	}
	return phraseScore
}
