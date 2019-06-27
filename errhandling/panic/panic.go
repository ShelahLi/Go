package main

import "fmt"

func main() {
	/**
	宕机前，defer 语句会优先被执行，由于第 12 行的 defer 后执行，
	因此会在宕机前，这个 defer 会优先处理，随后才是第 11 行的 defer 对应的语句。
	这个特性可以用来在宕机发生前进行宕机信息处理。
	 */
	defer fmt.Println("宕机后要做的事情1")
	defer fmt.Println("宕机后要做的事情2")
	panic("宕机")
}
