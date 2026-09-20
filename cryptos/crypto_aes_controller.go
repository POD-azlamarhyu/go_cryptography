package cryptos

import (
	"fmt"
	"log/slog"
)

func CryptoAESController() {
	slog.Info("暗号化のサンプルコードを実行する", slog.String("AES暗号方式","現時点でのデファクトスタンダード"))
	aesExecute := &AESExecuter{}
	aesService := NewAESService(aesExecute)

	var inputText string = "Unko!Unko!"
	fmt.Println("\n\n入力文字列:", inputText)

	cipherText, err := aesService.Encrypt(inputText)
	if err != nil {
		slog.Error("暗号化に失敗しました", slog.String("error", err.Error()))
		return
	}

	decryptedText, err := aesService.Decrypt(cipherText)
	if err != nil {
		slog.Error("復号化に失敗しました", slog.String("error", err.Error()))
		return
	}
	fmt.Println("復号化結果:", string(decryptedText))
	slog.Info("復号化に成功しました")
}