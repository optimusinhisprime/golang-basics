package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//make new bills

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func (b *bill) format() string {
	fs := "Bill Breakdown: \n"
	total := 0.0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ... $%v \n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "Total:", total+b.tip)

	return fs
}

// update bill
func (b *bill) updateTip(tip float64) {
	// When working with structs we dont need to dereference the pointer. Go automatically does it for us.
	// (*b).tip= tip

	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, amount float64) {
	b.items[name] = amount
}
