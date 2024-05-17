package main

import (
	"fmt"
)

func main() {
	prices := []float64{1.00, 1.99, 14.50}
	fmt.Println(prices)

	discountPrices := []float64{12.99, 18.49, 3.66}
	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}
