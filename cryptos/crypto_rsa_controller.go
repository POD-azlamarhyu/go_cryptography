package cryptos

import (
	"fmt"
	"log/slog"
)

func RSAController(){
	slog.Info("暗号化のサンプルコードを実行する", slog.String("RSA暗号方式","現時点ではまだ使えるが4098ビット推奨"))
	inputTextRsa := "Unko!!!!!!"
	fmt.Println("入力文字列: ",inputTextRsa)

	rsaExecuter := &RSAExecuter{}
	rsaService := NewRSAService(rsaExecuter)
	cipherText, err := rsaService.Encrypt(inputTextRsa)
	if err != nil {
		fmt.Println("暗号化に失敗しました")
		return
	}
	fmt.Println("暗号化後の文字列: ", cipherText)

	plainText, err := rsaService.Decrypt(cipherText)
	if err != nil {
		fmt.Println("復号に失敗しました")
		return
	}
	fmt.Println("復号後の文字列: ", plainText)
}