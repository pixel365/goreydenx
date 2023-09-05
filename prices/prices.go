package prices

import (
	"fmt"
	rx "github.com/pixel365/goreydenx"
	m "github.com/pixel365/goreydenx/model"
)

func request(c *rx.Client, platform string) (*m.Result[[]m.Price], error) {
	res, err := c.Get(fmt.Sprintf("/prices/%s/", platform))
	if err != nil {
		return nil, err
	}

	result := m.Result[[]m.Price]{}
	err = c.JSONDecoder(res, &result)
	return &result, err
}

// All prices for Twitch
// See https://api.reyden-x.com/docs#/default/prices_v1_prices__platform_code___get
func Twitch(c *rx.Client) (*m.Result[[]m.Price], error) {
	return request(c, rx.Twitch)
}

// All prices for YouTube
// See https://api.reyden-x.com/docs#/default/prices_v1_prices__platform_code___get
func YouTube(c *rx.Client) (*m.Result[[]m.Price], error) {
	return request(c, rx.YouTube)
}

// All prices for Trovo
// See https://api.reyden-x.com/docs#/default/prices_v1_prices__platform_code___get
func Trovo(c *rx.Client) (*m.Result[[]m.Price], error) {
	return request(c, rx.Trovo)
}

// All prices for GoodGame
// See https://api.reyden-x.com/docs#/default/prices_v1_prices__platform_code___get
func GoodGame(c *rx.Client) (*m.Result[[]m.Price], error) {
	return request(c, rx.GoodGame)
}

// All prices for VkPlay
// See https://api.reyden-x.com/docs#/default/prices_v1_prices__platform_code___get
func VkPlay(c *rx.Client) (*m.Result[[]m.Price], error) {
	return request(c, rx.VkPlay)
}
