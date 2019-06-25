package tree

import "fmt"

/**
	定义节点的结构体
	*Node指针类型
 */
type Node struct {
	Value       int
	Left, Right *Node
}

// 为结构体定义函数：成员函数
// (node Node):接收者，相当于java的this。说明这个Print函数是给Node来接收的，是Node的成员函数
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil  {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	node.Value = value
}

// CreateNode不是Node的成员函数，因为它没有接受者
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
