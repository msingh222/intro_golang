package main

import "fmt"

func main()  {

	i,j := 21,31
	p := &i
	fmt.Print(*p)
	q := &j
	fmt.Print(*q+*p)


}
