package main

import "fmt"

func add(x, y int) int {
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("now.")
	return x + y
}

func main() {
	defer fmt.Println("World!")

	fmt.Println(add(4, 5))

	fmt.Print("Hello ")
}
