package main

import (
	"fmt"
)

func main() {
	var root tree.TreeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	//root.right.left.setValue(4)//自动取地址
	//root.right.left.print()

	//root.setValue(100)
	//root.print()
	fmt.Println("11111111111")
	root.traverse()

	nodes := []treeNode{
		{value: 5},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	var pRoot *treeNode
	typeof(pRoot)
	fmt.Println(pRoot)
	pRoot = &root
	pRoot.setValue(200)
	pRoot.print()



}
