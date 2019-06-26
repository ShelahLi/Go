package main

import (
	"fmt"
	"os"

	"bufio"

	"learngo/functional/fib"
)

func tryDefer() {
	/**
		defer: 在函数结束之间，调用；有多个defer，先进后出
	 */
	for i := 0; i < 100; i++ {
		/**
			defer语句参与计算
		 */
		 defer fmt.Println(i)
		if i == 30 {
			/**
				即使有panic：defer也照样执行
			 */
			//panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	//创建文件
	file, err := os.Create(filename)
	if err != nil{
		panic(err)
	}

	// 函数退出之前执行关闭操作
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 在文件关闭之前，将buffer.io里的数据写入文件
	defer writer.Flush()

	//写入buffer.io
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
