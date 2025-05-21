package apiserver

import (
	"encoding/json"
	"net/http"
	"path"
	"quotation_book/models"
	"quotation_book/store"
	"strconv"
)

type server struct {
	router *http.ServeMux
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		store:  store,
		router: http.NewServeMux(),
	}
	return s
}

func (s *server) configureRouter() {
	s.router.HandleFunc("POST /quotes", s.addQuote())
	s.router.HandleFunc("GET /quotes", s.getQuote())
	s.router.HandleFunc("GET /quotes/random", s.getRandomQuote())
	s.router.HandleFunc("DELETE /quotes/", s.deleteQuote())
}

func (s *server) addQuote() http.HandlerFunc {
	type request struct {
		Author string `json:"author"`
		Quote  string `json:"quote"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if req.Author == "" || req.Quote == "" {
			http.Error(w, "Author and quote are required", http.StatusBadRequest)
			return
		}

		quote := &models.QuoteModel{
			Author: req.Author,
			Quote:  req.Quote,
		}

		ret, err := s.store.Quotes().AddQuote(r.Context(), quote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ret)
	}
}

func (s *server) getQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		quotes, err := s.store.Quotes().GetQuotes(r.Context(), author)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quotes)
	}
}

func (s *server) getRandomQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quote, err := s.store.Quotes().GetRandomQuote(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(quote)

	}
}

func (s *server) deleteQuote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := path.Base(r.URL.Path)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := s.store.Quotes().DeleteQuote(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		return
	}
}
