package storage

import (
	"bufio"
	"context"
	"errors"
	"os"
	"shortener/config"
	"shortener/model"
	"strings"
)

type InternalMemory struct{}

func (i InternalMemory) Contains(cont context.Context, short_url string) (string, bool, error) {
	var (
		value string
		count int = 0
	)
	file, err := os.OpenFile(config.Get().FileName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", false, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), " ")
		if pair[0] == short_url {
			count++
			value = pair[1]
		}
	}
	if count == 0 {
		return "", false, nil
	} else if count == 1 {
		// fmt.Println(key, value)
		return value, true, nil
	}
	return "", false, errors.New(MoreThanOneError)
}

func (i InternalMemory) Insert(cont context.Context, input model.Links) error {
	if _, contains, err := i.Contains(cont, input.ShortUrl); err != nil {
		return err
	} else {
		if !contains {
			if file, err := os.OpenFile(config.Get().FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644); err != nil {
				return err
			} else {
				defer file.Close()
				file.WriteString(input.ShortUrl + " " + input.Url + "\n")
			}

		} else {
			return errors.New(AlreadyContainsError)
		}
	}
	return nil
}
