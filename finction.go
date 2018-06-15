package main

import "fmt"

func addition(x int , y int) int {
	return x + y

}

func swap(a,b string)  (string,string) {
	return b, a

}
func main ()  {
	fmt.Println("7 + 10  =  ", addition(7,10))
	a, b := swap("hello","world")

	fmt.Println(a,b)

}