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
	// getFloatBoundaryValue()
	testEscapeCharacter()
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

// 获取浮点数边界值
func getFloatBoundaryValue() {
	fmt.Println(math.MaxFloat32, math.MaxFloat64, math.SmallestNonzeroFloat32, math.SmallestNonzeroFloat64)
}

// 测试转义字符
func testEscapeCharacter() {
	fullName := "John Doe \t(alias \"Foo\")\n"
	fmt.Println(fullName)
}
