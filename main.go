package main

import "fmt"

func main()  {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	// userNames[6] = "Luka"  <-- Would crash because out of bounce (higher than capacity defined in make() -> 5)

	userNames = append(userNames, "Kim")
	userNames = append(userNames, "Tobi")

	fmt.Println(userNames)
}