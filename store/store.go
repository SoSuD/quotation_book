package store

type Store interface {
	Quotes() QuoteRepository
}
