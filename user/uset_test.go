package user

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	rx "github.com/pixel365/goreydenx"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestAccount(t *testing.T) {
	token := &rx.Token{AccessToken: "token", ExpiresIn: "2030-01-01T00:00:00Z"}

	//nolint: lll
	data := "{\"id\":0,\"username\":\"string\",\"date_joined\":\"\",\"email\":\"string\",\"is_active\":true,\"is_blocked\":true,\"has_image\":true,\"image_extension\":\"string\",\"image_url\":\"string\",\"currency_id\":0,\"discount_value\":0,\"is_reseller\":false,\"twitch_id\":0,\"twitch_login\":\"\"}"
	client := rx.NewClientWithToken(token)
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(data)),
			Header:     make(http.Header),
		}
	})

	_, err := Account(client)
	require.NoError(t, err)
}

func TestBalance(t *testing.T) {
	token := &rx.Token{AccessToken: "token", ExpiresIn: "2030-01-01T00:00:00Z"}

	data := "{\"id\":0,\"amount\":0,\"currency_id\":0,\"user_id\":0,\"formatted_amount\":0,\"currency\":\"string\"}"
	client := rx.NewClientWithToken(token)
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(data)),
			Header:     make(http.Header),
		}
	})

	_, err := Balance(client)
	require.NoError(t, err)
}
