package main

import (
	"fmt"
	"runtime"
	"time"
)
/**
	goroutine: go func():并发的去执行这个函数
 */
func main() {

	for i := 0; i < 1000; i++ {
		go func(i int) { //开启多线程
			for {
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i) // (i)的作用：将for里面的i传到func(i int)中的i去
	}
	time.Sleep(time.Minute)//休眠：否则main函数执行太快，其他线程还没执行完，main函数已经结束了，数据无法被打印

	var a [10]int
	for i:= 0; i < 10; i++{
		go func(i int) {
			for {
				a[i]++
				// 手动交出控制权runtime.Gosched()
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}
