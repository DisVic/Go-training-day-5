package main

import (
	"fmt"
)

func main() {
	var w uint16
	_, _ = fmt.Scan(&w)
	if w%2 == 0 && w/2 != 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
