package model

import (
	"errors"

	rx "github.com/pixel365/goreydenx"
)

type TwitchParams struct {
	BaseOrderParams
	TwitchId int `json:"twitch_id"`
}

func (o *TwitchParams) IsValid() (bool, error) {
	if o.TwitchId < 1 || o.TwitchId > 10_000_000_000 {
		return false, errors.New("invalid twitch id")
	}

	switch o.LaunchMode {
	case rx.LaunchModeAuto, rx.LaunchModeDelay, rx.LaunchModeManual:
	default:
		return false, errors.New("invalid launch mode")
	}

	return true, nil
}

func (o *TwitchParams) PlatformCode() rx.PlatformCode {
	return rx.Twitch
}
