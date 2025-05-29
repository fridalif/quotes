package services

import (
	"quotes/internal/repository"
	model "quotes/pkg/quotes"
)

type QuotesServiceI interface {
	GetQuotes() []model.Quote
	GetQuotesByAuthor(author string) []model.Quote
	GetRandomQuote() model.Quote
	InsertQuote(quote model.Quote) error
	DeleteQuote(id uint) error
}

type QuotesService struct {
	repo repository.QuotesRepoI
}

func NewQuotesService(repo repository.QuotesRepoI) QuotesServiceI {
	return &QuotesService{
		repo: repo,
	}
}

func (s *QuotesService) GetQuotes() []model.Quote {
	return s.repo.GetQuotes()
}
func (s *QuotesService) GetQuotesByAuthor(author string) []model.Quote {
	return s.repo.GetQuotesByAuthor(author)
}
func (s *QuotesService) GetRandomQuote() model.Quote {
	return s.repo.GetRandomQuote()
}
func (s *QuotesService) InsertQuote(quote model.Quote) error {
	return s.repo.InsertQuote(quote)
}
func (s *QuotesService) DeleteQuote(id uint) error {
	return s.repo.DeleteQuote(id)
}
