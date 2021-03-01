package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printFile(filename string){
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	printFileContents(file)
}
func printFileContents(reader io.Reader){
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	s := `ab''""
	
	sdfasdf  lll`
	printFileContents(strings.NewReader(s))

	printFile("basic/loop/abc.txt")
}
