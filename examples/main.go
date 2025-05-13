package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/ricksantos88/go-errorx"
)

func main() {
	databaseExample()
	apiExample()
	businessLogicExample()
}

func databaseExample() {
	fmt.Println("\n--- Exemplo Banco de Dados ---")

	err := errorx.Try("db.query", func() error {
		// Simular erro de banco de dados
		if rand.Intn(2) == 0 {
			return fmt.Errorf("connection timeout")
		}
		return nil
	})

	if err != nil {
		log.Printf("Erro de banco de dados: %v", err)
		fmt.Printf("Contexto: %v\n", errorx.GetContext(err))
	}
}

func apiExample() {
	fmt.Println("\n--- Exemplo API ---")

	h := errorx.New("api.request").
		With("endpoint", "/users").
		With("method", "GET")

	// Simular chamada API
	_, err := mockAPICall()
	if h.Check(err) {
		fmt.Println("Erro na chamada API foi tratado")
		fmt.Printf("Operação: %s\n", errorx.GetOperation(err))
	}
}

func mockAPICall() (interface{}, error) {
	// Simular erro aleatório
	if rand.Intn(2) == 0 {
		return nil, fmt.Errorf("status 500 - internal server error")
	}
	return "response data", nil
}

func businessLogicExample() {
	fmt.Println("\n--- Exemplo Lógica de Negócios ---")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de panic:", r)
		}
	}()

	errorx.MustDo("process.payment", func() error {
		// Simular validação de pagamento
		if rand.Intn(2) == 0 {
			return fmt.Errorf("invalid credit card")
		}
		return nil
	})

	fmt.Println("Pagamento processado com sucesso")
}
