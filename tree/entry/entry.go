package main

import (
	"fmt"
	"start-go/tree"
)

//为结构定义的方法必须放在一个包内，可是是不同文件
//扩充类型，1、使用别名。2、使用组合
//只有变量才能取得地址

type myTreeNode struct {
	node *tree.Node
}
func (myNode *myTreeNode) postOrder(){
		if myNode == nil {
			return
		}
		if myNode.node == nil {
			return
		}
		left := myTreeNode{myNode.node.Left}
		left.postOrder()
		right := myTreeNode{myNode.node.Right}
		right.postOrder()
		myNode.node.Print()
}
func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}

	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4) //会自动把root.Right.lefe的地址传给setValue

	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount ++
	})
	fmt.Print(nodeCount)

}

