package main

import (
	"fmt"
	"regexp"
)

const text = `my email is ccmouse@gmail.com@abc.com
email1 is abc@def.org
email2 is kkk@qq.com.cn`
func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	math := re.FindAllStringSubmatch(text, -1)
	fmt.Println(math)
}
