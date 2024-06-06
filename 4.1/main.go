package main

import (
	"fmt"
)

func main() {
	var number int
	_, _ = fmt.Scan(&number)
	for i := 1; i <= number; i++ {
		fmt.Println(i * i)
	}
}
