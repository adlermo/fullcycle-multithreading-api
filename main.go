package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Result struct {
	API  string
	Data map[string]interface{}
	Err  error
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <CEP>")
		return
	}

	cep := os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan Result)

	go fetchBrasilAPI(ctx, cep, ch)
	go fetchViaCEP(ctx, cep, ch)

	select {
	case res := <-ch:
		if res.Err != nil {
			fmt.Println("Erro:", res.Err)
			return
		}

		fmt.Println("API vencedora:", res.API)
		fmt.Println("Resposta:", res.Data)

	case <-ctx.Done():
		fmt.Println("Timeout: nenhuma API respondeu em 1s")
	}
}

func fetchBrasilAPI(ctx context.Context, cep string, ch chan<- Result) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- Result{API: "BrasilAPI", Err: err}
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	ch <- Result{
		API:  "BrasilAPI",
		Data: data,
	}
}

func fetchViaCEP(ctx context.Context, cep string, ch chan<- Result) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- Result{API: "ViaCEP", Err: err}
		return
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)

	ch <- Result{
		API:  "ViaCEP",
		Data: data,
	}
}
