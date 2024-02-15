package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/arthurdiego/goexpert/desafio2/internal/entity"
	"github.com/arthurdiego/goexpert/desafio2/internal/usecases"
)

const timeout = 1 * time.Second

func main() {
	var cep string

	// Configurar flags
	flag.StringVar(&cep, "cep", "", "CEP para consulta")
	flag.Parse()

	if cep == "" {
		fmt.Println("Por favor, forneça um CEP como argumento usando -cep.")
		return
	}

	canal1 := make(chan *entity.Response1)
	canal2 := make(chan *entity.Response2)

	// Embaralhar a ordem de requisição para variar a ordem de resposta
	randomNumber := rand.Intn(3)
	timeout1, timeout2 := 3, 3
	switch randomNumber {
	case 0:
		timeout1 = 0
	case 1:
		timeout2 = 0
	}

	go usecases.GetCEP1(cep, canal1, time.Duration(timeout1)*time.Second)
	go usecases.GetCEP2(cep, canal2, time.Duration(timeout2)*time.Second)

	select {
	case result := <-canal1:
		if result != nil {
			fmt.Println("Resposta da Brasil API:")
			fmt.Printf("CEP: %s\n", result.Cep)
			fmt.Printf("Rua: %s\n", result.Street)
			fmt.Printf("Bairro: %s\n", result.Neighborhood)
			fmt.Printf("Cidade: %s\n", result.City)
			fmt.Printf("Estado: %s\n", result.State)
		}
	case result := <-canal2:
		if result != nil {
			fmt.Println("Resposta da Via Cep API:")
			fmt.Printf("CEP do cliente: %s\n", result.Cep)
			fmt.Printf("Rua: %s\n", result.Logradouro)
			fmt.Printf("Bairro: %s\n", result.Bairro)
			fmt.Printf("Cidade: %s\n", result.Localidade)
			fmt.Printf("Estado: %s\n", result.UF)

		}
	case <-time.After(timeout):
		fmt.Println("Erro: Timeout na requisição.")
	}
}
