// Code listing 33: https://play.golang.org/p/njiry2LJ_qk

package main

import (
	"fmt"
)

func main() {
	var anything interface{} = "something"
	aString := anything.(string)
	fmt.Println(aString)
}