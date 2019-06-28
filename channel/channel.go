package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	// 方法1：从<-c接收两个值，判断ok是否还有值
	//n, ok := <-c
	//if !ok {
	//	break
	//}
	// 方法2：range可以自动检测channel的close
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

/**
	chan<- int:表示这个channel是用来发数据的；只允许这样的形式：channel <- 1
	<-chan int:表示这个channel是用来收数据的；只允许这样的形式：n := <- channel
*/
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func chanDemo1(){
	// 创建一个channel of int
	channel := make(chan int)
	// 往channel发数据,由于channel是goroutine与goroutine之间的交互，为了避免死锁
	// 需要另一个goroutine接收数据
	/**
	开启一个goroutine
	*/
	go func() {
		// 从channel读取数据
		for {
			c := <- channel
			fmt.Println(c)
		}
	}()
	//go worker(0,channel)
	channel <- 1
	channel <- 2

	time.Sleep(time.Millisecond)

}

func chanDemo2()  {
	// channel作为数组
	var channels [10] chan int
	for i:=0;i<10;i++{
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for i:=0;i<10;i++{
		channels[i] <- 'a' + i
	}
	for i:=0;i<10;i++{
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}
func chanDemo3() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}


func bufferedChannel() {
	// make(chan int, 3): 创建大小为3的channel缓冲区；不创建goroutine也不会报错，但是传入超过缓冲区大小就会报错
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	// channel是可以关闭的
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo1()
	chanDemo2()
	chanDemo3()
	fmt.Println("Buffered channel")
	bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
