package main

import (
	"fmt"
	"unsafe"
)

type person struct {
	name, city string
	age        int8
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

func main() {
	var p person
	//var p = person         //错误写法
	fmt.Printf("%T\n", p)  //main.person
	fmt.Printf("%v\n", p)  //{  0}
	fmt.Printf("%#v\n", p) //main.person{name:"", city:"", age:0}

	fmt.Println()

	// var pp person{}//错误写法
	// fmt.Printf("%#v\n", pp)
	// //main.person{name:"", city:"", age:0}
	// fmt.Println()

	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("%v\n", p2)
	fmt.Printf("%#v\n", p2)
	fmt.Println()
	// *main.person
	// &{  0}
	// &main.person{name:"", city:"", age:0}

	//p3 := person//错误写法
	p3 := person{}
	fmt.Printf("%T\n", p3)  //main.person
	fmt.Printf("%v\n", p3)  //{  0}
	fmt.Printf("%#v\n", p3) //main.person{name:"", city:"", age:0}
	fmt.Println()

	p4 := &person{}
	// p4 := &person              //错误写法
	fmt.Printf("%T\n", p4)     //*main.person
	fmt.Printf("%v\n", p4)     //&{  0}
	fmt.Printf("p3=%#v\n", p4) //p3=&main.person{name:"",city:"",age:0}
	p4.name = "七米"
	p4.age = 30
	p4.city = "成都"
	fmt.Printf("p3=%#v\n", p4) //p3=&main.person{name:"七米",city:"成都",age:30}

	tst := test{
		1,
		2,
		3,
		4,
	}
	fmt.Printf("n.a %p\n", &tst.a)
	fmt.Printf("n.a %p\n", &tst.b)
	fmt.Printf("n.a %p\n", &tst.c)
	fmt.Printf("n.a %p\n", &tst.d)
	fmt.Println()
	// n.a 0xc0000a0160
	// n.a 0xc0000a0161
	// n.a 0xc0000a0162
	// n.a 0xc0000a0163

	//空结构体
	var v struct{}
	fmt.Printf("%T\n", v)         //struct {}
	fmt.Printf("%#v\n", v)        //struct {}{}
	fmt.Println(unsafe.Sizeof(v)) //0

}
