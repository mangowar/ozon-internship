package server

import (
	"context"
	"fmt"
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
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		links, err := shortener.Shorten(ctx, q.URL.Query().Get("url"))
		if err != nil {
			http.Error(r, "HandleShorten: can't shorten URL", http.StatusInternalServerError)
			return
		}
		short_url := shorten.TransfornLink(q.Host, links.ShortUrl)
		r.WriteHeader(http.StatusOK)
		fmt.Fprint(r, short_url)
	}
}
