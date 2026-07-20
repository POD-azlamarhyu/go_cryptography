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
	desServiceInterface := cryptos.NewDESService(desService)
	var inputText string = "Hello" // デフォルトの入力文字列を設定
	// fmt.Printf("暗号化する文字列：")
	// fmt.Scanln(&inputText)

	fmt.Println("\n\n入力文字列:", inputText)
	plainText := strings.FormatByteStringDES(inputText)
	fmt.Println("入力文字列:", plainText)
	fmt.Println("入力文字列:", []byte(inputText))

	cipherText, err := desServiceInterface.Encrypt(plainText)
	if err != nil {
		slog.Error("暗号化に失敗しました", slog.String("error", err.Error()))
		return
	}
	// slog.Info("暗号化に成功しました", slog.String("cipherText", string(cipherText)))

	decryptedText, err := desServiceInterface.Decrypt(cipherText)
	if err != nil {
		slog.Error("復号化に失敗しました", slog.String("error", err.Error()))
		return
	}

	//TODO: 64ビットずつ暗号化と復号化を行うので、任意の文字列数ではなく、64ビット（8バイト）の倍数の文字列を入力する必要がある。
	fmt.Println("\n\n暗号化後の文字列:", cipherText,string(cipherText))
	fmt.Println("復号化後の文字列:", decryptedText,string(decryptedText))
}