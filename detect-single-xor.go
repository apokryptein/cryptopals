package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("xor-strings-full.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var inLines []string

	for scanner.Scan() {
		inLines = append(inLines, scanner.Text())
	}

	file.Close()

	// float64 variable to track phrase with highest score
	highScore := 0.0

	// rune variable to track key value associated with the highest score
	var decryptKey rune

	// []byte array variable to track
	var xorPhrase []byte

	for _, line := range inLines {
		stringLine := strings.TrimSpace(line)
		byteString, err := hex.DecodeString(stringLine)

		if err != nil {
			panic(err)
		}

		for i := 0; i < 256; i++ {
			xorVal := xorBytes(byteString, rune(i))

			if isASCII(xorVal) == false {
				continue
			}

			currentScore := scoreResult(xorVal)

			// checks for highest score
			if currentScore > highScore {
				highScore = currentScore
				decryptKey = rune(i)
				xorPhrase = xorVal
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Generate and display result
	fmt.Printf("Key: %c\n", decryptKey)
	fmt.Println(string(xorPhrase))

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

// Function taken from https://stackoverflow.com/questions/53069040/checking-a-string-contains-only-ascii-characters
// Modified
func isASCII(s []byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}
