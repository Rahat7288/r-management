package main

import (
	"fmt"
	"os"
)

//function for the customer information and get the order information

func customerOrder(Name string, person int16, tNumber int8) Customer {

	custo := Customer{
		customerName:  Name,
		totalCustomer: person,
		tableNumber:   tNumber,
	}
	return custo

}

// creating a neew bill function which reaturn a bill type structure

func newBill(name string) Bill {

	bill := Bill{
		Name:      name,
		foodItems: map[string]float64{},
		tips:      0,
		vat:       0,
	}

	return bill

}

//bill formating function

func (bill *Bill) format() string {

	formateString := "Bill Breakdown:  \n"

	var total float64 = 0

	//fetching the foodItems

	for k, v := range bill.foodItems {
		formateString += fmt.Sprintf("%-25v  ...%v  \n", k+":", v)
		total += v
	}

	//adding tips

	formateString += fmt.Sprintf("%-25v  ...%v  \n", "tip :", bill.tips)

	///adding total bill

	formateString += fmt.Sprintf("%-25v  ...%0.2f", "Total", total+bill.tips)

	return formateString

}

//function for updating tips

func (bill *Bill) updateTip(tip float64) {
	bill.tips = tip
}

//add food item function

func (bill *Bill) addItems(name string, price float64) {
	bill.foodItems[name] = price
}

func (bill *Bill) save() {
	data := []byte(bill.format())
	err := os.WriteFile("bill/"+bill.Name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("Bill has been saved")
}
