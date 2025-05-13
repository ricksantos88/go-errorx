package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ricksantos88/go-errorx"
)

func main() {
	exampleCheck()
	exampleMust()
	exampleTry()
	exampleMustDo()
}

func exampleCheck() {
	fmt.Println("\n--- Exemplo Check ---")
	h := errorx.NewHandler("leitura de arquivo")
	
	// Simular erro
	_, err := os.Open("arquivo_inexistente.txt")
	if h.Check(err) {
		fmt.Println("Erro foi tratado e logado")
	}

	// Com contexto adicional
	h.WithContext("tentativa", 3).WithContext("modo", "leitura")
	_, err = os.Open("outro_arquivo_inexistente.txt")
	if h.Check(err) {
		fmt.Println("Erro com contexto adicional foi tratado")
	}
}

func exampleMust() {
	fmt.Println("\n--- Exemplo Must ---")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	h := errorx.NewHandler("operação crítica")
	
	// Isso vai causar panic
	_, err := os.Open("arquivo_inexistente.txt")
	h.Must(err)
}

func exampleTry() {
	fmt.Println("\n--- Exemplo Try ---")
	
	err := errorx.Try("processamento complexo", func() error {
		// Simular várias operações com erro
		_, err := os.Open("arquivo_inexistente.txt")
		if err != nil {
			return err
		}
		return nil
	})
	
	if err != nil {
		log.Printf("Erro tratado: %v\n", err)
	}
}

func exampleMustDo() {
	fmt.Println("\n--- Exemplo MustDo ---")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Isso vai causar panic se o arquivo não existir
	errorx.MustDo("escrever arquivo", func() error {
		return os.WriteFile("teste.txt", []byte("dados"), 0644)
	})
	
	fmt.Println("Operação concluída com sucesso")
}