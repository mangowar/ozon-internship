package model

import "context"

type Links struct {
	ShortUrl string `db:"short_url"`
	Url      string `db:"url"`
}

type Shortener interface {
	Shorten(cont context.Context, input string)
}
