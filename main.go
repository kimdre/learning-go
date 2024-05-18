package main

import "fmt"

func main() {
	// sum := sumup(numbers)
	
	sum := sumup(4, 10, 6, 20)
	fmt.Println(sum)
	
	numbers := []int{4,10,6}
	anotherSum := sumup(1, numbers...)
	fmt.Println(anotherSum)
}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
