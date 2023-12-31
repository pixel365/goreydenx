package goreydenx

import (
	"errors"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type JSONMarshal func(v interface{}) ([]byte, error)
type JSONUnmarshal func(data []byte, v interface{}) error

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

func (t *Token) IsValid() bool {
	//TODO implements
	return t.AccessToken != ""
}

type Client struct {
	*http.Client
	JSONEncoder JSONMarshal
	JSONDecoder JSONUnmarshal
	Token       *Token
	BaseUrl     string
	userName    string
	password    string
}

func (c *Client) isAuthenticated() bool {
	return c.Token != nil && c.Token.IsValid()
}

func (c *Client) Auth() *Client {
	if c.isAuthenticated() {
		return c
	}

	if c.Token == nil {
		form := url.Values{}
		form.Add("username", c.userName)
		form.Add("password", c.password)
		req, _ := http.NewRequest(http.MethodPost, c.BaseUrl+"/token/", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, err := c.Do(req)
		if err == nil && resp.StatusCode == 200 {
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				t := Token{}
				err := c.JSONDecoder(body, &t)
				if err == nil {
					c.Token = &t
				}
			}
		}
	} else {
		if !c.Token.IsValid() {
			c.refreshToken()
		}
	}

	return c
}

func (c *Client) refreshToken() {
	resp, err := request(c, http.MethodPatch, "/token/refresh/", nil)
	if err != nil {
		t := Token{}
		err := c.JSONDecoder(resp, &t)
		if err == nil {
			c.Token = &t
		}
	} else {
		c.Token = nil
	}
}

func request(client *Client, method string, path string, payload io.Reader) ([]byte, error) {
	if !client.isAuthenticated() {
		return nil, errors.New("unauthorized")
	}

	req, _ := http.NewRequest(method, client.BaseUrl+path, payload)
	req.Header.Set("Accept", "application/json")
	if client.Token != nil && client.Token.IsValid() && path != "/token/" {
		req.Header.Set("Authorization", "Bearer "+client.Token.AccessToken)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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

// Make new Client instance
func NewClient(userName string, password string) *Client {
	return &Client{
		Client: &http.Client{
			Timeout: time.Second * 5,
		},
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		BaseUrl:     BaseUrl,
		userName:    userName,
		password:    password,
	}
}
