package panic_recovery

import (
	"errors"
	"fmt"
	"testing"
)
// 当心recover成为恶魔，形成僵尸服务进程，导致health check失效。"Let it Crash"往往是我们恢复不确定错误最好的方法。
func TestPanicVxExit(t *testing.T){
	defer func() {
		fmt.Println("Finally!")
		if err := recover(); err != nil{
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("Start!")
	panic(errors.New("Something wrong!"))
}
//package 基本复用单元，以首字母大写表明可被包外代码访问
//代码的package可以和所在的目录不一致
//同一目录的Go代码的package要保持一致