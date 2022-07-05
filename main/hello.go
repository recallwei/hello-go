package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//a, b, c := initConst()
	//fmt.Println(a, b, c)

	// declareVariable()
	// getRuneArr("ABC")
	// getFloatBoundaryValue()
	// testEscapeCharacter()
	// testDefaultValue()
	// convertToInt32()
	// convertToInt()
	// printWithOrg()

	// _忽略函数返回值
	// a, _ := sumAndMul("1", "2")

}

// 函数指针
func changeName(name string) {
	name = "Bruce"
}

// 计算两数和、乘积
func sumAndMul(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
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

// 测试默认值
func testDefaultValue() {
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultBool, defaultFloat32, defaultFloat64, defaultString)
}

// 使用内置函数转换类型
func convertToInt32() {
	var integer16 int16 = 127
	var integer32 int32 = 32767
	fmt.Println(int32(integer16) + integer32)
}

// 使用 strconv 包将 string 转化为 int
func convertToInt() {
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}

// 访问命令行参数使用 os.org
func printWithOrg() {
	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])
	fmt.Println("Sum:", number1+number2)
}
