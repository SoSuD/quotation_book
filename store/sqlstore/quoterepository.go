package sqlstore

import (
	"context"
	"quotation_book/models"
)

type quoteRepository struct {
	store *Store
}

func (q *quoteRepository) AddQuote(ctx context.Context, quote *models.QuoteModel) (models.QuoteModel, error) {
	query := `INSERT INTO quotes (author, quote) VALUES ($1, $2) RETURNING id`
	ret := models.QuoteModel{}
	if err := q.store.db.QueryRow(ctx, query, quote.Author, quote.Quote).Scan(&ret.ID); err != nil {
		return models.QuoteModel{}, err
	}
	return ret, nil
}

func (q *quoteRepository) GetQuotes(ctx context.Context, author string) ([]models.QuoteModel, error) {
	sql := `SELECT id, quote, author FROM quotes`
	var args []interface{}

	// Если передан автор — добавляем WHERE и параметр
	if author != "" {
		sql += ` WHERE author = $1`
		args = append(args, author)
	}

	// Выполняем запрос
	rows, err := q.store.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Сканируем в срез
	var result []models.QuoteModel
	for rows.Next() {
		var q models.QuoteModel
		if err := rows.Scan(&q.ID, &q.Quote, &q.Author); err != nil {
			return nil, err
		}
		result = append(result, q)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return result, nil
}

func (q *quoteRepository) GetRandomQuote(ctx context.Context) (models.QuoteModel, error) {
	query := `SELECT id, quote, author FROM quotes ORDER BY random() LIMIT 1`
	ret := models.QuoteModel{}
	if err := q.store.db.QueryRow(ctx, query).Scan(&ret.ID, &ret.Quote, &ret.Author); err != nil {
		return models.QuoteModel{}, err
	}
	return ret, nil
}

func (q *quoteRepository) DeleteQuote(ctx context.Context, id int) error {
	query := `DELETE FROM quotes WHERE id = $1`
	_, err := q.store.db.Exec(ctx, query, id)
	return err
}
