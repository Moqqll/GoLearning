package main

import "fmt"

func main() {
	//声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔类型切片并初始化
	// var d = []bool{false, true} //声明一个布尔类型切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[flase,true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
	fmt.Println()

	//基于数组得到切片
	m := [5]int{55, 56, 57, 58, 59}
	//简单切片表达式
	n := m[1:3]           //设置索引值（左闭右开），从数组m得到切片n
	fmt.Println(n)        //[56,57]
	fmt.Printf("%T\n", n) //[]int
	fmt.Println(len(n))   //2
	fmt.Println(cap(n))   //4
	//完整切片表达式
	t := m[1:3:4]
	fmt.Printf("%T,%v,%v\n", t, len(t), cap(t)) //[]int,2,3
	//切片再次切片
	q := n[0:1]
	q1 := n[3:4]
	fmt.Println(q)        //[56]
	fmt.Println(q1)       //[59]
	fmt.Printf("%T\n", q) //[]int
	fmt.Println(len(q))   //1
	fmt.Println(cap(q))   //4
	fmt.Println(cap(q1))  //1
	fmt.Println()

	//通过make函数构造切片
	p := make([]int, 5, 10) //make(切片类型,切片长度,切片容量)
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(len(p)) //获取切片长度len()=5
	fmt.Println(cap(p)) //获取切片容量cap()=10
	fmt.Println()

	//长度、容量为0的切片，和，值为nil的切片
	var s1 []int                                   //声明切片未初始化，长度、容量为0，值为nil
	s2 := []int{}                                  //声明且初始化切片，长度、容量为0，值不为nil
	s3 := make([]int, 0)                           //make()切片，长度、容量为0，值不为nil
	fmt.Printf("%v,%v,%v\n", s1, len(s1), cap(s1)) //s1==nil, 0, 0
	fmt.Println(s1 == nil)                         //true
	fmt.Printf("%v,%v,%v\n", s2, len(s2), cap(s2)) //s2!=nil, 0, 0
	fmt.Println(s2 == nil)                         //false
	fmt.Printf("%v,%v,%v\n", s3, len(s3), cap(s3)) //s3!=nil, 0, 0
	fmt.Println(s3 == nil)                         //false
	fmt.Println()

	//切片的复制/拷贝
	r1 := make([]int, 3) //[0 0 0]
	r2 := r1             //将r1直接赋值给r2，r1和r2共用一个底层数组
	r2[0] = 100
	fmt.Println(r1) //[100 0 0]
	fmt.Println(r2) //[100 0 0]
	fmt.Println()

	//切片的遍历
	sr := []int{3, 7, 8}
	for i := 0; i < len(sr); i++ {
		fmt.Println(i, sr[i])
	}
	for index, value := range sr {
		fmt.Println(index, value)
	}
	fmt.Println()

	//append()函数向切片中追加元素
	//可以一次追加一个，也可以一次追加多个，还可以追加其他切片元素
	var st []int
	st = append(st, 1)
	fmt.Println(st) //[1]
	st = append(st, 2, 3, 4)
	fmt.Println(st) //[1,2,3,4]
	sv := []int{6, 7, 8}
	st = append(st, sv...)
	fmt.Println(st) //[1,2,3,4,6,7,8]
	fmt.Println()

	//append()函数添加元素和自动扩容
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
	fmt.Println()

	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Println(slice, newSlice) //[10,20,30,60,50][20,30,60]
	fmt.Println()

	// copy()复制切片
	h1 := []int{1, 2, 3, 4, 5}
	h2 := make([]int, 5, 5)
	copy(h2, h1)    //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(h1) //[1 2 3 4 5]
	fmt.Println(h2) //[1 2 3 4 5]
	h2[0] = 1000
	fmt.Println(h1) //[1 2 3 4 5]
	fmt.Println(h2) //[1000 2 3 4 5]
	fmt.Println()

	// 从切片中删除元素
	at := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	at = append(at[:2], at[3:]...)
	fmt.Println(at)                                          //[30 31 33 34 35 36 37]
	fmt.Printf("len(at):%d, cap(at):%d\n", len(at), cap(at)) //7, 8
	fmt.Println()

	var aa = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		aa = append(aa, fmt.Sprintf("%v", i))
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", aa, len(aa), cap(aa), aa)
	}
	fmt.Println(aa)
	fmt.Printf("len:%d, cap:%d\n", len(aa), cap(aa))
}
