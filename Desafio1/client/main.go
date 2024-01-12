package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Quotation struct {
	Bid string `json:"bid"`
}

func main() {
	// Criar um contexto com timeout de 300ms
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	// Criar cliente HTTP
	client := &http.Client{}

	// Criar uma nova requisição HTTP com o contexto definido
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Erro ao criar a requisição:", err)
		panic(err)
	}

	// Realizar a requisição HTTP
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		panic(err)
	}
	defer resp.Body.Close()

	// Decodificar a resposta JSON em uma estrutura de dados
	var quotation Quotation
	err = json.NewDecoder(resp.Body).Decode(&quotation)
	if err != nil {
		fmt.Println("Erro ao decodificar a resposta JSON:", err)
		panic(err)
	}

	// Escrever o conteúdo no arquivo
	err = writeToFile("cotacao.txt", fmt.Sprintf("Dólar: {%s}", quotation.Bid))
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		panic(err)
	}
}

func writeToFile(filename, content string) error {
	// Criar ou truncar o arquivo para escrita
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		panic(err)
	}

	// Escrever o conteúdo no arquivo
	_, err = f.Write([]byte(content))
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		panic(err)
	}

	// Fechar o arquivo
	f.Close()

	return err
}
