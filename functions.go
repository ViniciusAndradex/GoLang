package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x string, y int) (string, int) {
	y = y + 8
	return x + " Teste", y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	a, b := swap("Sou eu", 40)
	fmt.Println(a, b)
	fmt.Println(split(17))
}
