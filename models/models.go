package models

type QuoteModel struct {
	ID     int    `db:"id"`
	Author string `db:"author"`
	Quote  string `db:"quote"`
}
