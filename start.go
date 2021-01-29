package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
}
func swap2(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	fmt.Println("hello wprld")
	//指针不能计算
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
	c, b := 3, 4
	swap(c, b)
	fmt.Println(c, b)
	swap2(&c, &b)
	fmt.Println(c, b)
}
