package prices

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

func TestPrices(t *testing.T) {
	token := &rx.Token{AccessToken: "token", ExpiresIn: "2030-01-01T00:00:00Z"}

	//nolint: lll
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:50:18.430Z\",\"result\":[{\"id\":0,\"name\":\"string\",\"format\":\"string\",\"price\":0,\"views\":{\"min\":0,\"max\":0,\"step\":0},\"online_viewers\":{\"min\":0,\"max\":0,\"step\":0},\"description\":\"string\"}]}"
	client := rx.NewClientWithToken(token)
	client.Transport = RoundTripFunc(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(data)),
			Header:     make(http.Header),
		}
	})

	_, err := Twitch(client)
	require.NoError(t, err)

	_, err = YouTube(client)
	require.NoError(t, err)

	_, err = Trovo(client)
	require.NoError(t, err)

	_, err = GoodGame(client)
	require.NoError(t, err)

	_, err = VkPlay(client)
	require.NoError(t, err)

	_, err = Kick(client)
	require.NoError(t, err)
}
