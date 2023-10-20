package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"

	fmt.Println(a)
	fmt.Println(a[0], a[1])

	primes := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes)

	var s []int = primes[1:4]
	fmt.Println(s)

	s = primes[:2]
	fmt.Println(s)

	s = primes[1:]
	fmt.Println(s)

	s = primes[:]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Leitão",
		"George",
		"Douglas",
	}

	fmt.Println(names)

	j := names[1:3]
	k := names[0:2]

	fmt.Println(j, k)

	k[0] = "XXX"
	fmt.Println(a, k)
	fmt.Println(names)
}
