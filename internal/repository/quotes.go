package repository

import (
	"math/rand"
	"quotes/pkg/errors"
	model "quotes/pkg/quotes"
	"sync"
)

/*
Здесь должна быть логика работы с БД, но по ТЗ в памяти можно
*/
type QuotesRepoI interface {
	GetQuotes() []model.Quote
	GetQuotesByAuthor(author string) []model.Quote
	GetRandomQuote() model.Quote
	InsertQuote(quote model.Quote) error
	DeleteQuote(id uint) error
}

type QuotesRepo struct {
	quotes []model.Quote
	mu     *sync.RWMutex
}

func NewQuotesRepo(quotes []model.Quote) QuotesRepoI {
	return &QuotesRepo{
		quotes: quotes,
		mu:     &sync.RWMutex{},
	}
}

func (r *QuotesRepo) GetQuotes() []model.Quote {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.quotes
}
func (r *QuotesRepo) GetQuotesByAuthor(author string) []model.Quote {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var quotes []model.Quote
	for _, quote := range r.quotes {
		if quote.Author == author {
			quotes = append(quotes, quote)
		}
	}
	return quotes
}
func (r *QuotesRepo) GetRandomQuote() model.Quote {
	r.mu.RLock()
	defer r.mu.RUnlock()
	quote := r.quotes[rand.Intn(len(r.quotes))]
	return quote
}
func (r *QuotesRepo) InsertQuote(quote model.Quote) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, q := range r.quotes {
		if q.Text == quote.Text && q.Author == quote.Author {
			return errors.ErrQuoteAlreadyExists
		}
	}
	quote.Id = uint(len(r.quotes) + 1)
	r.quotes = append(r.quotes, quote)
	return nil
}

func (r *QuotesRepo) DeleteQuote(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, quote := range r.quotes {
		if quote.Id == id {
			r.quotes = append(r.quotes[:i], r.quotes[i+1:]...)
			return nil
		}
	}
	return errors.ErrQuoteNotFound
}
