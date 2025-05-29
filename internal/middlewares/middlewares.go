package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
)

type MiddlewaresI interface {
	RecoverMiddleware(next http.HandlerFunc) http.HandlerFunc
}

type Middlewares struct {
	host string
	port string
}

func NewMiddlewares(host, port string) MiddlewaresI {
	return &Middlewares{
		host: host,
		port: port,
	}
}

func (middlewares *Middlewares) RecoverMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("%s:%s Произошла паника: %v", middlewares.host, middlewares.port, err)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Внутренняя ошибка сервера",
				})
			}
		}()

		next(w, r)
	}
}
