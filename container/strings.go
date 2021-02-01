package main

import (
	"fmt"
	"unicode/utf8"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
func main() {
	s := "Ilove天天"

	fmt.Printf(typeof(s))
	fmt.Println()
	fmt.Println(len(s))
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", []byte(s))//字符序列
	fmt.Printf("%X\n", []byte(s))
	fmt.Println("11111111111111111")
	for _, b := range []byte(s) { // b is a byte //底层，字符序列
		fmt.Printf("%X ", b)
	}

	fmt.Println("\n22222222222222")
	for i, ch := range s { // ch is a rune, unicode是字符集，utf-8是编码规则，是具体的存储方案。//以utf-8解码后又编码成unicode，每个中文字符算3个
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	fmt.Println("rune count:",utf8.RuneCountInString(s))

	bytes := []byte(s)
	ch, size := utf8.DecodeRune(bytes) //上述代码中utf8.DecodeRune的作用是通过传入的utf8字节序列转为一个rune即unicode。
	fmt.Printf("%c %d\n",ch, size)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c %d", ch, size)
	}

	fmt.Println()

	for i, ch := range []rune(s) { //上层，新生成的rune数组，utf-8转发为unicode
		fmt.Printf("(%d %c )", i, ch)
	}

}
