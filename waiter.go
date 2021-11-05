package main

import (
	"strconv"
	"time"
)

//Waiter tunables
const getOrderTimeRequired = 3 * timeVar
const deliveryTimeRequired = 2 * timeVar

var waiterStatus = [...]string{"Waiting.", "Sending order id:", "Delivering delivery id:", "Waiting for orderList to clear."}

type Waiter struct {
	id         int
	atWork     int
	statusId   int
	modifierId int
}

func NewWaiter(id int, atWork int, statusId int, modifierId int) *Waiter {
	ret := new(Waiter)
	ret.id = id
	ret.atWork = atWork
	ret.statusId = statusId
	ret.modifierId = modifierId
	return ret
}

func (w *Waiter) startWorking() {
	w.atWork = 1
	for w.atWork == 1 {
		didATask := false

		order := restaurant.tableList.serveTable(w)

		for success := false; order != nil && !success; {
			success = restaurant.sendOrder(order)
			if success {
				didATask = true
				w.modifierId = order.Id
				w.statusId = 1
				time.Sleep(getOrderTimeRequired) //to remake, not ideal
			} else {
				w.statusId = 3
				time.Sleep(timeVar) //waits for orders to finish
			}
		}

		select {
		case delivery := <-restaurant.deliveryChan:   // // Receive from deliveryChan, and
														// assign value to delivery.
			didATask = true
			w.statusId = 2//Delivery to the table
			w.modifierId = delivery.OrderId
			time.Sleep(deliveryTimeRequired)
			now := genTime()
			restaurant.tableList.deliver(delivery, now)
		default:
			break
		}

		if !didATask {
			//Wait one unit because there are no tasks
			w.statusId = 0
			time.Sleep(timeVar)
		}

	}
}
func (w *Waiter) getStatus() string {
	status := "Waiter id:" + strconv.Itoa(w.id) + " Status:" + waiterStatus[w.statusId]
	if w.statusId == 1 || w.statusId == 2 {
		return status + strconv.Itoa(w.modifierId)
	}
	return status
}
