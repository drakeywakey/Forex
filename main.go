package main

import (
	"encoding/json"
	"net/http"
)

type ExchangeValue struct {
	ID                 int64
	From               string
	To                 string
	ConversionMultiple float32
	Port               int64
}

func ExchangeValueServer(w http.ResponseWriter, r *http.Request) {
	val := ExchangeValue{ID: 10002, From: "EUR", To: "INR", ConversionMultiple: 75, Port: 8000}
	json.NewEncoder(w).Encode(val)
	w.WriteHeader(http.StatusOK)
}

func main() {

}
