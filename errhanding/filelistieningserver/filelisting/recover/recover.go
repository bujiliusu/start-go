package main
//panic 停止当前函数执行，一直向上返回，并执行每一层的defer，如果没有recover，程序退出

//recover 仅在defer调用中使用，获取panic中的值，如何无法处理，就panic
// errors vs panic, 意料之中的：使用errors。如文件打不开 意料之外的：使用panic。如：数组越界
import (
	"fmt"
)

func tryRecover(){
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occured:", err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("this is a error"))
	//b := 0
	//a := 5 / b
	//fmt.Println(a)
	panic(123)

}
func main() {
	tryRecover()
}
