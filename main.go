package main

import (
	"fmt"
)

type transformFunc func(int) int

func main()  {
	numbers := []int{1,2,3,4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled, tripled)
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
