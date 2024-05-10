package main

import (
	"math"
	"fmt"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return value: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue :=  investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println("Expected return value:", futureValue)
	fmt.Println("Future real value:", futureRealValue)
}