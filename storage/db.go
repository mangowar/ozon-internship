package storage

import (
	"context"
	"errors"
	"fmt"

	"shortener/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	MoreThanOneError     = "more than one instance"
	AlreadyContainsError = "already contains such url"
)

type DataBase struct {
	db *sqlx.DB
}

func New(config string) (DataBase, error) {
	var err error
	var d DataBase
	d.db, err = sqlx.Connect("postgres", config)
	if err != nil {
		return d, err
	}
	d.db.Exec("create table if not exists urls ( short_url text, url text, primary key(short_url) );")
	return d, nil
}

func (d DataBase) Contains(cont context.Context, short_url string) (string, bool, error) {
	var result []model.Links
	err := d.db.Select(&result, fmt.Sprintf("SELECT * FROM urls WHERE short_url = '%s'", short_url))
	if err != nil {
		return "", false, err
	} else {
		if len(result) > 1 {
			return "", false, errors.New(MoreThanOneError)
		} else if len(result) == 1 {
			link_pair := result[0]
			return link_pair.Url, true, nil
		} else {
			return "", false, nil
		}
	}
}

func (d DataBase) Insert(cont context.Context, input model.Links) error {
	if _, contains, err := d.Contains(cont, input.ShortUrl); err != nil {
		return err
	} else {
		if !contains {
			query := fmt.Sprintf("INSERT INTO urls (short_url, url) VALUES ('%s', '%s');", input.ShortUrl, input.Url)
			_, err := d.db.Exec(query)
			if err != nil {
				return fmt.Errorf("can't do insert: %w", err)
			}
		} else {
			return nil
		}
	}
	return nil
}

func (d DataBase) Close() error {
	return d.db.Close()
}
