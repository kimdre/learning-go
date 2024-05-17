package main

import (
	"fmt"
)

type Product struct {
	title string
	id int
	price float64
}

func newProduct(title string, id int, price float64) Product {
	return Product{
		title: title,
		id: id,
		price: price,
	}
}

func main() {

	hobbies := [3]string{"mountainbike", "development", "documentaries"}
	fmt.Println(hobbies)

	fmt.Println(hobbies[0])
	otherHobbies := hobbies[1:3]
	fmt.Println(otherHobbies)

	sliceOne := hobbies[:1]
	fmt.Println(sliceOne)

	sliceTwo := hobbies[0:1]
	fmt.Println(sliceTwo)

	fourSlice := sliceOne[1:3]
	fmt.Println(fourSlice)

	goals := []string{"learning", "knowledge"}
	goals[1] = "foo"
	goals = append(goals, "bar")
	fmt.Println(goals)

	products := []Product{
		{
			title: "Apfel",
			id: 1,
			price: 0.49,
		},
		{
			title: "Birne",
			id: 2,
			price: 0.79,
		},
	}

	products = append(products, newProduct("Zitrone", 3, 1.19))

	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.


// func main() {
// 	prices := []float64{10.99, 8.99}
// 	fmt.Println(prices[0:1])
// 	prices[1] = 9.99

// 	prices = append(prices, 5.99)
// 	fmt.Println(prices)
// }

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
