package tree

//改造，使用函数遍历二叉树
func (node *Node) Traverse(){
	node.TraverseFunc(func(node *Node) {
		node.Print()
	})
}

func (node * Node) TraverseFunc(f func(node *Node)){
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}
