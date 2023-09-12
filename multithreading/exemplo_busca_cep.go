package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func fetchCEP(apiURL string, wg *sync.WaitGroup, resultChan chan []byte) {
	defer wg.Done()

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Erro ao fazer a requisição para %s: %s\n", apiURL, err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Erro ao ler a resposta de %s: %s\n", apiURL, err)
		return
	}

	resultChan <- body
}

func main() {
	cep := "06634-000"
	api1URL := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	api2URL := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	var wg sync.WaitGroup
	resultChan := make(chan []byte, 2)

	wg.Add(2)
	go fetchCEP(api1URL, &wg, resultChan)
	go fetchCEP(api2URL, &wg, resultChan)

	select {
	case result := <-resultChan:
		fmt.Printf("Resposta da API: %s\n", result)
	case <-time.After(5 * time.Second): // Defina um tempo limite aqui
		fmt.Println("Nenhuma resposta recebida dentro do tempo limite.")
	}

	close(resultChan)
	wg.Wait()
}
