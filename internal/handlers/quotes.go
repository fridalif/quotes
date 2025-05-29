package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"quotes/internal/services"
	model "quotes/pkg/quotes"
	"strconv"

	"github.com/gorilla/mux"
)

type QuotesHandlerI interface {
	GetQuotes(w http.ResponseWriter, r *http.Request)
	GetQuotesByAuthor(w http.ResponseWriter, r *http.Request)
	GetRandomQuote(w http.ResponseWriter, r *http.Request)
	InsertQuote(w http.ResponseWriter, r *http.Request)
	DeleteQuote(w http.ResponseWriter, r *http.Request)
}

type QuotesHandler struct {
	service services.QuotesServiceI
}

func NewQuotesHandler(service services.QuotesServiceI) QuotesHandlerI {
	return &QuotesHandler{
		service: service,
	}
}

func (handler *QuotesHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	quotes := handler.service.GetQuotes()
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"body": map[string]interface{}{
			"quotes": quotes,
		},
		"error": nil,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"body":  map[string][]string{},
			"error": "Что-то пошло не так",
		})
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *QuotesHandler) GetQuotesByAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	author := r.URL.Query().Get("author")
	quotes := handler.service.GetQuotesByAuthor(author)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"body": map[string]interface{}{
			"quotes": quotes,
		},
		"error": nil,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"body":  map[string][]string{},
			"error": "Что-то пошло не так",
		})
		log.Println(err)
		return
	}
}

func (handler *QuotesHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	quote := handler.service.GetRandomQuote()
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"body": map[string]interface{}{
			"quotes": []model.Quote{quote},
		},
		"error": nil,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"body":  map[string][]string{},
			"error": "Что-то пошло не так",
		})
		log.Println(err)
		return
	}
}

func (handler *QuotesHandler) InsertQuote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var quote model.Quote
	err := json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"body":  map[string][]string{},
			"error": "Неверный формат запроса",
		})
		return
	}
	if quote.Author == "" || quote.Text == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"body":  map[string][]string{},
			"error": "Неверный формат запроса",
		})
		return
	}
	err = handler.service.InsertQuote(quote)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"body":  map[string][]string{},
			"error": "Цитата уже существует",
		})
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (handler *QuotesHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"body":  map[string][]string{},
			"error": "Неверный формат запроса",
		})
		return
	}
	err = handler.service.DeleteQuote(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"body":  map[string][]string{},
			"error": "Цитата не найдена",
		})
		return
	}
	w.WriteHeader(http.StatusOK)
}
