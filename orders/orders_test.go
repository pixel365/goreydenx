package orders

import (
	"bytes"
	rx "github.com/pixel365/goreydenx"
	m "github.com/pixel365/goreydenx/model"
	"io"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestOrders(t *testing.T) {
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:38:27.937Z\",\"result\":[{\"id\":0,\"created_at\":\"2023-09-05T05:38:27.937Z\",\"updated_at\":\"2023-09-05T05:38:27.937Z\",\"uuid\":\"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\"status\":\"string\",\"ordered_view_qty\":0,\"price_per_view\":0,\"is_autostart\":true,\"online_users_limit\":0,\"platform\":\"string\",\"content_type\":\"string\",\"parameters\":{},\"extras\":{},\"statistics\":{}}],\"cursor\":\"string\"}"
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

	res, err := Orders(client, "")
	if err != nil {
		t.Error(err)
	}

	if !res.HasNext() {
		t.Error("invalid cursor")
	}
}

func TestDetails(t *testing.T) {
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:52:41.978Z\",\"result\":{\"id\":0,\"created_at\":\"2023-09-05T05:52:41.978Z\",\"updated_at\":\"2023-09-05T05:52:41.978Z\",\"uuid\":\"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\"status\":\"string\",\"ordered_view_qty\":0,\"price_per_view\":0,\"is_autostart\":true,\"online_users_limit\":0,\"platform\":\"string\",\"content_type\":\"string\",\"parameters\":{},\"extras\":{},\"statistics\":{}}}"
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

	_, err := Details(client, 123)
	if err != nil {
		t.Error(err)
	}
}

func TestPayments(t *testing.T) {
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:54:16.881Z\",\"result\":[{\"id\":0,\"created_at\":\"2023-09-05T05:54:16.881Z\",\"updated_at\":\"2023-09-05T05:54:16.881Z\",\"payed_at\":\"2023-09-05T05:54:16.881Z\",\"order_id\":0,\"amount\":0,\"external_id\":\"string\",\"uuid\":\"3fa85f64-5717-4562-b3fc-2c963f66afa6\",\"receipt\":\"string\"}],\"cursor\":\"string\"}"
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

	_, err := Payments(client, 123, "")
	if err != nil {
		t.Error(err)
	}
}

func TestOnlineStats(t *testing.T) {
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:55:32.729Z\",\"result\":[{\"created_at\":\"2023-09-05T05:55:32.729Z\",\"in_settings\":0,\"in_fact\":0}]}"
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

	_, err := OnlineStats(client, 123)
	if err != nil {
		t.Error(err)
	}
}

func TestClicksStats(t *testing.T) {
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:56:39.812Z\",\"result\":[{\"date\":\"string\",\"quantity\":0}]}"
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

	_, err := ClicksStats(client, 123)
	if err != nil {
		t.Error(err)
	}
}

func TestViewsStats(t *testing.T) {
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:56:39.812Z\",\"result\":[{\"date\":\"string\",\"quantity\":0}]}"
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

	_, err := ViewsStats(client, 123)
	if err != nil {
		t.Error(err)
	}
}

func TestSitesStats(t *testing.T) {
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:57:22.392Z\",\"result\":[{\"domain\":\"\",\"views\":0,\"clicks\":0,\"ctr\":0}]}"
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

	_, err := SitesStats(client, 123)
	if err != nil {
		t.Error(err)
	}
}
func TestMultipleViewsStats(t *testing.T) {
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:58:29.239Z\",\"result\":[{\"id\":0,\"quantity\":0}]}"
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

	_, err := MultipleViewsStats(client, m.Identifiers{Identifiers: make([]uint32, 5)})
	if err != nil {
		t.Error(err)
	}
}

func TestMultipleClicksStats(t *testing.T) {
	data := "{\"request_id\":\"string\",\"cached\":false,\"cache_expires_at\":\"2023-09-05T05:58:29.239Z\",\"result\":[{\"id\":0,\"quantity\":0}]}"
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

	_, err := MultipleClicksStats(client, m.Identifiers{Identifiers: make([]uint32, 5)})
	if err != nil {
		t.Error(err)
	}
}

func TestCreateStream(t *testing.T) {
	data := "{\"request_id\":\"string\",\"order_id\":0,\"action\":\"string\",\"value\":0,\"task\":{\"id\":\"string\",\"url\":\"string\",\"expires_at\":\"2023-09-05T06:02:06.925Z\"}}"
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

	_, err := CreateStream[*m.TwitchParams](client, &m.TwitchParams{
		BaseOrderParams: m.BaseOrderParams{
			LaunchMode: rx.LaunchModeAuto,
			SmoothGain: m.SmoothGain{
				Enabled: false,
				Minutes: 0,
			},
			PriceId:         1,
			NumberOfViews:   10_000,
			NumberOfViewers: 100,
			DelayTime:       0,
		},
		TwitchId: 12345,
	})
	if err != nil {
		t.Error(err)
	}

	_, err = CreateStream[*m.YouTubeParams](client, &m.YouTubeParams{
		BaseOrderParams: m.BaseOrderParams{
			LaunchMode: rx.LaunchModeAuto,
			SmoothGain: m.SmoothGain{
				Enabled: false,
				Minutes: 0,
			},
			PriceId:         1,
			NumberOfViews:   10_000,
			NumberOfViewers: 100,
			DelayTime:       0,
		},
		ChannelUrl: "https://www.youtube.com/channel/UCtI0Hodo5o5dUb67FeUjDeA",
	})
	if err != nil {
		t.Error(err)
	}
}
