package main

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/inancgumus/screen"
	"github.com/pastorilps/password-generator/password"
)

func main() {

	var (
		quit   string
		senhas int
	)

	for quit != "n" {
		screen.Clear()
		fmt.Println("Quantas senhas você deseja gerar?")
		fmt.Scanln(&senhas)
		fmt.Println("\n")

		fmt.Println("Digite o tamanho da senha desejada (8 ou mais):")
		var sizepass int
		fmt.Scanln(&sizepass)
		fmt.Println("\n")

		if sizepass < 8 {
			fmt.Println("O tamanho mínimo da senha é 8.\n")
			continue
		}

		for i := 0; i < senhas; i++ {
			pass := password.GeneratePassword(sizepass)
			cryptoPass := password.GetMd5Hash(pass)

			fmt.Println("Password: ", pass)
			fmt.Println("Encrypted Password: ", cryptoPass, "\n")
			copyToClipboard(pass)
			time.Sleep(2 * time.Second)

		}

		fmt.Println("\n", "Deseja criar mais senhas? y/n\n")

		fmt.Scanln(&quit)

	}

}

func copyToClipboard(text string) {
	err := clipboard.WriteAll(text)
	if err != nil {
		fmt.Println("Erro ao copiar para a área de transferência", err)
	}
}
