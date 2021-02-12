// Converts hex to base64

package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Printf("Hex: %s\n", hexString)
	fmt.Printf("Type: %T\n", hexString)

	decHex, err := hex.DecodeString(hexString)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Decoded Hex: %s\n", decHex)
	fmt.Printf("Type: %T\n", decHex)

	base64String := base64.StdEncoding.EncodeToString(decHex)
	fmt.Printf("Base64 Encoded: %s\n", base64String)
	fmt.Printf("Type: %T\n", base64String)
}
