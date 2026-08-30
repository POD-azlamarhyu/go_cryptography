package cryptos

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type RSAKey struct {
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
}

var key = &RSAKey{}

type RSAService struct{}

func NewRSAService()(*RSAService){
	return &RSAService{}
}

func (r *RSAService) Encrypt(plainText string) (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println("秘密鍵の生成に失敗しました")
		return "", err
	}

	key.privateKey = privateKey
	key.publicKey = &privateKey.PublicKey
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, key.publicKey, []byte(plainText), nil)
	if err != nil {
		fmt.Println("暗号化に失敗しました")
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (r *RSAService) Decrypt(cipherText string) (string, error) {
	decodedText, _ := base64.StdEncoding.DecodeString(cipherText)
	plainText, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key.privateKey, decodedText, nil)
	if err != nil {
		fmt.Println("復号に失敗しました")
		return "", err
	}
	return string(plainText), nil
}