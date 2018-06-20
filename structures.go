package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{34,45}
	v2 = Vertex{100 ,222}
	v3 = Vertex{}
	v4 = &Vertex{78,85}
)

func main()  {
	//p:=&Vertex{1,2}
	v:=Vertex{23,34}
	q:=&v
	q.X =345
	fmt.Print(v)

	fmt.Print(Vertex{1,2})

	fmt.Println(v1,v2,v3,*v4)
	m = make(map[string]Vertex)
}

var m map[string]Vertex




