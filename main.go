package main

import (
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v72"
)

func main() {
	// This is your test secret API key.
	stripe.Key = os.Getenv("STRIPE_SECRET_API_KEY")

	fs := http.FileServer(http.Dir("build"))
	http.Handle("/", fs)
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/products", handleGetProducts)
	http.HandleFunc("/prices", handleGetPrices)
	http.HandleFunc("/price", handleGetPrice)
	http.HandleFunc("/create-customer", handleCreateCustomer)
	http.HandleFunc("/delete-customer", handleDeleteCustomer)
	http.HandleFunc("/customer", handleGetCustomer)
	http.HandleFunc("/customers", handleGetCustomers)

	addr := "localhost:4242"
	log.Printf("Listening on %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
