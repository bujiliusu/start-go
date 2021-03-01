package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	request,err := http.NewRequest(http.MethodGet,"http://www.imooc.com",nil)//同样返回结构体指针
	//将跳转指令放入requst中
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1")
	resp, err := http.DefaultClient.Do(request)//引用Client{}结构体 将Client结构体指针返回到变量中
	if err != nil{    //错误返回报告
		panic(err)
	}
	defer resp.Body.Close()  //预先处理resp关闭
	s, err := httputil.DumpResponse(resp,true)//分析结构体指针是否正确 如果正确 返回utf8格式
	if err != nil{
		panic(err)
	}

	fmt.Printf("%s", s)//打印utf8
}
