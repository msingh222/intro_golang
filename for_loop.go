package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	if x < 0 {
		return math.Sqrt(2*(-x) + x )

	}
	return math.Sqrt(x)
}

func main () {

	sum := 0
	for sum<10   {
		sum = sum + 1
	}
	fmt.Print(sum)
	fmt.Print(sqrt(3), sqrt(-2))
}