package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetExchangeValue(t *testing.T) {
	t.Run("Gets correct exchang value from EUR to INR", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/currency-exchange/from/EUR/to/INR", nil)
		res := httptest.NewRecorder()

		ExchangeValueServer(res, req)

		var got ExchangeValue
		want := ExchangeValue{
			ID:                 10002,
			From:               "EUR",
			To:                 "INR",
			ConversionMultiple: 75,
			Port:               8000,
		}

		err := json.NewDecoder(res.Body).Decode(&got)

		if err != nil {
			t.Errorf("got an error decoding json, %v", err)
		}

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
