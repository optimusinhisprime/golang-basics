package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput(reader, "Create a new bill name: ")

	b := newBill(name)
	fmt.Println("Created the new bill - ", name)

	return b

}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput(reader, "Choose option (a - add item, s - save bill, t - add tip): ")

	switch opt {
	case "a":
		name, _ := getInput(reader, "Item Name: ")
		price, _ := getInput(reader, "Item Price: ")

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number.")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added -", name, p)
		promptOptions(b)

	case "s":
		b.saveBill()
		fmt.Println("You saved the bill file.")
	case "t":
		tip, _ := getInput(reader, "Enter tip amount: ")

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number.")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip has been updated - ", t)
		promptOptions(b)

	default:
		fmt.Println("That was not a valid option.")
		promptOptions(b)
	}
}

func main() {

	myBill := createBill()
	promptOptions(myBill)

}
