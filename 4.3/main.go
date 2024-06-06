package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	product := 1
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			product *= i
		}
	}
	fmt.Println(product)
}
