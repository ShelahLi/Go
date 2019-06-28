package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex // 锁
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	/**
		使用匿名函数在一块代码区里面对a.value进行保护
	 */
	func() {
		// 在increment之前上锁，防止其他线程对其进行操作
		a.lock.Lock()
		// defer控制在匿名函数里面，函数退出之前解锁
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	// 在get之前上锁，防止其他线程对其进行操作
	a.lock.Lock()
	// 函数退出之前解锁
	defer a.lock.Unlock()

	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
		//a.value++
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
/**
	main()与go func(){}并发，a.value++时可能遇到a.get()，造成数据冲突
 */
