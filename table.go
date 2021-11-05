package main

import (
	"strconv"
	"time"
)

//Table tunables
const waitForCustomersTime = timeVar * 5
const maxWaitForCustomersMultiplier = 5
const eatingTimePerItem = timeVar * 5

var tableStatus = [...]string{"Waiting for customers.", "Waiting for waiter.", "Waiting for delivery.", "Eating."}

type Table struct {
	id     int
	status int
	order  *Order
}

func NewTable(id int, status int, order *Order) *Table {
	ret := new(Table)
	ret.id = id
	ret.status = status
	ret.order = order
	return ret
}

func (t *Table) deliver(delivery *Delivery, now int64) {
	t.status = 3
	t.order = nil

	rating := 0
	maxWaitF := float64(delivery.MaxWait)
	timeWaitedF := float64(now - delivery.PickUpTime)
	if maxWaitF > timeWaitedF {
		rating += 1
	}
	if maxWaitF*1.1 > timeWaitedF {
		rating += 1
	}
	if maxWaitF*1.2 > timeWaitedF {
		rating += 1
	}
	if maxWaitF*1.3 > timeWaitedF {
		rating += 1
	}
	if maxWaitF*1.4 > timeWaitedF {
		rating += 1
	}
	restaurant.ratings.addValue(rating)
	go func() {
		time.Sleep(eatingTimePerItem * (time.Duration(len(delivery.Items) + 1)))
		t.waitCustomers()
	}()
}

func (t *Table) waitCustomers() {
	t.status = 0
	time.Sleep(waitForCustomersTime * time.Duration(maxWaitForCustomersMultiplier))
	t.order = makeOrder(t)
	t.status = 1
}

func (t *Table) serve(waiter *Waiter) *Order {
	t.status = 2
	t.order.WaiterId = waiter.id
	t.order.PickUpTime = genTime()
	return t.order
}

func (t *Table) getStatus() string {
	waitStatus := ""
	if t.order != nil && t.status == 2 {
		waitStatus = " Waiting for:" + strconv.Itoa(int(genTime()-t.order.PickUpTime)) + "sec" + " Max wait:" + strconv.Itoa(t.order.MaxWait)
	}
	return "Table id:" + strconv.Itoa(t.id) + " Status:" + tableStatus[t.status] + waitStatus
}
