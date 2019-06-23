package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//函数的定义：函数名(输入参数)(返回参数1，返回参数2)
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		// fmt.Errorf(), error类型
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数的输入参数嵌套函数：
// op是一个函数，输入参数两个int，返回一个int
func apply(op func(int, int) int, a, b int) int {
	// 获取函数的指针
	p := reflect.ValueOf(op).Pointer()
	// 获取函数的名
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+
		"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	// range numbers表示输入参数列表里的每一个数的下标
	for i := range numbers {
		s += numbers[i]
	}
	return s
}



func main() {
	fmt.Println("Error handling")
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Printf("13 div 3 is %d mod %d\n", q, r)

	fmt.Println("pow(3, 4) is:", apply(func(a int, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println("1+2+...+5 =", sum(1, 2, 3, 4, 5))

	
}
