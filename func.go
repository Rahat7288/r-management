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


//bill formating function 

func (bill *Bill) format() string{

	formateString := "Bill Breakdown:  \n"

	total := 0

	//fetching the foodItems 

	for k,v: range bill.foodItems {
		fs += fmt.Sprintf("%-25v  ...%v  \n",k+,":",v )
		total += v
	}

	//adding tips 

	fs += fmt.Sprintf("%-25v  ...%v  \n","tip :",bill.tips)


	///adding total bill

	fs += fmt.Sprintf("%-25v  ...%0.2f","Total", total+bill.tips)


	return formatingString



}

//function for updating tips 

func (bill *Bill) updateTip(tip float64){
	bill.tips = tip
}



//