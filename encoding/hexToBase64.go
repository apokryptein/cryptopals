package encoding

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(hexString string) string {
	decHex, err := hex.DecodeString(hexString)

	if err != nil {
		fmt.Println(err)
	}

	base64String := base64.StdEncoding.EncodeToString(decHex)

	return base64String

}
