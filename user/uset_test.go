package user

import (
	"bytes"
	rx "github.com/pixel365/goreydenx"
	"io"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestAccount(t *testing.T) {
	data := "{\"id\":0,\"username\":\"string\",\"date_joined\":\"\",\"email\":\"string\",\"is_active\":true,\"is_blocked\":true,\"has_image\":true,\"image_extension\":\"string\",\"image_url\":\"string\",\"currency_id\":0,\"discount_value\":0,\"is_reseller\":false,\"twitch_id\":0,\"twitch_login\":\"\"}"
	client := rx.NewClient("", "")
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(data)),
			Header:     make(http.Header),
		}
	})
	client.Token = &rx.Token{
		AccessToken: "fake token",
		ExpiresIn:   "fake date",
	}

	_, err := Account(client)
	if err != nil {
		t.Error(err)
	}
}

func TestBalance(t *testing.T) {
	data := "{\"id\":0,\"amount\":0,\"currency_id\":0,\"user_id\":0,\"formatted_amount\":0,\"currency\":\"string\"}"
	client := rx.NewClient("", "")
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(data)),
			Header:     make(http.Header),
		}
	})
	client.Token = &rx.Token{
		AccessToken: "fake token",
		ExpiresIn:   "fake date",
	}

	_, err := Balance(client)
	if err != nil {
		t.Error(err)
	}
}
