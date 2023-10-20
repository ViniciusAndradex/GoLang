package main

import "fmt"

type Vertex struct {
	X int // Fields
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	x  = &Vertex{1, 2}
)

func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{4, 5}
	v.X = 25
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, v2, v3, x)
}
