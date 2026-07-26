package cryptos

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"cryptography_tutorial/config"
	"fmt"
	"io"
	"log"
)

type AESExecuter struct {}

func (s *AESExecuter) Encrypt(plainText []byte) ([]byte, error){
	aesKey := config.GetAESKey()
	block, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil{
		log.Fatalln("エラーが発生")
		return nil, err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plainText, nil)
	fmt.Printf("暗号化した文字列：%x\n\n", ciphertext)
	return ciphertext,nil
}

func (s *AESExecuter) Decrypt(cipherText []byte) ([]byte, error){
	aesKey := config.GetAESKey()
	block, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return nil, err
	}
	
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	
	nonceSize := aesgcm.NonceSize()
	if len(cipherText) < nonceSize {
		log.Fatalln("不正な暗号文です")
		return nil, err
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("復号化した文字列：%s\n\n", plaintext)
	return plaintext, nil
}