package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello from goroutine %d\n", i)
				//io等待会出现切换,出现协程切换
				//a[i]++
				//非抢占，会锁死在第一个协程内
				//runtime.Gosched()
				//交出控制权
			}
		}(i)
		//不传入i的话，就是闭包，引入了外部变量
		//当for循环结束的时候，i=10，a[10]出现了超界
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
	//main也是一个协程，也会一起调度
}
//协程coroutine,编译器级别的多任务调度
//轻量级线程
//非抢占式任务处理，由协程主动交出控制权
//多个协程可能在一个或多个线程运行
//go run -race *.go,检测数据访问冲突
//使用channel解决并发访问的问题

//子程序是协程的一个特例
//以main调用dowork为例
//普通函数，运行在一个线程，main单向调用dowork
//协程，man与dowork双向调用，可能运行在多个线程里

//其他语言的协程
//python，使用yield关键字实现协程
//python3.5加入async def对协程的原生支持

//不需要在定义的时候区分是否是异步函数，相对与python
//调度器在合适的点进行切换
//使用-race检测数据冲突

//goroutine可能的切换点
//i/o,select,channel,等待锁,函数调用（有时）,runtime.Gosched()
