package strings

import (
	"crypto/des"
	"fmt"
)

func FormatByteStringDES(input string) []byte {
	blockSize := des.BlockSize
	fmt.Println("block size:", blockSize)
	fmt.Println("input length:", len(input))
	fmt.Println("input length mod block size:", len(input) % blockSize)
	padding := blockSize - len(input) % blockSize
	fmt.Println("padding:", padding)
	padText := make([]byte, padding)
	for i := range padText {
		padText[i] = byte(0)
	}
	return append([]byte(input), padText...)
}