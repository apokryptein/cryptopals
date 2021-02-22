package main

import (
	"fmt"

	"github.com/apokryptein/cryptopals/cryptanalysis"
)

func main() {
	string1 := "this is a test"
	string2 := "wokka wokka!!!"

	hamDist := cryptanalysis.HammingDistance([]byte(string1), []byte(string2))

	fmt.Println(hamDist)
}
