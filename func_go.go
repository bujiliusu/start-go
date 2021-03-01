package main

import "fmt"

//go只有值传递，值传递结合指针实现饮用传递
//值传递，var a int --> func f(a int)
//var a int --> func(pa *int)，拷贝&a的值给到pa
//var cache Cache --> func f(cache Cache) 一般无法拷贝cache的值，所以cache存放的一般是一个指向数据的指针
//go自定义对象的时候，注意是当值用还是当指针用，即是值封装还是指针封装

func swap(a, b int) (int, int){
	//a, b = b, a
	return b, a
}
func swap1(a, b *int){
	*a, *b = *b, *a
}
func main() {
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)

	swap1(&a, &b)
	fmt.Println(a, b)
}
