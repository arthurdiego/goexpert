package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/arthurdiego/goexpert/desafio2/internal/entity"
)

func GetCEP2(cep string, canal chan<- *entity.Response2, timeout time.Duration) {
	req, err := http.NewRequest(http.MethodGet, "http://viacep.com.br/ws/"+cep+"/json", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar a requisição: %v\n", err)
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
		return
	}

	var data entity.Response2
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
