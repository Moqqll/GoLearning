package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	//Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件错误, err：", err)
		return
	}
	name := "沙河小王子"
	//向打开的文件句柄中写入内容main
	fmt.Fprintf(fileObj, "往文件中写入信息：%s", name)

	//Sprint系列函数会把传入的数据生成并返回一个字符串。
	s1 := fmt.Sprint("moqqll")
	name1 := "沙河娜扎"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name1, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1, s2, s3)

	//Errorf
	err1 := fmt.Errorf("这是一个错误")
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误：%w", e)
	fmt.Println(err1)
	fmt.Println(w)

	//格式化占位符
	//通用占位符
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"moqqll"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%+v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")
	//布尔型
	fmt.Printf("%t\n", true)
	//整型
	n := 999999
	fmt.Printf("二进制：%b\n", n)
	fmt.Printf("Unicode码值：%c\n", n)
	fmt.Printf("十进制：%d\n", n)
	fmt.Printf("八进制：%#o\n", n)
	fmt.Printf("十六进制：%#x\n", n)
	fmt.Printf("十六进制：%X\n", n)
	fmt.Printf("Unicode：%U\n", n)

	m := 12.34
	fmt.Printf("%f\n", m)
	fmt.Printf("%9f\n", m)
	fmt.Printf("%.2f\n", m)
	fmt.Printf("%9.2f\n", m)
	fmt.Printf("%9.f\n", m)

	//Scan
	var (
		nameq    string
		ageq     int
		marriedq bool
	)
	fmt.Scan(&nameq, &ageq, &marriedq)
	fmt.Printf("name:%s, age:%d, married:%t\n", nameq, ageq, marriedq)
	//Scanf
	var (
		namer    string
		ager     int
		marriedr bool
	)
	fmt.Scanf("name:%s age:%d married:%t\n", &namer, &ager, &marriedr)
	fmt.Printf("扫描结果 name:%s age:%d married:%t", namer, ager, marriedr)
	bufioDemo()
}

func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) //从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n')
	// text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
