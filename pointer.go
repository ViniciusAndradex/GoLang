package main

func main() {
	var p *int
	i := 42
	p = &i
	println(i)
	*p = 21
	println(i)

	k, j := 42, 2701

	l := &k
	println("l: ", *l)
	*l = 21
	println(k)

	l = &j
	*l = *l / 37
	println("j: ", j)
}
