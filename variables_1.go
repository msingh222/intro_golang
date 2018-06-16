package main

import (
	"fmt"
	"math"
)

func main()  {
	var x  = 6.0
	var y = 7.0
	var f  = math.Sqrt(float64(x * y ))
	fmt.Println(f)
	fmt.Printf("x is of type %T and %v", x, x)

}
