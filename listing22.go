// Code listing 22: https://play.golang.org/p/bOG8oBahGnm

package main

import "fmt"

func main() {

	plus := func(a int, b int) int {

		return a + b
	}

	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}