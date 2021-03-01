package main

//接口变量里有啥，1、实现者类型，2、实现者的值或者指针
//接口变量自带指针
//指针变量同样采用值传递，几乎不需要使用接口的指针
//指针接受者实现只能以指针方式使用，值接受者都可

//struct可以组合，接口也可以
import (
	"fmt"
	"start-go/retriever/mock"
	"start-go/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

const url = "http://www.imooc.com"
func downlaod(r Retriever) string {
	return r.Get(url)
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster)  {
	poster.Post(url, map[string]string{
		"name": "ccmouse",
		"course": "golang",
	})
}
type RetrieverPoster interface {
	Retriever
	Poster
}
func session(s RetrieverPoster) string{
	s.Post(url, map[string]string{"contents":"nimei"})
	return s.Get(url)
}
func inspect(r Retriever){
	fmt.Println("Insepecting", r)
	fmt.Printf("> %T %v\n", r, r) //接口内部是有值的，还有类型
	fmt.Println("> type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println()
}
func main() {

	var r Retriever
	r = &mock.Retriever{"this is a fake immooc.com"}
	//实现了接口的结构体实例可以赋值给接口变量
	//fmt.Print(downlaod(r))
	inspect(r)
	//type assertion
	mockRetriever := r.(*mock.Retriever)
	fmt.Println(mockRetriever)
	fmt.Println(mockRetriever.Contents)

	fmt.Println("try a session")
	var retriever RetrieverPoster
	retriever = &mock.Retriever{"this is a fake immooc.com"}
	fmt.Println(session(retriever))
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	//inspect(r)


}
