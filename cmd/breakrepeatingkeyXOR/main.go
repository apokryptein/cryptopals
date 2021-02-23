package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"sort"

	"github.com/apokryptein/cryptopals/cryptanalysis"
)

// Following code idea taken from the following URL for sorting dictionaries based on values
// https://medium.com/@kdnotes/how-to-sort-golang-maps-by-value-and-key-eedc1199d944

// Struct type holding keysize and hamming distance pair
type hamMap struct {
	keySize         int
	hammingDistance int
}

type hamDistances []hamMap

func (p hamDistances) Len() int           { return len(p) }
func (p hamDistances) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p hamDistances) Less(i, j int) bool { return p[i].hammingDistance < p[j].hammingDistance }

func main() {

	// Hamming Distance test
	// Result should be 37
	/*
		string1 := "this is a test"
		string2 := "wokka wokka!!!"
		hamDist := cryptanalysis.HammingDistance([]byte(string1), []byte(string2))
		fmt.Println(hamDist)
	*/

	hamDists := hamDistances{}

	// Read in data from file
	inData, err := os.ReadFile("../../data/6.txt")

	// Check for read error
	if err != nil {
		fmt.Println("Error reading file.")
		os.Exit(0)
	}

	// Unencode data from base64 and store in byte array
	encText, err := base64.StdEncoding.DecodeString(string(inData))

	// Check for unencoding error
	if err != nil {
		fmt.Println("Error decoding data.")
		os.Exit(0)
	}

	// Go through all key sizes from 2 to 40, take two blocks the size of each key size
	// Take hamming distance of two blocks, and store in struct list
	for keySize := 2; keySize <= 40; keySize++ {
		hamDist := cryptanalysis.HammingDistance(encText[:keySize], encText[keySize:keySize*2])
		hamDists = append(hamDists, hamMap{
			keySize:         keySize,
			hammingDistance: hamDist,
		})
	}

	// Sort according to hamming distance
	sort.Sort(hamDists)
	fmt.Println(hamDists)

}
