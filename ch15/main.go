package main

import (
	"fmt"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
}
// 第一种定义方式在实例对应方法调用时，实例的成员会进行值复制
func (e Employee) String() string{
	fmt.Printf("address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Nane:%s-Age:%d", e.Id, e.Name, e.Age)
}
//通常情况下为了避免内存拷贝我们使用第二种方式
func (e *Employee) String2() string{
	fmt.Printf("address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Nane:%s-Age:%d", e.Id, e.Name, e.Age)
}
func main() {
	e := Employee{"0", "Bob", 20}
	//e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) //注意这里返回的引用/指针，相当于 e:= &Employee{}
	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 22
	//fmt.Println(e, e1)
	fmt.Println(e.String())
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))

	fmt.Println(e.String2())
	fmt.Printf("address is %x", unsafe.Pointer(&e.Name))



	//fmt.Println(e2.String())
	//fmt.Println(e2.String2())

}
