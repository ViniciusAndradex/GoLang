package main

import "fmt"

func adder(sum int) func(int) int {
	//sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(2), adder(3)

	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
