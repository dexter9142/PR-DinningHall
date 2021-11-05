package main

import (
	"encoding/json"
	"log"
	"math/rand"
)

type FoodDelivery struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

type Delivery struct {
	OrderId        int            `json:"order_id"`
	TableId        int            `json:"table_id"`
	Items          []int          `json:"items"`
	Priority       int            `json:"priority"`
	MaxWait        int            `json:"max_wait"`
	PickUpTime     int64          `json:"pick_up_time"`
	CookingTime    int            `json:"cooking_time"`
	CookingDetails []FoodDelivery `json:"cooking_details"`
} 															//This is the main concurrency channel

type Rating struct {
	values   []int
	maxSize  int
	orderNum int
	average  float32
	full     bool
}

type Order struct {
	Id         int   `json:"order_id"`
	TableId    int   `json:"table_id"`
	WaiterId   int   `json:"waiter_id"`
	Items      []int `json:"items"`
	Priority   int   `json:"priority"`
	MaxWait    int   `json:"max_wait"`
	PickUpTime int64 `json:"pick_up_time"`
}

var orderIdCounter = 1

func (o *Order) getPayload() []byte {
	result, err := json.Marshal(*o)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return result
}

func getOrderId() int {
	orderIdCounter++
	return orderIdCounter - 1
}

func makeOrder(table *Table) *Order {

	itemNum := rand.Intn(5) + 1
	var items []int
	maxWait := -1
	for i := 0; i < itemNum; i++ {
		item := rand.Intn(len(menu)) // randomizer 
		items = append(items, item)
		itemWait := menu[item].preparationTime * 3
		if itemWait > maxWait {
			maxWait = itemWait
		}
	}
	ret := new(Order)

	ret.Id = getOrderId()
	ret.TableId = table.id
	ret.WaiterId = -1
	ret.Items = items
	ret.Priority = rand.Intn(3)
	ret.MaxWait = maxWait
	ret.PickUpTime = genTime()

	return ret
}

func NewRating() *Rating {
	maxSize := 100
	return &Rating{maxSize: maxSize, values: make([]int, maxSize), full: false, average: 0}
}

func (r *Rating) addValue(rating int) {
	if !r.full && r.orderNum >= r.maxSize {
		r.full = true
	}

	r.values[r.orderNum%r.maxSize] = rating
	r.orderNum++

	numberOfReviews := 0
	if r.full {
		numberOfReviews = r.maxSize
	} else {
		numberOfReviews = r.orderNum
	}

	sum := 0
	for i := 0; i < numberOfReviews; i++ {
		sum += r.values[i]
	}

	r.average = float32(sum) / float32(numberOfReviews)
}

func (r *Rating) getAverage() float32 {
	return r.average
}

func (r *Rating) getNumOfOrders() int {
	return r.orderNum
}
