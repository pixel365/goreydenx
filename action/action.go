package action

import (
	"fmt"
	rx "github.com/pixel365/goreydenx"
	m "github.com/pixel365/goreydenx/model"
)

func request(c *rx.Client, path string) (*m.ActionResult, error) {
	res, err := c.Patch(path, nil)
	if err != nil {
		return nil, err
	}

	result := m.ActionResult{}
	err = c.JSONDecoder(res, &result)
	return &result, err
}

// TaskStatus Check the task status
// See https://api.reyden-x.com/docs#/Orders/order_get_task_status_v1_orders__order_id__task__task_id__status__get
func TaskStatus(c *rx.Client, orderId uint32, taskId string) (string, error) {
	res, err := c.Get(fmt.Sprintf("/orders/%d/task/%s/status/", orderId, taskId))
	if err != nil {
		return "", err
	}

	result := struct {
		Status string `json:"status"`
	}{}
	err = c.JSONDecoder(res, &result)
	return result.Status, err
}

// Run the order
// See https://api.reyden-x.com/docs#/Orders/order_run_v1_orders__order_id__action_run__patch
func Run(c *rx.Client, orderId uint32) (*m.ActionResult, error) {
	return request(c, fmt.Sprintf("/orders/%d/action/run/", orderId))
}

// Stop the order
// See https://api.reyden-x.com/docs#/Orders/order_stop_v1_orders__order_id__action_stop__patch
func Stop(c *rx.Client, orderId uint32) (*m.ActionResult, error) {
	return request(c, fmt.Sprintf("/orders/%d/action/stop/", orderId))
}

// Cancel the order
// See https://api.reyden-x.com/docs#/Orders/order_cancel_v1_orders__order_id__action_cancel__patch
func Cancel(c *rx.Client, orderId uint32) (*m.ActionResult, error) {
	return request(c, fmt.Sprintf("/orders/%d/action/cancel/", orderId))
}

// ChangeOnline Change the number of viewers
// See https://api.reyden-x.com/docs#/Orders/order_change_online_v1_orders__order_id__action_change_online__value___patch
func ChangeOnline(c *rx.Client, orderId uint32, value uint32) (*m.ActionResult, error) {
	return request(c, fmt.Sprintf("/orders/%d/action/change/online/%d/", orderId, value))
}

// ChangeIncreaseValue Change the time of the smooth set of viewers
// See https://api.reyden-x.com/docs#/Orders/change_increase_value_v1_orders__order_id__action_increase_change__value___patch
func ChangeIncreaseValue(c *rx.Client, orderId uint32, value uint32) (*m.ActionResult, error) {
	return request(c, fmt.Sprintf("/orders/%d/action/increase/change/%d/", orderId, value))
}

// IncreaseOn Enable smooth increase of viewers
// See https://api.reyden-x.com/docs#/Orders/increase_on_v1_orders__order_id__action_increase_on__value___patch
func IncreaseOn(c *rx.Client, orderId uint32, value uint32) (*m.ActionResult, error) {
	return request(c, fmt.Sprintf("/orders/%d/action/increase/on/%d/", orderId, value))
}

// IncreaseOff Disable smooth increase of viewers
// See https://api.reyden-x.com/docs#/Orders/increase_off_v1_orders__order_id__action_increase_off__patch
func IncreaseOff(c *rx.Client, orderId uint32) (*m.ActionResult, error) {
	return request(c, fmt.Sprintf("/orders/%d/action/increase/off/", orderId))
}

// AddViews Add views to order
// See https://api.reyden-x.com/docs#/Orders/add_views_v1_orders__order_id__action_add_views__value___patch
func AddViews(c *rx.Client, orderId uint32, value uint32) (*m.ActionResult, error) {
	return request(c, fmt.Sprintf("/orders/%d/action/add/views/%d/", orderId, value))
}
