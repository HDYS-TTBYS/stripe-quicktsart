package main

import (
	"net/http"

	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/price"
)

func handleGetPrices(w http.ResponseWriter, r *http.Request) {
	setCors(w, r)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// すべての価格を一覧表示する
	params := &stripe.PriceListParams{}
	params.Filters.AddFilter("limit", "", "10")
	i := price.List(params)
	priceList := []*stripe.Price{}
	for i.Next() {
		p := i.Price()
		priceList = append(priceList, p)
	}

	writeJSON(w, priceList)
}

func handleGetPrice(w http.ResponseWriter, r *http.Request) {
	setCors(w, r)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")
	p, err := price.Get(id, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeJSON(w, p)
}
