package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

//get input function

func getInput(prompt string, re *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := re.ReadString('\n')

	return strings.TrimSpace(input), err
}

//bill creating function

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	CustomerName, _ := getInput("Enter Customer Name: ", reader)

	//taking input for customer order

	//var totalNumber int16
	//var table int8

	//fmt.Println("Enter total number of guest please: \n")
	//fmt.Scanf("%d", &totalNumber)

	//fmt.Println("Please Enter the table number: \n")
	//fmt.Scanf("%d", &table)

	//order := customerOrder(CustomerName, totalNumber, table)

	billName := newBill(CustomerName)

	fmt.Println("Bill has been created -", billName.Name)

	return billName
}

// writhing prompt option function
func promptOption(bill Bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option ( A - order place, B - Add Food Items, C - Add Tips, D - Save Bill):", reader)

	switch option {
	case "A":
		name, _ := getInput("Enter yout name please: ", reader)
		var tNumber int8
		var pNumber int16
		fmt.Println("Please enter total number of guest: \n")
		fmt.Scanf("%d", &pNumber)
		fmt.Println("Please enter table number: \n")
		fmt.Scanf("%d", &tNumber)

		customerOrder(name, pNumber, tNumber)

		promptOption(bill)

	case "B":
		name, _ := getInput("Enter Name please : ", reader)
		price, _ := getInput("Enter price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("price must be a number!!")
			promptOption(bill)
		}

		bill.addItems(name, p)
		promptOption(bill)

	case "C":
		tip, _ := getInput("Enter tip Amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOption(bill)
		}
		bill.updateTip(t)

		fmt.Println("Tip added - ", tip)
		promptOption(bill)

	case "D":
		bill.save()
		fmt.Println("you saved the file -", bill.Name)

	default:
		fmt.Println("that was not an option !")
		promptOption(bill)

	}
}
func main() {

	fmt.Println("Welcome to Amar Dokan")

	menu()

	mybill := createBill()

	promptOption(mybill)

}
