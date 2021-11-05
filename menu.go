package main

var menu = []Menu{pizza, salad, zeama, sswmlc, idwmm, waffles, aubergine, lasagna, burger, gyros}
var apparatus = map[string]int{"": 0, "oven": 1, "stove": 2}

type Menu struct {
	id               int
	name             string
	preparationTime  int
	complexity       int
	cookingApparatus string
}

var pizza = Menu{1, "pizza", 20, 2, "oven"}
var salad = Menu{2, "salad", 10, 1, ""}
var zeama = Menu{3, "zeama", 7, 1, "stove"}
var sswmlc = Menu{4, "Scallop Sashimi with Meyer Lemon Confit", 32, 3, ""}
var idwmm = Menu{5, "Island Duck with Mulberry Mustard", 35, 3, "oven"}
var waffles = Menu{6, "Waffles", 10, 1, "stove"}
var aubergine = Menu{7, "Aubergine", 20, 2, ""}
var lasagna = Menu{8, "Lasagna", 30, 2, "oven"}
var burger = Menu{9, "Burger", 15, 1, "oven"}
var gyros = Menu{10, "Gyros", 15, 1, ""}
