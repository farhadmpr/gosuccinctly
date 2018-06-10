// Code listing 12: https://play.golang.org/p/kHqSobJC2Yv

package main

import (
	"fmt"
)

func main() {
	mySlice := []string{"Apples", "Oranges", "Bananas"}
	fmt.Printf("Initial slice values: %v\n", mySlice)
	myFunction(mySlice)
	fmt.Printf("Final slice values: %v\n", mySlice)
}

func myFunction(fruits []string) {
	// Change Oranges to Strawberries
	fruits[1] = "Strawberries"
	fmt.Printf("Slice values in myFunction(): %v\n", fruits)
}