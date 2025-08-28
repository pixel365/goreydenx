package goreydenx

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-module/carbon"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	TokenType   string `json:"token_type,omitempty"`
}

func (t *Token) IsValid() bool {
	if t.AccessToken == "" || t.ExpiresIn == "" {
		return false
	}

	now := carbon.Now(carbon.UTC)
	return carbon.Parse(t.ExpiresIn, carbon.UTC).Compare(">=", now)
}

type RxClient interface {
	Get(string) ([]byte, error)
	Post(string, io.Reader) ([]byte, error)
	Patch(string, io.Reader) ([]byte, error)
}

type Client struct {
	*http.Client
	Token    *Token
	BaseUrl  string
	userName string
	password string
}

func (c *Client) isAuthenticated() bool {
	if c.Token == nil {
		return false
	}
	return c.Token.IsValid()
}

func (c *Client) Auth() *Client {
	if c.isAuthenticated() {
		return c
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if c.Token == nil {
		form := url.Values{}
		form.Add("username", c.userName)
		form.Add("password", c.password)
		req, _ := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			c.BaseUrl+"/token/",
			strings.NewReader(form.Encode()),
		)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := c.Do(req)
		defer func() {
			_ = resp.Body.Close()
		}()

		if err == nil && resp.StatusCode == 200 {
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				t := Token{}
				err := json.Unmarshal(body, &t)
				if err == nil {
					c.Token = &t
				}
			}
		}

		return c
	}

	if !c.Token.IsValid() {
		token, _ := c.RefreshToken()
		c.Token = token
	}

	return c
}

func (c *Client) RefreshToken() (*Token, error) {
	resp, err := request(c, http.MethodPatch, "/token/refresh/", nil)
	if err == nil {
		token := Token{}
		if decodeError := json.Unmarshal(resp, &token); decodeError != nil {
			return nil, decodeError
		}
		return &token, nil
	}

	return nil, err
}

func request(client *Client, method string, path string, payload io.Reader) ([]byte, error) {
	if !client.isAuthenticated() {
		return nil, errors.New("unauthorized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, method, client.BaseUrl+path, payload)
	req.Header.Set("Accept", "application/json")
	if client.Token != nil && client.Token.IsValid() && path != "/token/" {
		req.Header.Set("Authorization", "Bearer "+client.Token.AccessToken)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	switch resp.StatusCode {
	case http.StatusBadRequest:
		return nil, errors.New("bad request")
	case http.StatusTooManyRequests:
		return nil, errors.New("too many requests")
	case http.StatusNotFound:
		return nil, errors.New("not found")
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized")
	case http.StatusOK:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, err
	default:
		return nil, errors.New("unknown error")
	}
}

func (c *Client) Get(path string) ([]byte, error) {
	return request(c, http.MethodGet, path, nil)
}

func (c *Client) Post(path string, payload io.Reader) ([]byte, error) {
	return request(c, http.MethodPost, path, payload)
}

func (c *Client) Patch(path string, payload io.Reader) ([]byte, error) {
	return request(c, http.MethodPatch, path, payload)
}

// NewClient Make Client instance
func NewClient(userName string, password string) *Client {
	return &Client{
		Client: &http.Client{
			Timeout: time.Second * 5,
		},
		BaseUrl:  BaseUrl,
		userName: userName,
		password: password,
	}
}

// NewClientWithToken Make Client instance
func NewClientWithToken(token *Token) *Client {
	return &Client{
		Client: &http.Client{
			Timeout: time.Second * 5,
		},
		BaseUrl: BaseUrl,
		Token:   token,
	}
}
