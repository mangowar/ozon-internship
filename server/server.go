package server

import (
	"context"
	"encoding/json"
	"net/http"
	"shortener/model"
	"shortener/shorten"
	"time"
)

type Shortener interface {
	Shorten(context.Context, string) (model.Links, error)
}

type Request struct {
	URL string
}

type Response struct {
	Short_url string
	// Message   string
}

func HandleShorten(shortener Shortener) http.HandlerFunc {
	return func(r http.ResponseWriter, q *http.Request) {
		if q.Method == http.MethodPost {
			var request Request
			json.NewDecoder(q.Body).Decode(&request)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			links, err := shortener.Shorten(ctx, request.URL)
			if err != nil {
				http.Error(r, "HandleShorten: can't shorten URL", http.StatusInternalServerError)
				return
			}
			short_url := shorten.TransfornLink(q.Host, links.ShortUrl)
			r.WriteHeader(http.StatusOK)
			json.NewEncoder(r).Encode(Response{Short_url: short_url})
		}
	}
}

func MainHandler(h handler) http.HandlerFunc {
	return func(r http.ResponseWriter, q *http.Request) {
		if q.Method == http.MethodPost {
			HandleShorten(h)(r, q)
		} else if q.Method == http.MethodGet {
			HandleRedirect(h)(r, q)
		}
	}
}
