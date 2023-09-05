# REYDEN-X

###### Reyden-X is an automated service for promoting live broadcasts on external sites with integrated system of viewers and views management.

- [Website](https://reyden-x.com/en)

- [API Documentation](https://api.reyden-x.com/docs)


### Quickstart

```go
package main

import (
	"fmt"
	rx "github.com/pixel365/goreydenx"
	m "github.com/pixel365/goreydenx/model"
	"github.com/pixel365/goreydenx/orders"
	"github.com/pixel365/goreydenx/prices"
	"github.com/pixel365/goreydenx/user"
)

func main() {
	// new client
	client := rx.NewClient("EMAIL", "PASSWORD").Auth()
	
	// user account details
	res, err := user.Account(client)
	if err != nil {
		fmt.Println(res)
	}

	// order details
	res, err = orders.Details(client, 12345)
	if err != nil {
		fmt.Println(res)
	}

	// prices for Twitch
	res, err = prices.Twitch(client)
	if err != nil {
		fmt.Println(res)
	}

	// make new order
	res, err = orders.CreateStream[*m.TwitchParams](client, &m.TwitchParams{
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
		fmt.Println(res)
	}
}
```
