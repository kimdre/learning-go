package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)

// 	// fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	higlightedPrices := featuredPrices[:1]
// 	fmt.Println(higlightedPrices)
// 	fmt.Println(len(higlightedPrices), cap(higlightedPrices))
// }
