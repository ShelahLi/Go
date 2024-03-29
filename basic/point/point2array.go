package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	/*
		当想要维护一个数组，它可以存储指向int或字符串或任何其他可用的数据类型的指针。
		下面是一个指向整数的指针数组的声明：
	*/
	var ptr [MAX]*int

	/*
		这里将ptr声明为一个MAX整数的指针数组。 因此，ptr中的每个元素现在保存一个指向int值的指针。
	*/
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer. */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i])
	}
}
