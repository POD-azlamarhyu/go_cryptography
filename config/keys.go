package config

import (
	"os"
	"log/slog"
	"crypto/rand"
)

const (
	DES_KEY = "DES_KEY"
	AES_KEY = "AES_KEY"
)

func GetDESKeyName() string {
	return DES_KEY
}

func GetAESKeyName() string {
	return AES_KEY
}

func GetAESKey() string {
	return os.Getenv(GetAESKeyName())
}

func GetDESKey() []byte {
	key, err := os.LookupEnv(GetDESKeyName())
	if !err || len(key) != 8 {
		// slog.Error("DES_KEYが環境変数に設定されていません、または無効な長さです。")
		randKey := make([]byte, 8) // DESキーは64bit（8バイト）である必要があります

		_, err := rand.Read(randKey)
		if err != nil {
			slog.Error("ランダムなDESキーの生成に失敗しました。", slog.String("error", err.Error()))
			return nil
		}

		// slog.Info("ランダムなDESキーを生成しました。", slog.String("DES_KEY", string(randKey)))
		os.Setenv(GetDESKeyName(), string(randKey)) // 環境変数に設定
		return randKey
	}
	return []byte(key)
}