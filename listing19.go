// Code listing 19: https://play.golang.org/p/ueJmA28U5EF

package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	delete(m, "k2")

	if _, ok := m["k2"]; ok {
		fmt.Println("Ok")
	} else {
		fmt.Println("NotOk")
	}
}