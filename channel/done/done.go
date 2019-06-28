package main
/**
	通过通信来共享内存
 */
import (
	"fmt"
	"sync"
)

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		/**
			每个任务做完之后就要done
		 */
		w.done()
	}
}

type worker struct {
	in   chan int
	/**
		把*sync.WaitGroup的Done方法包装在done里面
	 */
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	/**
	sync.WaitGroup.Add(任务数)：10个小写+10个大写
	 */
	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	/**
		sync.WaitGroup.Wait():等待这20个任务全部做完
	 */
	wg.Wait()
}

func main() {
	chanDemo()
}
