package main

import (
	"fmt"
	"math"
)

func main() {
	var d1, d2, d3 float64
	_, _ = fmt.Scan(&d1, &d2, &d3)
	dist := math.Min(d1, d2)
	if d3 > d1+d2 {
		dist += math.Min(d1, d2)
		dist += (math.Max(d1, d2) * 2)
	} else {
		dist += d3
		if math.Max(d1, d2) > (d3 + math.Min(d1, d2)) {
			dist += (d3 + math.Min(d1, d2))
		} else {
			dist += math.Max(d1, d2)
		}
	}
	fmt.Println(int(dist))
}
