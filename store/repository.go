package store

import (
	"context"
	"quotation_book/models"
)

type QuoteRepository interface {
	AddQuote(ctx context.Context, quote *models.QuoteModel) (models.QuoteModel, error)
	GetQuotes(ctx context.Context, author string) ([]models.QuoteModel, error)
	GetRandomQuote(ctx context.Context) (models.QuoteModel, error)
	DeleteQuote(ctx context.Context, id int) error
}
