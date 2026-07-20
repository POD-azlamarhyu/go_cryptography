package cryptos

import (
	"crypto/des"
	"cryptography_tutorial/config"
)

type DESExecuter struct {}

func (s *DESExecuter) Encrypt(plainText []byte) ([]byte, error) {
	desKey := config.GetDESKey()
	cipherBlock, err := des.NewCipher(desKey)
	if err != nil {
		return nil, err
	}
	
	cipherText := make([]byte, len(plainText))
	cipherBlock.Encrypt(cipherText, plainText)
	return cipherText, nil
}

func (s *DESExecuter) Decrypt(cipherText []byte) ([]byte, error) {
	desKey := config.GetDESKey()
	cipherBlock, err := des.NewCipher(desKey)
	if err != nil {
		return nil, err
	}
	
	plainText := make([]byte, len(cipherText))
	cipherBlock.Decrypt(plainText, cipherText)
	return plainText, nil
}