package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

type person struct {
	name, city string
	age        int8
	dreams     []string
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}

type student struct { //结构体字段名不能小写否则其他包拿不到数据
	ID   int
	Name string
	Age  int
}

func newStudent(id, age int, name string) *student {
	return &student{
		ID:   id,
		Age:  age,
		Name: name,
	}
}

type class struct { //结构体字段名不能小写否则其他包拿不到数据
	Title    string
	Students []*student
}

func (p *person) SetDreams(dreams []string) {
	p.dreams = dreams //这会导致p.dreams指向的底层数组和dreams指向的底层数组相同，从而想着相互影响
	// p.dreams = make([]string, len(dreams))//为p.dreams开辟一块新内存
	// copy(p.dreams, dreams)//把dreams指向的底层数组的值，拷贝一份给p.dreams指向的底层数组。
}

func main() {
	var p person
	//var p = person         //错误写法
	fmt.Printf("%T\n", p)  //main.person
	fmt.Printf("%v\n", p)  //{  0}
	fmt.Printf("%#v\n", p) //main.person{name:"", city:"", age:0}
	fmt.Printf("%v\n", &p) //&{  0}
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
	fmt.Println()

	m := make(map[string]*student, 10)
	stus := []student{
		{Name: "小王子", Age: 18},
		{Name: "沙河娜扎", Age: 20},
		{Name: "张雅婷小王八", Age: 1900},
	}

	for _, stu := range stus {
		temp := stu
		m[stu.Name] = &temp
		// m[temp.name] = &temp//这种写法也可以
		// 最后如预期的结果输出如下：
		// 小王子 => 小王子
		// 沙河娜扎 =>沙河娜扎
		// 张雅婷小王八 => 张雅婷小王八
		fmt.Println(&stu)
		// &{小王子 18}
		// &{沙河娜扎 20}
		// &{张雅婷小王八 1900}
		fmt.Printf("%p\n", &stu)
		// 0xc0000044c0
		// 0xc0000044c0
		// 0xc0000044c0
		fmt.Println(*(&stu))
		//{小王子 18}
		//{沙河娜扎 20}
		//{张雅婷小王八 1900}
		fmt.Println((&stu).Name)
		//小王子
		//沙河娜扎
		//张雅婷小王八
	}
	fmt.Println()
	fmt.Println(m)
	//map[小王子:0xc0000044c0 张雅婷小王八:0xc0000044c0 沙河娜扎:0xc0000044c0]
	fmt.Printf("%p\n", m)
	//0xc00006e600
	for k, v := range m {
		fmt.Println(v)
		//&{张雅婷小王八 1900}
		//&{张雅婷小王八 1900}
		//&{张雅婷小王八 1900}
		fmt.Printf("%#v\n", *v)
		// main.student{name:"张雅婷小王八", age:1900}
		// main.student{name:"张雅婷小王八", age:1900}
		// main.student{name:"张雅婷小王八", age:1900}
		fmt.Printf("%p\n", &k)
		//0xc00003c250
		//0xc00003c250
		//0xc00003c250
		fmt.Printf("%p\n", v)
		// 0xc0000044c0
		// 0xc0000044c0
		// 0xc0000044c0
		fmt.Printf("%p\n", &v)
		//0xc000006030
		//0xc000006030
		//0xc000006030
		fmt.Println(*(&v))
		//&{张雅婷小王八 1900}
		//&{张雅婷小王八 1900}
		//&{张雅婷小王八 1900}
		fmt.Println(k, "=>", v.Name)
		//小王子=>张雅婷小王八
		//沙河娜扎=>张雅婷小王八
		//张雅婷小王八=>张雅婷小王八
	}
	fmt.Println()

	c1 := &class{
		Title:    "火箭101",
		Students: make([]*student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		tmpstu := newStudent(i, 18, fmt.Sprintf("stu%02d", i))
		fmt.Printf("%v\n", tmpstu)
		// &{0 stu00 18}
		// &{1 stu01 18}
		// &{2 stu02 18}
		// &{3 stu03 18}
		// &{4 stu04 18}
		// &{5 stu05 18}
		// &{6 stu06 18}
		// &{7 stu07 18}
		// &{8 stu08 18}
		// &{9 stu09 18}
		c1.Students = append(c1.Students, tmpstu)
	}
	fmt.Printf("%v\n", c1)
	//&{火箭101 [0xc000004740 0xc000004780 0xc0000047c0 0xc000004800 0xc000004840 0xc000004880 0xc0000048c0 0xc000004900 0xc000004940 0xc000004980]}
	fmt.Println()

	//JSON序列化：结构体数据 --> JSON格式数据
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed.")
		return
	}
	fmt.Printf("%T\n", data) //[]uint8
	fmt.Printf("%s\n", data) //结构体字段名不能小写否则其他包拿不到数据
	//{"Title":"火箭101","Students":[{"ID":0,"Name":"stu00","Age":18},{"ID":1,"Name":"stu01","Age":18},{"ID":2,"Name":"stu02","Age":18},{"ID":3,"Name":"stu03","Age":18},{"ID":4,"Name":"stu04","Age":18},{"ID":5,"Name":"stu05","Age":18},{"ID":6,"Name":"stu06","Age":18},{"ID":7,"Name":"stu07","Age":18},{"ID":8,"Name":"stu08","Age":18},{"ID":9,"Name":"stu09","Age":18}]}

	//JSON反序列化：JSON格式数据 --> 结构体数据
	str := `{"Title":"火箭101","Students":[{"ID":0,"Name":"stu00","Age":18},{"ID":1,"Name":"stu01","Age":18},{"ID":2,"Name":"stu02","Age":18},{"ID":3,"Name":"stu03","Age":18},{"ID":4,"Name":"stu04","Age":18},{"ID":5,"Name":"stu05","Age":18},{"ID":6,"Name":"stu06","Age":18},{"ID":7,"Name":"stu07","Age":18},{"ID":8,"Name":"stu08","Age":18},{"ID":9,"Name":"stu09","Age":18}]}`

	c2 := &class{}
	err = json.Unmarshal([]byte(str), c2) //c2必须是个指针
	if err != nil {
		fmt.Println("json unmarshal failed.")
		return
	}
	fmt.Printf("%v\n", c2)
	//&{火箭101 [0xc000004c20 0xc000004c60 0xc000004c80 0xc000004ca0 0xc000004ce0 0xc000004d00 0xc000004d40 0xc000004d60 0xc000004d80 0xc000004dc0]}
	data2, err1 := json.Marshal(c1)
	if err1 != nil {
		fmt.Println("json marshal failed.")
		return
	}
	fmt.Printf("%T\n", data2) //[]uint8
	fmt.Printf("%s\n", data2)
	//{"Title":"火箭101","Students":[{"ID":0,"Name":"stu00","Age":18},{"ID":1,"Name":"stu01","Age":18},{"ID":2,"Name":"stu02","Age":18},{"ID":3,"Name":"stu03","Age":18},{"ID":4,"Name":"stu04","Age":18},{"ID":5,"Name":"stu05","Age":18},{"ID":6,"Name":"stu06","Age":18},{"ID":7,"Name":"stu07","Age":18},{"ID":8,"Name":"stu08","Age":18},{"ID":9,"Name":"stu09","Age":18}]}

	//
	pp := &person{
		name: "moqqll",
		age:  18,
	}
	dreamdata := []string{"吃饭", "睡觉", "打豆豆"}
	fmt.Printf("%p\n", dreamdata) //0xc000070750
	pp.SetDreams(dreamdata)
	fmt.Printf("%v\n", pp)
	fmt.Printf("%p\n", pp.dreams) //0xc000070750,这不是所期望的。

	dreamdata[1] = "不睡觉" //会把pp.dreams修改，这不是所期望的。
	fmt.Printf("%v\n", pp)
}
