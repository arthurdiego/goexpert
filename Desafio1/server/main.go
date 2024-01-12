package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type BodyJSON struct {
	Quotation Quotation `json:"USDBRL"`
}

type Quotation struct {
	gorm.Model
	Bid string `json:"bid"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("desafio1.sqlite"), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro ao abrir o banco de dados:", err)
		panic(err)
	}
	err = db.AutoMigrate(&Quotation{})
	if err != nil {
		fmt.Println("Erro ao executar AutoMigrate:", err)
		panic(err)
	}

	http.HandleFunc("/cotacao", handleCotacao)
	http.ListenAndServe(":8080", nil)
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	exchangeRate, err := getExchangeRate(ctx)
	if err != nil {
		fmt.Println("Erro ao obter a cotação:", err)
		panic(err)
	}

	// Salvar a cotação no banco de dados
	db.Create(&exchangeRate)

	// Responder com a cotação em formato JSON
	jsonResponse, err := json.Marshal(exchangeRate)
	if err != nil {
		fmt.Println("Erro ao converter a cotação para JSON:", err)
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getExchangeRate(ctx context.Context) (Quotation, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		fmt.Println("Erro ao criar a requisição:", err)
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erro: resposta do servidor não foi OK")
		panic(err)
	}

	var result BodyJSON

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Erro ao decodificar a resposta JSON:", err)
		panic(err)
	}

	return result.Quotation, nil
}
