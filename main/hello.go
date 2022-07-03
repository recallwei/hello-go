package main

import "fmt"

func main() {
	//a, b, c := initConst()
	//fmt.Println(a, b, c)

	getRune()
	// declareVariable()
}

// variable declaration
func declareVariable() {
	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Println(firstName, lastName, age)
}

// iota
func initConst() (int, int, int) {
	const (
		a = 1 << iota
		b = 1 << iota
		c = 1 << iota
	)
	return a, b, c
}

// rune
func getRune() {
	rune := 'G'
	fmt.Println(rune)
}
