package main

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
