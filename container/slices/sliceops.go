package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func sliceOps() {
	fmt.Println("Creating slice")
	// 创建切片
	var s []int // Zero value for slice is nil

	for i := 0; i < 10; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	//创建有初始值得slice
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	// make: 创建长度16，cap为16的切面
	s2 := make([]int, 16)
	// make: 创建长度为10，cap为32的切片
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	// copy: 复制s1的值到s2
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	//删除第三位：前两位+第4位开始往后
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	// 删除第一个元素
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	//删除最后一个元素
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}

func main()  {
	sliceOps()
}
