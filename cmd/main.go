package main

import (
	"quotes/internal/handlers"
	"quotes/internal/middlewares"
	"quotes/internal/repository"
	"quotes/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	quotesRepo := repository.NewQuotesRepo()
	quotesService := services.NewQuotesService(quotesRepo)
	quotesHandler := handlers.NewQuotesHandler(quotesService)
	middlewares := middlewares.NewMiddlewares("localhost", "8080")

	r := mux.NewRouter()

	r.HandleFunc("/quotes", middlewares.RecoverMiddleware(quotesHandler.GetQuotes)).Methods("GET")
	r.HandleFunc("/quotes/random", middlewares.RecoverMiddleware(quotesHandler.GetRandomQuote)).Methods("GET")
	r.HandleFunc("/quotes", middlewares.RecoverMiddleware(quotesHandler.InsertQuote)).Methods("POST")
	r.HandleFunc("/quotes/{id}", middlewares.RecoverMiddleware(quotesHandler.DeleteQuote)).Methods("DELETE")
}
