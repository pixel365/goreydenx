package prices

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

func TestPrices(t *testing.T) {
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:50:18.430Z\",\"result\":[{\"id\":0,\"name\":\"string\",\"format\":\"string\",\"price\":0,\"views\":{\"min\":0,\"max\":0,\"step\":0},\"online_viewers\":{\"min\":0,\"max\":0,\"step\":0},\"description\":\"string\"}]}"
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

	_, err := Twitch(client)
	if err != nil {
		t.Error(err)
	}

	_, err = YouTube(client)
	if err != nil {
		t.Error(err)
	}

	_, err = Trovo(client)
	if err != nil {
		t.Error(err)
	}

	_, err = GoodGame(client)
	if err != nil {
		t.Error(err)
	}

	_, err = VkPlay(client)
	if err != nil {
		t.Error(err)
	}
}
