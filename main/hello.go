package main

import "fmt"

func main() {
	a, b, c := initConst()
	d := "Unused"
	fmt.Println(a, b, c)
	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Println(firstName, lastName, age)
}

func initConst() (int, int, int) {
	const (
		a = 1 << iota
		b = 1 << iota
		c = 1 << iota
	)
	return a, b, c
}
