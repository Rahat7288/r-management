package main

import (
	"fmt"
)

//function for the menu

func menu() {
	food := map[string]float64{
		"Bread":   2.55,
		"Cake":    5.00,
		"Pie":     3.25,
		"Pudding": 10.25,
		"Bugger":  50.00,
		"Pizza":   150.25,
		"Brownee": 40.35,
	}

	beverage := map[string]float64{
		"Coke":  20.00,
		"Pepsi": 20.00,
		"Juice": 30.23,
		"Water": 10.12,
	}

	fmt.Println("======MENU========")
	fmt.Println("\n")

	for solid, price := range food {

		fmt.Println(solid, ":", price, "$")

	}

	fmt.Println("\n")

	fmt.Println("========Beverage========")
	fmt.Println("\n")

	for bev, pri := range beverage {

		fmt.Println(bev, ":", pri, "$")
	}

}

func main() {

	fmt.Println("Welcome to Amar Dokan")

	menu()

	//myBill := newBill("Rahat")

	//fmt.Println("\n")

	//cus := customerOrder("Rahat", 6, 20)

	//fmt.Println(myBill)
	//fmt.Println(cus)

}
