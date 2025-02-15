package shorten

import (
	"context"
	"errors"
	"fmt"
	"shortener/model"
)

type Storage interface {
	Insert(cont context.Context, input model.Links) error
	Contains(cont context.Context, short_url string) (string, bool, error)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) Shorten(cont context.Context, url string) (model.Links, error) {
	input := model.Links{
		ShortUrl: ShortenLink(url),
		Url:      url,
	}
	if err := s.storage.Insert(cont, input); err != nil {
		fmt.Println(err)
		return input, err
	}
	return input, nil
}

func (s *Service) Redirect(cont context.Context, shoort_url string) (string, error) {
	url, contains, err := s.storage.Contains(cont, shoort_url)
	if err != nil {
		return "", err
	} else if !contains {
		return "", errors.New("don't have such url")
	}
	return url, nil
}
