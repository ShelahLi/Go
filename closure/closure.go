package main

import "fmt"
/**
	闭包的实现
	闭包的实现确保只要闭包还被使用，那么被闭包引用的变量会一直存在。
	如 sum的值会一直存在，有记忆效应
	函数 + 引用环境 = 闭包
 */
func Accumulate() func(int) int{

	sum := 0
	return func(i int) int {
		sum = sum + i
		return sum
	}
}
func main() {
	s := Accumulate()
	fmt.Println(s(1))
	fmt.Println(s(2))
}
