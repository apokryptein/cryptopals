package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func main() {

	// Hamming Distance test
	// Result should be 37
	/*
		string1 := "this is a test"
		string2 := "wokka wokka!!!"
		hamDist := cryptanalysis.HammingDistance([]byte(string1), []byte(string2))
		fmt.Println(hamDist)
	*/

	inData, err := os.ReadFile("../../data/6.txt")

	if err != nil {
		fmt.Println("Error reading file.")
		os.Exit(0)
	}

	encText, err := base64.StdEncoding.DecodeString(string(inData))

	if err != nil {
		fmt.Println("Error decoding data.")
		os.Exit(0)
	}

	fmt.Println(encText)

}
