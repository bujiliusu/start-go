package tree

import (
	"fmt"
)
//值接受着VS指针接受者，只有go才有值接受者，值/指针接受者均可接受值指针
type Node struct {
	Value int
	Left,Right *Node
}
func (node Node) Print(){ //结构体Node的方法，node Node 表示接受者
	fmt.Print(node.Value)
}
func print1(node Node){ //效果跟print一致，调用方法：print1(root)
	fmt.Println(node.Value)
}

func (node *Node) SetValue(Value int){ //只有指针可以改变结构内容，nil指针也可以调用方法
	if node == nil { //值不会是nil，指针可以是nil
		fmt.Println("Setting nil, ignore")
		return
	}
	node.Value = Value
}
func CreateNode(Value int) *Node{
	return &Node{Value: Value} //返回局部变量的地址，工厂函数
}

