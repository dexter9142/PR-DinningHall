package main

import (
	"os"
	"time"
)

var kitchenServerHost = "http://localhost"

const kitchenPort = ":8000"
const restaurantPort = ":7500"

const tableNumber = 6
const waiterQuantity = 3
const timeVar = 100 * time.Millisecond

var restaurant Restaurant

func main() {
	args := os.Args
	if len(args) > 1 {
		kitchenServerHost = args[1]
	}

	restaurant.start()
}
