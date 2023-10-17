package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		fmt.Printf("%v°: %v\n", i, i)
		sum += i
	}
	fmt.Print(sum)
}
