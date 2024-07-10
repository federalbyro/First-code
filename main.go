package main

import (
	"fmt"
)

func main() {
	var a, b int
	n := 1
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&n)
	max_b := (b + n - 1) / n
	if max_b >= a {
		fmt.Println("No")
	}
	if max_b < a {
		fmt.Println("Yes")
	}
}
