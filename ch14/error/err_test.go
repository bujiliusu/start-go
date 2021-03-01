package error



//及早失败，避免嵌套
// panic用于不可恢复的错误，panic退出前会执行defer指定的内容
//panic vs is.Exit os.Exit 退出时不会调用defer指定的函数，os.Exit退出是不会输出当前调用栈信息
import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func GetFibonacci(n int) ([]int, error)  {
	if n < 2 || n > 100 {
		return nil,errors.New("n should be in [2, 100]")
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil

}

func GetFibonacci1(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}
func GetFibonacci2(str string) {
	var (
		i int
		err error
		list []int
	)
	if i, err = strconv.Atoi(str); err == nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err == nil {
		fmt.Println(list)
		return
	}
	fmt.Println(list)
}
func TestGetFib(t *testing.T)  {
	if v, err := GetFibonacci(10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
