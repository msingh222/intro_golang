package main

import "fmt"

func variable_dec() {
	var a, b bool
	var x,y = 7,8
	i, j := 5, 10
	fmt.Println(i,j,x,y,a,b)
}

var (
	ty,yu,tg,xs bool
)

func main() {
	variable_dec()
	fmt.Println(ty,yu,tg,xs)

}