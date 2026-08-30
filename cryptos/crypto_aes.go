package cryptos

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"cryptography_tutorial/config"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
)

type AESExecuter struct {}

func (s *AESExecuter) Encrypt(plainText string) (string, error){
	aesKey := config.GetAESKey()
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil{
		log.Fatalln("エラーが発生")
		return "", err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, []byte(plainText), nil)
	fmt.Printf("暗号化した文字列：%x\n\n", ciphertext)
	return base64.StdEncoding.EncodeToString(ciphertext),nil
}

func (s *AESExecuter) Decrypt(cipherText string) (string, error){
	aesKey := config.GetAESKey()
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}
	
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		slog.Error("cipher NewGCM")
		return "", err
	}
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil{
		slog.Error("base64 encoding errr")
		return "", err
	}
	nonceSize := aesgcm.NonceSize()
	if len(cipherText) < nonceSize {
		slog.Error("不正な値です")
		return "",errors.New("不正な値")
	}

	nonce, cipherTextByte := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, cipherTextByte, nil)
	if err != nil {
		return "", err
	}
	fmt.Printf("復号化した文字列：%s\n\n", plaintext)
	return string(plaintext), nil
}