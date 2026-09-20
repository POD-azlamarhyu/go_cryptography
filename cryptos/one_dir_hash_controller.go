package cryptos

import "log/slog"

func HashController(){
	slog.Info("ハッシュのサンプルコードを実行する",slog.String("SHA2 256, SHA2 512","デファクトスタンダード"))

	hashExecuter := &hashExecuter{}
	hashService := NewHashService(hashExecuter)

	var message string = "Hello, World!"
	hashedText, err := hashService.ExecuteSHA256(message)
	if err != nil {
		slog.Error("ハッシュの実行に失敗しました", slog.String("error", err.Error()))
		return
	}
	slog.Info("ハッシュの実行に成功しました", slog.Any("hashedText", string(hashedText)))
}