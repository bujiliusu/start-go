package tree

import (
	"fmt"
)

type TreeNode struct {
	value int
	left, right *TreeNode
}
func (node TreeNode) print() {
	fmt.Println(node.value)
}
func (node *TreeNode) setValue(value int) {
	if node == nil {
		fmt.Println("ignore")
		return//nil无法拿到value，所以需要返回
	}
	node.value = value
}
func createNode(value int) *TreeNode {
	return &TreeNode{value: value} //此局部变量由编译器决定分配在栈上还是堆上
}

func (node *TreeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

