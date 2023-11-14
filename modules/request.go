package modules

import (
	"errors"
	"io"
	"net/http"
)

type CatInt interface {
	GetCatTags(string) ([]byte, error)
}

type Cat struct{}

func (c *Cat) GetCatTags(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
