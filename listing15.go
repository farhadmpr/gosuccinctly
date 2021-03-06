// Code listing 15: https://play.golang.org/p/yE-1Rhccebv

package main

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 0, 8)
	mySlice = append(mySlice, 1, 3, 5, 7, 9, 11, 13, 17)

	fmt.Printf("Contents: %v\n", mySlice)
	fmt.Printf("Number of Items: %d\n", len(mySlice))
	fmt.Printf("Capacity: %d\n", cap(mySlice))

	mySlice = append(mySlice, 19)

	fmt.Printf("Contents: %v\n", mySlice)
	fmt.Printf("Number of Items: %d\n", len(mySlice))
	fmt.Printf("Capacity: %d\n", cap(mySlice))

}