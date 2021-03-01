package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"
	fmt.Printf("%T", s)
	fmt.Println(s)
	r := strings.NewReader(s)
	fmt.Println(r)
	fmt.Println(r.Size())
	fmt.Println(r.Len())
	fmt.Println(len(s))
	for r.Len() > 5 {
		b, err := r.ReadByte()
		fmt.Println(string(b), err, r.Len(), r.Size())
	}
	remainStr := make([]byte, 5)
	n, err := r.Read(remainStr)
	fmt.Println(string(remainStr), n, err)
	fmt.Println(r.Size())
	fmt.Println(r.Len())
	fmt.Printf("%T", r)

	s = string("nimei")
	fmt.Printf("%T", s)
	fmt.Println(s)
}
