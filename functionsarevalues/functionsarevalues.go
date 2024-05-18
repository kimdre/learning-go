package main

import (
	"fmt"
)

type transformFunc func(int) int

func main()  {
	numbers := []int{1,2,3,4}
	moreNumbers:= []int{5,1,2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled, tripled)

	transorfmerFn1 := getTransformerFunction(&numbers)
	transorfmerFn2 := getTransformerFunction(&moreNumbers)

	tranformedNumbers := transformNumbers(&numbers, transorfmerFn1)
	moreTranformedNumbers := transformNumbers(&numbers, transorfmerFn2)
	fmt.Println(tranformedNumbers)
	fmt.Println(moreTranformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFunc {
	if(*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
