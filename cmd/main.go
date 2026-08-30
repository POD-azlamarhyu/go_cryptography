package main

import (
	"fmt"
	"cryptography_tutorial/cryptos"
	"log/slog"
	"cryptography_tutorial/strings"
)

func main() {
	slog.Info("暗号化のサンプルコードを実行する", slog.String("DES暗号方式","非推奨"))

	desService := &cryptos.DESExecuter{}
	aesService := &cryptos.AESExecuter{}
	rsaService := cryptos.NewRSAService()
	aesServiceInterface := cryptos.NewAESService(aesService)
	desServiceInterface := cryptos.NewDESService(desService)
	var inputText string = "Hello" // デフォルトの入力文字列を設定

	fmt.Println("\n\n入力文字列:", inputText)
	plainText := strings.FormatByteStringDES(inputText)

	cipherText, err := desServiceInterface.Encrypt(plainText)
	if err != nil {
		slog.Error("暗号化に失敗しました", slog.String("error", err.Error()))
		return
	}

	decryptedText, err := desServiceInterface.Decrypt(cipherText)
	if err != nil {
		slog.Error("復号化に失敗しました", slog.String("error", err.Error()))
		return
	}

	//TODO: 64ビットずつ暗号化と復号化を行うので、任意の文字列数ではなく、64ビット（8バイト）の倍数の文字列を入力する必要がある。
	fmt.Println("\n暗号化後の文字列:",string(cipherText))
	fmt.Println("復号化後の文字列:",string(decryptedText))

	slog.Info("暗号化のサンプルコードを実行する", slog.String("AES暗号方式","現時点でのデファクトスタンダード"))
	inputTextAes := "Unko!"
	fmt.Println("入力文字列: ",inputTextAes)
	cipherTextAes,err := aesServiceInterface.Encrypt(inputTextAes)
	if err != nil{
		slog.Error("暗号化に失敗しました", slog.String("error:", err.Error()))
	}
	decryptedTextAes,err := aesServiceInterface.Decrypt(cipherTextAes)
	if err != nil{
		slog.Error("エラーが発生", slog.String("error:", err.Error()))
	}
	fmt.Println("\n暗号化後の文字列:", string(cipherTextAes))
	fmt.Println("復号化後の文字列:", string(decryptedTextAes))

	slog.Info("暗号化のサンプルコードを実行する", slog.String("RSA暗号方式","現時点ではまだ使えるが4098ビット推奨"))

	inputTextRsa := "Unko!!!!!!"
	fmt.Println("入力文字列: ",inputTextRsa)
	cipherTextRsa, err := rsaService.Encrypt(inputTextRsa)
	decryptedTextRsa, err := rsaService.Decrypt(cipherTextRsa)

	fmt.Println("\n暗号化後の文字列:", string(cipherTextRsa))
	fmt.Println("復号化後の文字列:", string(decryptedTextRsa))
}