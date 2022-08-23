package main

import (
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/product"
)

func handleGetProducts(w http.ResponseWriter, r *http.Request) {
	setCors(w, r)

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// すべての製品を一覧表示する
	params := &stripe.ProductListParams{}
	params.Filters.AddFilter("limit", "", "")
	i := product.List(params)
	productsList := []*stripe.Product{}
	for i.Next() {
		p := i.Product()
		productsList = append(productsList, p)
		pp.Print(p)
	}

	writeJSON(w, productsList)
}
