package cryptos

import (
	"crypto/sha256"
	"crypto/sha512"
)

type hashExecuter struct {}


func (h *hashExecuter) ExecuteSHA256(message string) ([]byte,error) {
	byteMessage := []byte(message)
	sha256 := sha256.Sum256(byteMessage)
	return sha256[:], nil
}

func (h *hashExecuter) ExecuteSHA512(message string) ([]byte,error) {
	byteMessage := []byte(message)
	sha512 := sha512.Sum512(byteMessage)
	return sha512[:], nil
}