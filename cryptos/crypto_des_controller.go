package cryptos

import (
	"cryptography_tutorial/strings"
	"fmt"
	"log/slog"
)

func CryptoDESController() {
	slog.Info("暗号化のサンプルコードを実行する", slog.String("DES暗号方式","非推奨"))
	desExecute := &DESExecuter{}
	desService := NewDESService(desExecute)

	var inputText string = "Hello"
	fmt.Println("\n\n入力文字列:", inputText)
	plainText := strings.FormatByteStringDES(inputText)

	cipherText, err := desService.Encrypt(plainText)
	if err != nil {
		slog.Error("暗号化に失敗しました", slog.String("error", err.Error()))
		return
	}

	decryptedText, err := desService.Decrypt(cipherText)
	if err != nil {
		slog.Error("復号化に失敗しました", slog.String("error", err.Error()))
		return
	}
	//TODO: 64ビットずつ暗号化と復号化を行うので、任意の文字列数ではなく、64ビット（8バイト）の倍数の文字列を入力する必要がある。
	fmt.Println("復号化結果:", string(decryptedText))
}