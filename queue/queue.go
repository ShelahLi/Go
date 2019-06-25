package queue

/**
	通过别名扩充系统类型
 */
// A FIFO queue.
// Queue就是一个int的silce

/**
	将int改为interface后，该Queue支持任何类型
 */
type Queue []interface{}

// Pushes the element into the queue.
// 		e.g. q.Push(123)
/**
	该Queue支持所有类型，但是Push限定为int类型，把Push的入参改为int类型即可
	或者是*q = append(*q, v.(int))
 */
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// Pops element from head.
/**
	该Queue支持所有类型，但是Pop限定为int类型，在head后面.(int)
 */
//func (q *Queue) Pop() int {
//	head := (*q)[0]
//	*q = (*q)[1:]
//	return head.(int)
//}
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
