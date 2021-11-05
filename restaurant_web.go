package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

type RestaurantWeb struct {
	restaurantServer http.Server
	restaurantHan    RestaurantHandler
	restaurantClient http.Client
	connectionError  error
}

func (dhw *RestaurantWeb) start() {
	dhw.restaurantServer.Addr = restaurantPort
	dhw.restaurantServer.Handler = &dhw.restaurantHan

	fmt.Println(time.Now())
	if err := dhw.restaurantServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (dhw *RestaurantWeb) sendOrder(order *Order) bool {
	requestBody := order.getPayload()
	request, _ := http.NewRequest(http.MethodPost, kitchenServerHost+kitchenPort+"/order", bytes.NewBuffer(requestBody))
	response, err := dhw.restaurantClient.Do(request)

	if err != nil {
		fmt.Println("Could not send order to kitchen.")
		log.Fatal(err)
		return false
	}
	var responseBody = make([]byte, response.ContentLength)
	response.Body.Read(responseBody)
	if string(responseBody) != "OK" {
		return false
	}

	return true
}

func (dhw *RestaurantWeb) establishConnection() bool {
	if restaurant.connected == true {
		return false
	}
	request, _ := http.NewRequest(http.MethodConnect, kitchenServerHost+kitchenPort+"/", bytes.NewBuffer([]byte{}))
	response, err := dhw.restaurantClient.Do(request)
	if err != nil {
		dhw.connectionError = err
		return false
	}
	var responseBody = make([]byte, response.ContentLength)
	response.Body.Read(responseBody)
	if string(responseBody) != "OK" {
		return false
	}

	return true
}
