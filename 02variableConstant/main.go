package main

import "fmt"

var m = 100
var y = "moqqll"

const (
	n1 = iota //n1=0
	n2        //n2=1
	n3        //n3=2
	n4        //n4=3
)

const (
	m1 = iota //m1=0
	m2        //m2=1
	_
	m4 //m4=3
)

const (
	a1 = iota //a1=0
	a2        //a2=1
	a3 = 100  //a3=2
	a4 = iota //a4=3
)
const a5 = iota //a5=0

const (
	_ = iota
	//KB ...
	KB = 1 << (10 * iota)
	//MB ...
	MB = 1 << (10 * iota)
	//GB ...
	GB = 1 << (10 * iota)
	//TB ...
	TB = 1 << (10 * iota)
	//PB ...
	PB = 1 << (10 * iota)
)

const (
	aa, bb = iota + 1, iota + 2
	cc, dd
	ee, ff
)

func foo() (int, string) {
	return 10, "moqqll"
}

func main() {
	//标准声明
	var name = "moqqll"
	var age int
	var gender bool
	fmt.Println(name, age, gender) //moqqll 0 false

	//批量声明
	var (
		a int
		b bool
		c = "moqqll"
		d float32
	)
	fmt.Println(a, b, c, d) // 0 false moqqll 0

	n := 10
	m := 200   //覆盖全局变量m的值
	y = "moke" //覆盖全局变量y的值
	fmt.Println(n, m, y)

	q, _ := foo()
	_, z := foo()
	fmt.Println("q =", q)
	fmt.Println("z =", z)

	fmt.Println(n1, n2, n3, n4)         //0 1 2 3
	fmt.Println(m1, m2, m4)             //0 1 3
	fmt.Println(a1, a2, a3, a4, a5)     //0 1 100 3 0
	fmt.Println(KB, MB, GB, TB, PB)     //1024 ...
	fmt.Println(aa, bb, cc, dd, ee, ff) //1,2,2,3,3,4
}
