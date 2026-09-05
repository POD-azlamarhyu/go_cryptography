package cryptos

import (
	"crypto/des"
	"cryptography_tutorial/config"
)

type DESExecuter struct {}

type DESKeyStruct struct{
	shareKey *[]byte
}

var DESkey = &AESKeyStruct{}

func (s *DESExecuter) Encrypt(plainText []byte) ([]byte, error) {
	desKey := config.GetDESKey()
	DESkey.shareKey = &desKey
	cipherBlock, err := des.NewCipher(*DESkey.shareKey)
	if err != nil {
		return nil, err
	}
	
	cipherText := make([]byte, len(plainText))
	cipherBlock.Encrypt(cipherText, plainText)
	return cipherText, nil
}

func (s *DESExecuter) Decrypt(cipherText []byte) ([]byte, error) {

	cipherBlock, err := des.NewCipher(*DESkey.shareKey)
	if err != nil {
		return nil, err
	}
	
	plainText := make([]byte, len(cipherText))
	cipherBlock.Decrypt(plainText, cipherText)
	return plainText, nil
}