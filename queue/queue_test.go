package queue

import "fmt"

/**
	更加特别的test测试
	示例代码ExampleQueue_Pop()
	Output对应上方的输出结果
 */
func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())


	// Output:
	// 0
	// 2
	// false
	// 3
	// true
}
