package tree

import "fmt"

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}

/**
	闭包遍历树
*/
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	fmt.Printf("%d ",node.Value)
	node.Right.TraverseFunc(f)
}

/**
	使用channel实现树的遍历
 */
func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(
			// 在闭包里面实现另一个函数:f(node)，将node传入chan中
			func(node *Node) {
				out <- node
			})
		close(out)
	}()
	return out
}
