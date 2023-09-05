package action

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

func TestActions(t *testing.T) {
	data := "{\"request_id\":\"string\",\"order_id\":0,\"action\":\"string\",\"value\":0,\"task\":{\"id\":\"string\",\"url\":\"string\",\"expires_at\":\"2023-09-05T05:45:26.018Z\"}}"
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

	_, err := Run(client, 123)
	if err != nil {
		t.Error(err)
	}

	_, err = Stop(client, 123)
	if err != nil {
		t.Error(err)
	}

	_, err = Cancel(client, 123)
	if err != nil {
		t.Error(err)
	}

	_, err = ChangeOnline(client, 123, 100)
	if err != nil {
		t.Error(err)
	}

	_, err = ChangeIncreaseValue(client, 123, 100)
	if err != nil {
		t.Error(err)
	}

	_, err = IncreaseOn(client, 123, 100)
	if err != nil {
		t.Error(err)
	}

	_, err = IncreaseOff(client, 123)
	if err != nil {
		t.Error(err)
	}

	_, err = AddViews(client, 123, 100)
	if err != nil {
		t.Error(err)
	}
}

func TestTaskStatus(t *testing.T) {
	data := "{\"status\":\"pending\"}"
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

	_, err := TaskStatus(client, 123, "TaskID")
	if err != nil {
		t.Error(err)
	}
}
