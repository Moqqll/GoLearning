package main

import (
	"fmt"
	"strings"
)

//这是不是闭包，只是函数作为返回值
func a() func() {
	return func() {
		fmt.Println("沙河小王子")
	}
}

//闭包=函数+外层变量的引用
//此时a1的返回值就是一个闭包
func a1() func() {
	name := "沙河娜扎"
	return func() {
		fmt.Println("hello,", name)
	}
}

func a2(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

//使用闭包写一个文件后缀名检测的函数
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func intSum(x ...int) (sum int) {
	fmt.Println(x)
	for _, v := range x {
		sum += v
	}
	return sum
}

func intSum1(x int, y ...int) int {
	fmt.Println(x, y)
	// 100 []
	// 100 [10]
	// 100 [10 20]
	// 100 [10 20 30]
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("panic in func B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	ret1 := intSum()
	ret2 := intSum(10)
	ret3 := intSum(10, 20)
	ret4 := intSum(10, 30)
	fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 40
	fmt.Println()

	ret5 := intSum1(100)
	ret6 := intSum1(100, 10)
	ret7 := intSum1(100, 10, 20)
	ret8 := intSum1(100, 10, 20, 30)
	fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160

	//defer语句
	//先被defer的语句后执行，后被defer的语句先执行
	fmt.Println("start")
	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	fmt.Println("end")
	fmt.Println()
	// fmt.Println("start")
	// defer fmt.Println("end")
	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")

	//匿名函数和闭包
	r := a()
	r()
	r1 := a1()
	//闭包=函数+外层变量的引用,a1()中匿名函数包含了一个匿名函数外层的变量name，此时a1的返回值就是一个闭包
	//因此r1就是一个闭包
	r1()
	r3 := a2("沙河小王子")
	r3()

	//用闭包实现检测文件后缀的函数
	test := makeSuffixFunc(".txt")
	ret := test("moqqll.txt")
	fmt.Println(ret)

	//依次调用funcA，funcB，funcC
	funcA()
	funcB()
	funcC()

	//defer经典案例
	fmt.Println("defer经典案例")
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5

	//defer面试题
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
