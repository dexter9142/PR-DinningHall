package main

import (
	"fmt"
	"time"
)

type Restaurant struct {
	restaurantWeb RestaurantWeb
	waiterList    *WaiterList
	tableList     *TableList
	deliveryChan  chan *Delivery
	ratings       *Rating
	connected     bool
	startTime     time.Time
}

func (dh *Restaurant) start() {
	dh.ratings = NewRating()
	dh.waiterList = NewWaiterList()
	dh.tableList = NewTableList()
	go dh.tryConnectKitchen() //it requires a goroutine, because in case the Kitchen hasn't yet started, so that it waits for the connection to occour
	dh.restaurantWeb.start()
}

func (dh *Restaurant) connectionSuccessful() {
	dh.startTime = time.Now()
	if dh.connected {
		return
	}
	dh.connected = true
	dh.deliveryChan = make(chan *Delivery) //Buffered Channel 
	//This creates a buffered channel with a capacity of 1. Normally channels are synchronous;
	//both sides of the channel will wait until the other side is ready. A buffered channel is asynchronous; sending or receiving a message will not wait unless the channel is already full.
	dh.tableList.start()
	dh.waiterList.start()
}

func (dh *Restaurant) tryConnectKitchen() {
	dh.connected = false
	for !dh.connected {
		if dh.restaurantWeb.establishConnection() {
			dh.connectionSuccessful()
			break
		} else {
			time.Sleep(timeVar)
		}
	}
}

func (dh *Restaurant) sendOrder(order *Order) bool {
	return dh.restaurantWeb.sendOrder(order)
}

func (dh *Restaurant) getStatus() string {
	ret := "Total runtime:" + fmt.Sprintf("%v", time.Since(dh.startTime))
	ret += Div("Rating:" + fmt.Sprintf("%f", dh.ratings.getAverage()) + " Total reviews:" + fmt.Sprintf("%d", dh.ratings.getNumOfOrders()))
	ret += "Waiters:"
	for _, waiter := range dh.waiterList.waiterList {
		ret += Div(waiter.getStatus())
	}
	ret += "Tables:"
	for _, table := range dh.tableList.tableList {
		ret += Div(table.getStatus())
	}

	return ret
}
