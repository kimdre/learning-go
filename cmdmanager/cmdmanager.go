package cmdmanager

import "fmt"

type CmdManager struct{}

func (cmd CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		prices = append(prices, price)

		if price == "0" {
			break
		}
	}

	return prices, nil
}

func (cmd CmdManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CmdManager {
	return CmdManager{}
}
