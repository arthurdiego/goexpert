package usecases

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/arthurdiego/goexpert/desafio2/internal/entity"
)

func GetCEP1(cep string, canal chan<- *entity.Response1, timeout time.Duration) {
	req, err := http.NewRequest(http.MethodGet, "https://brasilapi.com.br/api/cep/v1/"+cep, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao enviar a requisição: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Status Code: %d - %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("body:", string(body))
		return
	}

	var data entity.Response1
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Erro ao decodificar a resposta JSON:", err)
		return
	}
	if timeout > 0 {
		time.Sleep(timeout)
	}
	canal <- &data
	close(canal)
}
