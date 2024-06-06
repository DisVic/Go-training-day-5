package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	counter := true
	for i := 1; i <= n; i++ {
		var number int
		_, _ = fmt.Scan(&number)
		if number == 0 {
			counter = false
			break
		}
	}
	if counter == false {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
