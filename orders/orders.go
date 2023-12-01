package orders

import (
	"bytes"
	"fmt"
	rx "github.com/pixel365/goreydenx"
	h "github.com/pixel365/goreydenx/helpers"
	m "github.com/pixel365/goreydenx/model"
)

// All orders
// See https://api.reyden-x.com/docs#/default/orders_v1_orders__get
func Orders(c *rx.Client, cursor string) (*m.Result[[]m.Order], error) {
	path := "/orders/"
	if len(cursor) > 0 {
		path = fmt.Sprintf("/orders/?cursor=%s", cursor)
	}
	return h.Get[[]m.Order](c, path)
}

// Order details
// See https://api.reyden-x.com/docs#/default/order_details_v1_orders__order_id___get
func Details(c *rx.Client, orderId uint32) (*m.Result[m.Order], error) {
	return h.Get[m.Order](c, fmt.Sprintf("/orders/%d/", orderId))
}

// Order payments
// See https://api.reyden-x.com/docs#/default/order_payments_v1_orders__order_id__payments__get
func Payments(c *rx.Client, orderId uint32, cursor string) (*m.Result[[]m.Payment], error) {
	path := fmt.Sprintf("/orders/%d/payments/", orderId)
	if len(cursor) > 0 {
		path = fmt.Sprintf("/orders/%d/payments/?cursor=%s", orderId, cursor)
	}
	return h.Get[[]m.Payment](c, path)
}

// Detailed information about users online
// See https://api.reyden-x.com/docs#/default/order_stats_online_v1_orders__order_id__statistics_online__get
func OnlineStats(c *rx.Client, orderId uint32) (*m.Result[[]m.OnlineStats], error) {
	return h.Get[[]m.OnlineStats](c, fmt.Sprintf("/orders/%d/statistics/online/", orderId))
}

// Detailed information about clicks
// See https://api.reyden-x.com/docs#/default/order_stats_clicks_v1_orders__order_id__statistics_clicks__get
func ClicksStats(c *rx.Client, orderId uint32) (*m.Result[[]m.DateQuantity], error) {
	return h.Get[[]m.DateQuantity](c, fmt.Sprintf("/orders/%d/statistics/clicks/", orderId))
}

// Detailed information about views
// See https://api.reyden-x.com/docs#/default/order_stats_views_v1_orders__order_id__statistics_views__get
func ViewsStats(c *rx.Client, orderId uint32) (*m.Result[[]m.DateQuantity], error) {
	return h.Get[[]m.DateQuantity](c, fmt.Sprintf("/orders/%d/statistics/views/", orderId))
}

// Detailed information about sites
// See https://api.reyden-x.com/docs#/default/order_stats_sites_v1_orders__order_id__statistics_sites__get
func SitesStats(c *rx.Client, orderId uint32) (*m.Result[[]m.SiteStats], error) {
	return h.Get[[]m.SiteStats](c, fmt.Sprintf("/orders/%d/statistics/views/", orderId))
}

// View statistics for multiple orders
// See https://api.reyden-x.com/docs#/default/multiple_views_v1_orders_multiple_views__post
func MultipleViewsStats(c *rx.Client, identifiers m.Identifiers) (*m.Result[[]m.IdQuantity], error) {
	payload, err := c.JSONEncoder(identifiers)
	if err != nil {
		return nil, err
	}
	return h.Post[[]m.IdQuantity](c, "/orders/multiple/views/", bytes.NewBuffer(payload))
}

// Click-through statistics for multiple orders
// See https://api.reyden-x.com/docs#/default/multiple_clicks_v1_orders_multiple_clicks__post
func MultipleClicksStats(c *rx.Client, identifiers m.Identifiers) (*m.Result[[]m.IdQuantity], error) {
	payload, err := c.JSONEncoder(identifiers)
	if err != nil {
		return nil, err
	}
	return h.Post[[]m.IdQuantity](c, "/orders/multiple/clicks/", bytes.NewBuffer(payload))
}

// Create new order for Twitch or Youtube stream
// See https://api.reyden-x.com/docs#/default/twitch_stream_v1_orders_create_twitch_stream__post
func CreateStream[T m.OrderParams](c *rx.Client, params T) (*m.ActionResult, error) {
	valid, err := params.IsValid()
	if !valid {
		return nil, err
	}

	payload, err := c.JSONEncoder(params)
	if err != nil {
		return nil, err
	}

	res, err := c.Post(fmt.Sprintf("/orders/create/%s/stream/", params.PlatformCode()), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	result := m.ActionResult{}
	err = c.JSONDecoder(res, &result)
	return &result, err
}
