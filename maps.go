package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	fmt.Println(m)
	m = make(map[string]Vertex)
	fmt.Println(m)
	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])

	x := make(map[string][]int)
	x["Teste"] = append(x["Teste"], 0, 0, 7)
	println(x)
	println(x["Teste"])

	delete(x, "Teste")
	println(x)
	println(x["Teste"])

}
