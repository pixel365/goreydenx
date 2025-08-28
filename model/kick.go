package model

import (
	"errors"
	"net/url"

	rx "github.com/pixel365/goreydenx"
)

type KickParams struct {
	ChannelUrl string `json:"channel_url"`
	BaseOrderParams
}

func (o *KickParams) IsValid() (bool, error) {
	if len(o.ChannelUrl) == 0 {
		return false, errors.New("invalid channel url")
	}

	_url, err := url.ParseRequestURI(o.ChannelUrl)
	if err != nil {
		return false, errors.New("invalid channel url")
	}

	if _url.Scheme != "https" {
		return false, errors.New("invalid channel url scheme")
	}

	switch _url.Host {
	case "kick.com", "www.kick.com":
	default:
		return false, errors.New("invalid channel url host")
	}

	switch o.LaunchMode {
	case rx.LaunchModeAuto, rx.LaunchModeDelay, rx.LaunchModeManual:
	default:
		return false, errors.New("invalid launch mode")
	}

	return true, nil
}

func (o *KickParams) PlatformCode() rx.PlatformCode {
	return rx.Kick
}
