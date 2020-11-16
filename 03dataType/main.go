package main

import "fmt"

func main() {
	//十进制数打印为二进制
	n := 10
	fmt.Printf("%b\n", n)

	//八进制数打印为八进制
	m := 077
	fmt.Printf("%o\n", m)
	//八进制数打印为十进制
	fmt.Printf("%d\n", m)
	//八进制数打印为二进制
	fmt.Printf("%b\n", m)

	//十六进制数打印为十六进制
	f := 0xff
	fmt.Printf("%x\n", a)
}
