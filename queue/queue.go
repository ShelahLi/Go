package queue

/**
	通过别名扩充系统类型
 */
// A FIFO queue.
// Queue就是一个int的silce
type Queue []int

// Pushes the element into the queue.
// 		e.g. q.Push(123)
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

// Pops element from head.
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
