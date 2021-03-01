package custome_type

import (
	"fmt"
	"testing"
	"time"
)
//自定义类型，或者函数的别名
type IntConv func (op int) int
func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println(time.Since(start).Seconds(), "yyyyyyyyyyyyyyyyyy")
		return ret

	}
}

func slowFun(op int) int{
	time.Sleep(time.Second*1)
	return op
}
func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
