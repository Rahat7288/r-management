package main

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
