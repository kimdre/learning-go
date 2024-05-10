package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Println("-- Investment Calculator --")

	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter expected return value: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Expected return value: %.2f €\n", futureValue)
	fmt.Printf("Future real value: %.2f €\n", futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)

	return fv, frv
}
