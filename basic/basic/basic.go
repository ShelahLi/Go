package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	//函数外面也可以定义变量，但是函数外面定义变量不能用(:=)
	//这些变量不是全局变量，只是包类变量,在包类使用
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	//变量名在前，类型在后
	//函数内的变量只能在函数内部使用，函数外面无法使用
	var a int
	var s string
	//printf可以跟变量的格式%d，%q。%q可以把空字符串打印出来，是""。
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	//变量定义了一定要用，不用会报错
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	//go语言可以推断变量的类型(a,b := 1,"abc")
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	//更简单的写法，省略var，第一使用可以用(:=)，第二次赋值就不能再用了。
	//:相当于声明这个变量
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	//复数c
	c := 3 + 4i
	//复数包：cmplx, Abs取模
	fmt.Println(cmplx.Abs(c))
	//1i
	fmt.Print("欧拉公式: ")
	//%.3f：只要小数点后三位
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {

	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a , b int) int {
	var c int
	//强制类型转换：float64(), int()
	//Sqrt():需要float64类型
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	// go中常量类型定义 const(...)
	// 常量名称不需要大写,go中大小写有特殊含义
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	// 如果定义了c是int类型，就要强制转换
	//由于定义a, b时没有规定类型，所以sqrt会把a, b当做float用
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	// 枚举类型的定义
	//go中没有枚举类型，用const来代替，const变量后面必须要赋值
	// iota：表示这组变量是自增值，所以没有赋值, 自增值从0开始
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		// b等于1左移10乘以iota位，2进制
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	consts()
	enums()
}
