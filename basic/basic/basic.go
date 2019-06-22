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


func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)

	euler()
	triangle()
	
}
