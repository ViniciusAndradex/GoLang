package main

import "fmt"

func printSlice(array []int) {
	fmt.Printf("len = %d, cap=%d, %v\n", len(array), cap(array), array)
}

func main() {
	q := []int{2, 3, 5, 4, 7}
	fmt.Println(q)

	q = q[:0]
	printSlice(q)
	q = q[:4]
	printSlice(q)

	r := []bool{true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	var y []int
	fmt.Println(y, len(y), cap(y))
	if y == nil {
		fmt.Println("nil")
	}

	g := make([]int, 2, 5)
	print(g)
}
