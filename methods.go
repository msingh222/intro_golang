package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}


func (v Vertex) Abs() float64{
	return math.Sqrt(v.X*v.Y + v.X*v.Y)
}

func Abs1(v Vertex)  float64{
	return math.Sqrt(v.X*v.Y + v.X*v.Y)
}

func main()  {
	v:= Vertex{3,5}
	fmt.Println(v.Abs())
	fmt.Println(Abs1(v))

}