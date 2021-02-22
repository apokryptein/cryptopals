package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/apokryptein/cryptopals/cryptanalysis"
	"github.com/apokryptein/cryptopals/encoding"
)

func main() {
	// open file containing strings
	file, err := os.Open("data/xor-strings-full.txt")

	// check for errors
	if err != nil {
		log.Fatal(err)
	}

	// instantiate a scanner object to iterate over lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// string array to store each line of file
	var inLines []string

	// checking for errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// scan through each line of file and append to string array
	for scanner.Scan() {
		inLines = append(inLines, scanner.Text())
	}

	// close file
	file.Close()

	// float64 variable to track phrase with highest score
	highScore := 0.0

	// rune variable to track key value associated with the highest score
	var decryptKey rune

	// []byte array variable to track
	var xorPhrase []byte

	// iterate over string array, trim whitespace, decode hex and store in byte array
	for _, line := range inLines {
		stringLine := strings.TrimSpace(line)
		byteString, err := hex.DecodeString(stringLine)

		if err != nil {
			fmt.Println(err)
		}

		// xor line against each of first 256 unicode characters, score, if score is highest, store relevant values in variables for later retrieval
		for i := 0; i < 256; i++ {
			xorVal := encoding.SingleByteXor(byteString, i)

			if cryptanalysis.IsASCII(xorVal) == false {
				continue
			}

			currentScore := cryptanalysis.ScoreResult(xorVal)

			// checks for highest score
			if currentScore > highScore {
				highScore = currentScore
				decryptKey = rune(i)
				xorPhrase = xorVal
			}
		}
	}

	// Generate and display result
	fmt.Printf("Key: %c\n", decryptKey)
	fmt.Println(string(xorPhrase))
}
