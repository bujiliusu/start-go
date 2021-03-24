package client

import (
	"start-go/ch15/series"
	"testing"
)

func TestPackage(t *testing.T){
	t.Log(series.GetFibonacciSerie(5))
}

//package 基本复用单元，以首字母大写表明可被包外代码访问
//代码的package可以和所在的目录不一致
//同一目录的Go代码的package要保持一致

//init方法
//在main被执行前，所有依赖的packge的init方法都会执行
//不同包的init函数按照包导入的依赖顺序执行
//每个包可以有多个init函数
//包的每个源文件可以有多个init函数，这点比较特殊