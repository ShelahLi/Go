package main

import "fmt"

/**
adder函数：没有入参，返回是一个函数，这个函数入参是int，返回int
*/
func adder() func(int) int {
	// sum: 自由变量
	sum := 0
	return func(v int) int {
		// v: 局部变量
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {

	s := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(s(i))
	}
	fmt.Println(s(11));

	// a := adder() is trivial and also works.
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}
}
