package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main()  {
	userNames := make([]string, 2, 5) // 2: initial array length, 5: capacity/max array length

	userNames[0] = "Julie"
	// userNames[6] = "Luka"  <-- Would crash because out of bounce (higher than capacity defined in make() -> 5)

	userNames = append(userNames, "Kim")
	userNames = append(userNames, "Tobi")

	fmt.Println(userNames)

	courseRating := make(floatMap, 2)

	courseRating["go"] = 4.7
	courseRating["react"] = 4.8
	courseRating["test"] = 2

	courseRating.output()

	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courseRating {
		fmt.Println(key, value)
	}
}