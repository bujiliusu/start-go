package extension

import (
	"fmt"
	"testing"
)
//拓展与复合
type Pet struct {

}
func (p *Pet) Speak(){
	fmt.Println("...")
}
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}
type Dog struct {
	Pet //复合，匿名嵌套类型,不支持重载，不支持LSP
}




func (d *Dog) Speak(){
	fmt.Println("wang")
}
//func (d *Dog) SpeakTo(host string) {
//	d.Speak()
//	fmt.Println(" ", host)
//}
func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
	dog.SpeakTo("nimei")
}

