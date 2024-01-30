package traffic

import (
	rx "github.com/pixel365/goreydenx"
	h "github.com/pixel365/goreydenx/helpers"
	m "github.com/pixel365/goreydenx/model"
)

// Countries Traffic statistics by country
// See https://api.reyden-x.com/docs#/Traffic/Traffic_statistics_by_country_v1_traffic_countries__get
func Countries(c *rx.Client) (*m.Result[[]m.Traffic], error) {
	return h.Get[[]m.Traffic](c, "/traffic/countries/")
}

// Languages Traffic statistics by language
// See https://api.reyden-x.com/docs#/Traffic/Traffic_statistics_by_language_v1_traffic_languages__get
func Languages(c *rx.Client) (*m.Result[[]m.Traffic], error) {
	return h.Get[[]m.Traffic](c, "/traffic/languages/")
}

// Devices Traffic statistics by device type
// See https://api.reyden-x.com/docs#/Traffic/Traffic_statistics_by_device_type_v1_traffic_devices__get
func Devices(c *rx.Client) (*m.Result[[]m.Traffic], error) {
	return h.Get[[]m.Traffic](c, "/traffic/devices/")
}
