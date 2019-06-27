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
	//file, err := os.Create(filename)
	// Openfile实际上就是Create，可以查看Create源码，0666权限
	file, err := os.OpenFile(filename,
		os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)

	//err = errors.New("this is customer error")
	if err != nil {
		/**
			如果pathError等于err.(*os.PathError)，ok为true，否则为false
		 */
		if pathError, ok := err.(*os.PathError); !ok {
			/**
				panic会引起程序崩溃，因此panic一般用于严重错误
			 */
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	// 函数退出之前执行关闭操作
	defer file.Close()

	//NewWriter创建一个具有默认大小缓冲、写入file的*Writer。
	writer := bufio.NewWriter(file)
	// 在文件关闭之前，将buffer.io里的数据写入文件
	defer writer.Flush()

	//写入buffer.io
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprint(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
