package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type redirecter interface {
	Redirect(cont context.Context, shoort_url string) (string, error)
}

type handler interface {
	redirecter
	Shortener
}

func HandleRedirect(red redirecter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			shoort_url := strings.TrimPrefix(r.URL.Path, "/")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			url, err := red.Redirect(ctx, shoort_url)
			if err != nil {
				fmt.Println(err)
			}
			http.Redirect(w, r, url, http.StatusSeeOther)
		}
	}
}
