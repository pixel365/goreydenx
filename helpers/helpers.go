package helpers

import (
	"encoding/json"
	"io"

	rx "github.com/pixel365/goreydenx"
	m "github.com/pixel365/goreydenx/model"
)

func Get[T any](c rx.RxClient, path string) (*m.Result[T], error) {
	res, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	result := m.Result[T]{}
	err = json.Unmarshal(res, &result)
	return &result, err
}

func Post[T any](c rx.RxClient, path string, payload io.Reader) (*m.Result[T], error) {
	res, err := c.Post(path, payload)
	if err != nil {
		return nil, err
	}

	result := m.Result[T]{}
	err = json.Unmarshal(res, &result)
	return &result, err
}
