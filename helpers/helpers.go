package helpers

import (
	"github.com/pixel365/goreydenx"
	m "github.com/pixel365/goreydenx/model"
	"io"
)

func Get[T any](c *goreydenx.Client, path string) (*m.Result[T], error) {
	res, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	result := m.Result[T]{}
	err = c.JSONDecoder(res, &result)
	return &result, err
}

func Post[T any](c *goreydenx.Client, path string, payload io.Reader) (*m.Result[T], error) {
	res, err := c.Post(path, payload)
	if err != nil {
		return nil, err
	}

	result := m.Result[T]{}
	err = c.JSONDecoder(res, &result)
	return &result, err
}
