// Code listing 34: https://play.golang.org/p/9KT-MkAf3UJ

package main

import (
	"fmt"
)

func main() {
	var anything interface{} = "something"
	aInt := anything.(int)
	fmt.Println(aInt)
}