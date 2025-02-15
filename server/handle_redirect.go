package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type redirecter interface {
	Redirect(cont context.Context, shoort_url string) (string, error)
}

func HandleRedirect(red redirecter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shoort_url := r.URL.Query().Get("url")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		url, err := red.Redirect(ctx, shoort_url)
		fmt.Println(shoort_url, url)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
