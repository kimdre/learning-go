// A package that prints a greeting to you
package main

import (
	"fmt"
	"time"
)

// returns a nice, small greeting
func greeting() string {
	return "Hello, the time is " + time.Now().String()
}

func main() {
	fmt.Println(greeting())
}
