package main

import (
	"errors"
	"fmt"
)

func tryRecover() {
	/**
		使用recover必须在defer里面
	 */
	 defer func() {
		r := recover()
		// recover可以获取panic的值，但panic不一定存在，所以recover可能为空
		if r == nil {
			fmt.Println("Nothing to recover. " +
				"Please try uncomment errors " +
				"below.")
			return
		}
		// 如果panic存在
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}()

	//
	panic(errors.New("this is an error"))

	// Division by zero
	//b := 0
	//a := 5 / b
	//fmt.Println(a)

	// Causes re-panic
	//panic(123)


}

func main() {
	tryRecover()
	fmt.Println("over!")
}
