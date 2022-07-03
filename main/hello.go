package main

import (
	"fmt"
	"math"
)

func main() {
	//a, b, c := initConst()
	//fmt.Println(a, b, c)

	// declareVariable()
	// getRuneArr("ABC")
	getFloatMaxValue()
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
func getRuneArr(str string) {
	rune := []rune(str)
	fmt.Println(rune)
}

//
func getFloatMaxValue() {
	fmt.Println(math.MaxFloat32, math.MaxFloat64)
}
