package model

import (
	"errors"

	rx "github.com/pixel365/goreydenx"
)

type Result[T any] struct {
	Result         T      `json:"result"`
	RequestId      string `json:"request_id"`
	CacheExpiresAt string `json:"cache_expires_at"`
	Cursor         string `json:"cursor"`
	Cached         bool   `json:"cached"`
}

func (o *Result[T]) HasNext() bool {
	return o.Cursor != ""
}

type Task struct {
	Id        string `json:"id"`
	Url       string `json:"url"`
	ExpiresAt string `json:"expires_at"`
}

type ActionResult struct {
	Task      Task   `json:"task"`
	RequestId string `json:"request_id"`
	Action    string `json:"action"`
	OrderId   uint32 `json:"order_id"`
	Value     uint32 `json:"value"`
}

type User struct {
	Username      string `json:"username"`
	DateJoined    string `json:"date_joined"`
	Email         string `json:"email"`
	ImageUrl      string `json:"image_url"`
	TwitchLogin   string `json:"twitch_login"`
	Id            uint32 `json:"id"`
	CurrencyId    uint32 `json:"currency_id"`
	DiscountValue uint32 `json:"discount_value"`
	TwitchId      uint32 `json:"twitch_id"`
	IsActive      bool   `json:"is_active"`
	IsBlocked     bool   `json:"is_blocked"`
	IsReseller    bool   `json:"is_reseller"`
}

type Balance struct {
	Currency        string  `json:"currency"`
	FormattedAmount float64 `json:"formatted_amount"`
	Amount          float64 `json:"amount"`
	Id              uint32  `json:"id"`
	CurrencyId      uint32  `json:"currency_id"`
}

type MinMaxStep struct {
	Min  uint32 `json:"min"`
	Max  uint32 `json:"max"`
	Step uint32 `json:"step"`
}

type Price struct {
	Name          string     `json:"name"`
	Format        string     `json:"format"`
	Description   string     `json:"description"`
	Price         float64    `json:"price"`
	Views         MinMaxStep `json:"views"`
	OnlineViewers MinMaxStep `json:"online_viewers"`
	Id            uint32     `json:"id"`
	CategoryId    uint32     `json:"category_id"`
}

type PriceCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Id          uint32 `json:"id"`
	IsActive    bool   `json:"is_active"`
}

type Statistics struct {
	ActiveTimeInSeconds uint32  `json:"active_time_in_seconds,omitempty"`
	Views               uint32  `json:"views,omitempty"`
	Clicks              uint32  `json:"clicks,omitempty"`
	Ctr                 float32 `json:"ctr,omitempty"`
}

type Parameters struct {
	LaunchMode           string `json:"launch_mode,omitempty"`
	WorkMode             string `json:"work_mode,omitempty"`
	Delay                bool   `json:"delay,omitempty"`
	DelayTime            uint32 `json:"delay_time,omitempty"`
	EvenDistribution     bool   `json:"even_distribution,omitempty"`
	EvenDistributionTime uint32 `json:"even_distribution_time,omitempty"`
}

type Order struct {
	Platform                    string     `json:"platform"`
	Uuid                        string     `json:"uuid"`
	Status                      string     `json:"status"`
	UpdatedAt                   string     `json:"updated_at"`
	CreatedAt                   string     `json:"created_at"`
	ContentType                 string     `json:"content_type"`
	ContentClassificationLabels []string   `json:"content_classification_labels,omitempty"`
	Parameters                  Parameters `json:"parameters,omitempty"`
	PricePerView                float64    `json:"price_per_view"`
	Statistics                  Statistics `json:"statistics,omitempty"`
	OrderedViewQty              uint32     `json:"ordered_view_qty"`
	OnlineUsersLimit            uint32     `json:"online_users_limit"`
	PriceId                     uint32     `json:"tariff_id"`
	Id                          uint32     `json:"id"`
	IsAutostart                 bool       `json:"is_autostart"`
}

type Payment struct {
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	PayedAt    string  `json:"payed_at"`
	ExternalId string  `json:"external_id"`
	Receipt    string  `json:"receipt"`
	Amount     float64 `json:"amount"`
	Id         uint32  `json:"id"`
	OrderId    uint32  `json:"order_id"`
}

type OnlineStats struct {
	CreatedAt  string `json:"created_at"`
	InSettings uint32 `json:"in_settings"`
	InFact     uint32 `json:"in_fact"`
}

type DateQuantity struct {
	Date     string `json:"date"`
	Quantity uint32 `json:"quantity"`
}

type IdQuantity struct {
	Id       uint32 `json:"id"`
	Quantity uint32 `json:"quantity"`
}

type SiteStats struct {
	Domain string  `json:"domain"`
	Views  uint32  `json:"views"`
	Clicks uint32  `json:"clicks"`
	Ctr    float32 `json:"ctr"`
}

type Identifiers struct {
	Identifiers []uint32 `json:"identifiers"`
}

type SmoothGain struct {
	Enabled bool   `json:"enabled"`
	Minutes uint32 `json:"minutes"`
}

type BaseOrderParams struct {
	LaunchMode      string     `json:"launch_mode"`
	SmoothGain      SmoothGain `json:"smooth_gain"`
	PriceId         uint32     `json:"price_id"`
	NumberOfViews   uint32     `json:"number_of_views"`
	NumberOfViewers uint32     `json:"number_of_viewers"`
	DelayTime       uint32     `json:"delay_time"`
}

type OrderParams interface {
	IsValid() (bool, error)
	PlatformCode() string
}

type TwitchParams struct {
	BaseOrderParams
	TwitchId uint32 `json:"twitch_id"`
}

func (o *TwitchParams) IsValid() (bool, error) {
	if o.TwitchId < 1 {
		return false, errors.New("twitch id must be greater than zero")
	}
	switch o.LaunchMode {
	case rx.LaunchModeAuto, rx.LaunchModeDelay, rx.LaunchModeManual:
		return true, nil
	default:
		return false, errors.New("invalid launch mode")
	}
}

func (o *TwitchParams) PlatformCode() string {
	return rx.Twitch
}

type YouTubeParams struct {
	ChannelUrl string `json:"channel_url"`
	BaseOrderParams
}

func (o *YouTubeParams) IsValid() (bool, error) {
	if len(o.ChannelUrl) == 0 {
		return false, errors.New("invalid channel url")
	}
	switch o.LaunchMode {
	case rx.LaunchModeAuto, rx.LaunchModeDelay, rx.LaunchModeManual:
		return true, nil
	default:
		return false, errors.New("invalid launch mode")
	}
}

func (o *YouTubeParams) PlatformCode() string {
	return rx.YouTube
}

type Traffic struct {
	Code     string `json:"code"`
	Quantity uint32 `json:"quantity"`
}
