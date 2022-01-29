package main

type Bill struct {
	Name      string
	foodItems map[string]float64
	tips      float64
	vat       float64
}

type Customer struct {
	customerName  string
	totalCustomer int16
	tableNumber   int8
}
