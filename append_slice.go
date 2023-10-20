package main

import "fmt"

func printSlice(array []int) {
	fmt.Printf("len = %d, cap=%d, %v\n", len(array), cap(array), array)
}

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0, 5)
	printSlice(s)

	s = append(s, 4)
	printSlice(s)

	s = append(s, 2, 3, 5, 5)
	printSlice(s)

	var array = []int{1, 1, 2, 3, 4}

	for i, v := range array {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
	for _, v := range array {
		fmt.Printf("value: %d=\n", v)
	}
	for i, _ := range array {
		fmt.Printf("index: %d\n", i)
	}
	for i := range array {
		fmt.Printf("index: %d\n", i)
	}

}
