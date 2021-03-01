package interface_test

import "testing"
//go接口与其他语言的区别
//1。接口非入侵式，实现不依赖与接口定义
//2。所以接口的定义可以包含在接口使用者包内
type Programmer interface {
	WriteHelloWorld() string
}
type GoProgrammer struct {

}
func (gopro *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hello world\")"
}
func TestClient(t *testing.T){
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}