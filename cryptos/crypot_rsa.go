package cryptos

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"fmt"
)

type RSAKey struct {
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
}

var key = &RSAKey{}

func Encrypt(plainText string) ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		fmt.Println("秘密鍵の生成に失敗しました")
		return nil, err
	}

	key.privateKey = privateKey
	key.publicKey = &privateKey.PublicKey
	cipherText, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, key.publicKey, []byte(plainText), nil)
	if err != nil {
		fmt.Println("暗号化に失敗しました")
		return nil, err
	}
	return cipherText, nil
}

func Decrypt(cipherText string) ([]byte, error) {
	plainText, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, key.privateKey, []byte(cipherText), nil)
	if err != nil {
		fmt.Println("復号に失敗しました")
		return nil, err
	}
	return plainText, nil
}