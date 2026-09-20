package main

import (
	"fmt"
	"cryptography_tutorial/cryptos"
	"log/slog"
	"os"
)

func main() {
	args := os.Args
	slog.Info("Go cryptography cli running.",slog.Any("args",args))

	if len(args) < 2 || len(args) > 2{
		fmt.Println("Usage: go_cryptography <command>")
		fmt.Println("help - Show this help message")
		return
	}

	switch args[1] {
	case "des":
		cryptos.CryptoDESController()
	case "aes":
		cryptos.CryptoAESController()
	case "rsa":
		cryptos.RSAController()
	case "hash":
		cryptos.HashController()
	case "help":
		fmt.Println("Usage: go_cryptography <command>")
		fmt.Println("des - Run DES encryption sample")
		fmt.Println("aes - Run AES encryption sample")
		fmt.Println("rsa - Run RSA encryption sample")
		fmt.Println("hash - Run hash sample")
	default:
		slog.Error("Unknown command", slog.String("command", args[1]))
		return
	}

	slog.Info("Go cryptography cli finished.")
}