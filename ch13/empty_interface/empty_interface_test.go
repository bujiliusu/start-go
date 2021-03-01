package empty_interface

import (
	"fmt"
	"testing"
)

//空接口与断言
//空接口可以便是任何类型
//通过断言来将空接口转换成制定类型，v, ok := p.(int)

func DoSomething(p interface{}){
	//if i,ok := p.(int); ok{
	//	fmt.Println(i)
	//	return
	//}
	//fmt.Print("unknown")
	switch v:=p.(type){
	case int:
		fmt.Println(v)
	default:
		fmt.Print("unknown")
	}
}
func TestDo(t *testing.T){
	DoSomething(10)
	DoSomething("nimei")
}