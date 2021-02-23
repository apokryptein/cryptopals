package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// Hex buffer values, declared as strings
	buffer1 := "1c0111001f010100061a024b53535009181c"
	buffer2 := "686974207468652062756c6c277320657965"

	// Decode buffer1 to bytes []uint8
	b1Bytes, err := hex.DecodeString(buffer1)
	if err != nil {
		fmt.Println("Unable to decode Buffer 1")
	}

	// Decode buffer2 to bytes []uint8
	b2Bytes, err := hex.DecodeString(buffer2)
	if err != nil {
		fmt.Println("Unable to decode Buffer 2")
	}

	xorBytes := make([]byte, len(b1Bytes))

	for i := 0; i < len(b1Bytes); i++ {
		xorBytes[i] = b1Bytes[i] ^ b2Bytes[i]
	}

	fmt.Println(hex.EncodeToString(xorBytes))

}
