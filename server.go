package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

type Cotacao struct {
	ID   int `gorm:"primaryKey"`
	rate CurrencyRate
	gorm.Model
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("./data/cotacao.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening database:", err)
		panic(err)
	}

	if err := db.AutoMigrate(&Cotacao{}); err != nil {
		fmt.Println("Error migrating database:", err)
		panic(err)
	}

	http.HandleFunc("/cotacao", BuscaCotacaoHandler)
	http.ListenAndServe(":8080", nil)

}

func BuscaCotacaoHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Millisecond)
	defer cancel()

	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cotacao, err := BuscaCotacao()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error fetching data: %v", err)
		return
	}

	if err := db.WithContext(ctx).Create(&Cotacao{rate: cotacao.USDBRL}).Error; err != nil {
		fmt.Println("Error saving cotacao:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cotacao)

}

func BuscaCotacao() (*ApiResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	cotacao, error := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if error != nil {
		fmt.Println("Error fetching data:", error)
		return nil, error
	}

	res, err := http.DefaultClient.Do(cotacao)
	if err != nil {
		fmt.Println("Error performing request:", err)
		return nil, err
	}

	defer res.Body.Close()
	body, error := io.ReadAll(res.Body)
	if error != nil {
		fmt.Println("Error reading response body:", error)
		return nil, error
	}

	var apiResponse ApiResponse
	error = json.Unmarshal(body, &apiResponse)
	if error != nil {
		fmt.Println("Error unmarshalling JSON:", error)
		return nil, error
	}

	return &apiResponse, nil

}
