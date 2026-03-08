package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// CurrencyRate representa os campos de uma cotação
type CurrencyRate struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

// ApiResponse é o objecto de topo retornado pela API
type ApiResponse struct {
	USDBRL CurrencyRate `json:"USDBRL"`
}

func main() {
	cotacao, err := BuscaCotacao()
	if err != nil {
		fmt.Println("Error fetching cotacao:", err)
		return
	}

	fileWrite(*cotacao)
}

func BuscaCotacao() (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	cotacao, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(cotacao)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d\n", res.StatusCode)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return &apiResponse.USDBRL.Bid, nil
}

func fileWrite(cotacao string) {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	tamanho, err := file.WriteString("Dólar: " + cotacao)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Printf("Tamanho do arquivo: %d bytes\n", tamanho)

}
