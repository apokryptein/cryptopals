package main

import (
	"fmt"

	"github.com/apokryptein/cryptopals/encoding"
)

func main() {
	// Phrase to XOR encode using  a given key
	inPhrase := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	// Key used to encode inPhrase
	key := "ICE"

	// Answer to check against for accuracy
	answerCheck := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	result := encoding.RepeatKeyXor(inPhrase, key)

	// Output initial phrase and result
	fmt.Println(inPhrase)
	fmt.Println(result)

	// Check validity of result against answerCheck
	if result == answerCheck {
		fmt.Println("Answer checks out.")
	} else {
		fmt.Println(answerCheck)
	}
}
