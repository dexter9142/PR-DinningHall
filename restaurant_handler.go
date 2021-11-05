package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RestaurantHandler struct {
	// packetsReceived int32
	// postReceived    int32
}

func (d RestaurantHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		{
			lastDelivery := new(Delivery)
			var requestBody = make([]byte, r.ContentLength)
			r.Body.Read(requestBody)
			json.Unmarshal(requestBody, lastDelivery)
			restaurant.deliveryChan <- lastDelivery // Send lastDelivery to channel delivery.
			fmt.Fprint(w, "OK")
		}
	case http.MethodGet:
		{
			fmt.Fprintln(w, "<head><meta http-equiv=\"refresh\" content=\"2\" /></head>")
			if restaurant.connected {
				fmt.Fprintln(w, Div("Restaurant successfully connected"))
			} else {
				fmt.Fprintln(w, Div("Did not establish connection"))
				err := restaurant.restaurantWeb.connectionError
				if err != nil {
					fmt.Fprintln(w, Div("Connection error: "+err.Error()))
				}
			}
			fmt.Fprintln(w, Div(restaurant.getStatus()))
		}
	case http.MethodConnect:
		{
			restaurant.connectionSuccessful()
			fmt.Fprint(w, "OK")
		}
	default:
		{
			fmt.Fprintln(w, "UNSUPPORTED METHOD")
		}
	}
}
