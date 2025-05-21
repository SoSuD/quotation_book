package sqlstore

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"quotation_book/store"
)

type Store struct {
	db              *pgxpool.Pool
	quoteRepository *quoteRepository
}

func New(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Quotes() store.QuoteRepository {
	if s.quoteRepository != nil {
		return s.quoteRepository
	}
	s.quoteRepository = &quoteRepository{
		store: s,
	}
	return s.quoteRepository
}
