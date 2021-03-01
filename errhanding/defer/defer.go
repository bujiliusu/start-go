package main
//defer调用，确保调用在函数结束时发生
//参数在defer语句时计算
//defer列表为后进先出
import (
	"bufio"
	"fmt"
	"os"
	"start-go/functional/fib"
)

func tryDefer(){
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}

}
func writeFile(filename string){
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error:", err.Error() )
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()
	//fmt.Printf("%T", file)
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	//fmt.Printf("%T", writer)


	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}
func main() {
	//tryDefer()
	writeFile("def.txt")
}
