package user

import (
	rx "github.com/pixel365/goreydenx"
	m "github.com/pixel365/goreydenx/model"
)

// Account User Account
// See https://api.reyden-x.com/docs#/User/get_user_v1_user__get
func Account(c *rx.Client) (*m.User, error) {
	res, err := c.Get("/user/")
	if err != nil {
		return nil, err
	}

	result := m.User{}
	err = c.JSONDecoder(res, &result)
	return &result, err
}

// Balance User Balance
// See https://api.reyden-x.com/docs#/User/get_balance_v1_user_balance__get
func Balance(c *rx.Client) (*m.Balance, error) {
	res, err := c.Get("/user/balance/")
	if err != nil {
		return nil, err
	}

	result := m.Balance{}
	err = c.JSONDecoder(res, &result)
	return &result, err
}
