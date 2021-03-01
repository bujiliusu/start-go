package main

import "fmt"

//函数式编程实例，1、斐波那契数列，即输出一个函数2、为函数实现接口，即为函数添加功能
//3、使用函数遍历二叉树，即把函数作为参数传入，或者为函数注入不同函数实现不同的功能
//局部变量、自由变量
func adder() func(int) int {
	sum := 0
	return func(v int) int{
		sum += v
		return sum
	}
}
func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... +%d = %d\n", i, a(i))
	}
}

