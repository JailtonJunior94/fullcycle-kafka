package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jailtonjunior94/fullcycle-kafka/cmd/api/configs"
	"github.com/jailtonjunior94/fullcycle-kafka/cmd/api/handlers"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/customers", handlers.CreateCustomer)
	r.Post("/transactions", handlers.CreateTransaction)

	http.ListenAndServe(":9000", r)
}
