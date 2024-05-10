package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64 = 19  // tax rate percentage

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	// Earning before tax
	ebt := revenue - expenses

	// Earning after tax
	profit := ebt * (1 - taxRate / 100)

	// Ratio
	ratio := ebt / profit

	fmt.Printf("Earnings before tax: %v €\n", ebt)
	fmt.Printf("Profit after tax: %v €\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)

	var a [2]string

	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
}

