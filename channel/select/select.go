package main

/**
	用select进行调度
 */
import (
	"fmt"
	"math/rand"
	"time"
)

// 生成chan int
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	//var channel1, channel2 = generator(), generator()
	//for{
	//	select {
	//	case n := <- channel1:
	//		fmt.Println("Received from channel1", n)
	//	case n := <- channel2:
	//		fmt.Println("Received from channel2", n)
	//}

	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	// 10s之后会往channel里面送时间
	tm := time.After(10 * time.Second)
	// time.Tick定时，每一秒钟想看一下队列的长度
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}

		/**
			通过select来进行任务分发
		 */
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		// 800毫秒还没有获得数据，输出timeout
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		//	每一秒钟输出队列的长度
		case <-tick:
			fmt.Println(
				"queue len =", len(values))
		// 10s之后程序结束
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
