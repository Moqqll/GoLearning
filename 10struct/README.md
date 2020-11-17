# Struct

Go 语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了，Go 语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体，英文名称`struct`。 也就是我们可以通过`struct`来定义自己的类型了。

Go 语言中通过`struct`来实现面向对象。

## 结构体的定义

### 普通结构体

使用`type`和`struct`关键字来定义结构体，具体代码格式如下：

```go
type 类型名 struct {
    字段名 字段类型
    字段名	字段类型
    ...
}
```

其中：

- 类型名：标识自定义结构体的名称，在同一个包内不能重名。
- 字段名：表示结构体字段的名称，结构体中的字段名必须唯一。
- 字段类型：表示结构体中字段的类型。

举个例子，我们定义一个`Person`（人）结构体，代码如下：

```go
type person struct {
	name string
	city string
	age  int8
}
```

同样类型的字段也可以写在一行，**_Go 内存对齐_**

```go
type person1 struct {
	name, city string
	age        int8
}
```

这样我们就拥有了一个`person`的自定义类型，它有`name`、`city`、`age`三个字段，分别表示姓名、城市和年龄。这样我们使用这个`person`结构体就能够很方便的在程序中表示和存储人信息了。

语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型。

## 结构体实例化

### 基本实例化

只有**_当结构体实例化时，才会真正地分配内存_**。也就是必须实例化后才能使用结构体的字段。

结构体本身也是一种类型，我们可以像声明内置类型一样使用`var`关键字来声明结构体类型。

```go
var 结构体实例 结构体类型
```

代码示例如下：

我们通过`.`来访问结构体的字段（成员变量）,例如`p1.name`和`p1.age`等。

```go
type person struct {
    name city string
    age 	  int8
}

func main() {
    var p1 person
    //var p1 = person         //错误写法
    //var p1 person{}		  //错误写法
    p1.name = "沙河娜扎"
    p1.city = "北京"
    p1.age = 18
    fmt.Printf("p1=%v\n",p1)//p1={沙河娜扎 北京 18}
    fmt.Printf("p1=%#v\n",p1)//p1=main.person{name:"沙河娜扎",city:"北京",age:18}
}
```

### 匿名结构体

有些时候仅需要直接使用一次结构体，可采取匿名结构体定义的形式

```go
var user struct{
    name     string
    married  bool
}
user.name = "小王子"
user.age = false
fmt.Printf("%#v\n",user)//main.user{name:"小王子",married:false}
```

### 创建指针类型结构体

我们可以通过`new`关键字对结构体进行实例化，得到的是结构体的地址，即指针变量。

格式如下：

```go
var p2 = new(person)
fmt.Printf("%T\n", p2)
fmt.Printf("%v\n", p2)
fmt.Printf("%#v\n", p2)
//*main.person
//&{  0}
//&main.person{name:"",city:"",age:0}
```

从打印的结果中可以看到`p2`是一个结构体指针。

**注意**：在 Go 语言中支持对结构体指针直接使用`.`来访问结构体的成员。

```go
var p2 = new(person)
p2.name = "小王子"
p2.city = "北京"
p2.age = 18
fmt.Printf("p2=%#v",p2)
//p2=main.person{name:"小王子",city:"北京",age:18}
```

### 通过取结构体的地址进行实例化

使用`&`（取地址符）对结构体进行取地址操作相当于对该结构体进行了一次`new`实例化操作，得到的也是结构体的地址，即指针变量。

代码示例如下：

```go
p4 := &person{}
// p4 := &person              //错误写法
fmt.Printf("%T\n", p4)     //*main.person
fmt.Printf("%v\n", p4)     //&{  0}
fmt.Printf("p3=%#v\n", p4) //p3=&main.person{name:"",city:"",age:0}
p4.name = "七米"
p4.age = 30
p4.city = "成都"
fmt.Printf("p3=%#v\n", p4) //p3=&main.person{name:"七米",city:"成都",age:30}
```

`p3.name = "七米"`，在底层是`(*p3).name = "七米"`，这是 Go 语言实现的语法糖。

## 结构体的初始化

实例化但没有初始化的结构体，其字段的值均为对应类型的零值。

```go
type person struct {
	name string
	city string
	age  int8
}

func main() {
	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
}
```

### 使用键值对初始化

使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值。代码示例如下：

```go
p5 := person {//实例化并初始化p5
    name: "moqqll",
    city: "上海",
    age: 18,
}
fmt.Printf("p5=%#v\n",p5)
//p5=main.person{name:"moqqll",city:"上海",age:18}
```

也可以对结构体指针采用键值对进行初始化，代码示例如下：

```go
p6 = &person {
    name: "moqqll",
    city: "上海",
    age: 18,
}
fmt.Printf("p6=%#v\n")
//p6=&main.person{name:"moqqll",city:"上海",age:18}
```

当某些字段没有初始值的时候，该字段可以不写。此时，没有指定初始值的字段的值就是该字段类型的零值。

```go
p7 := &person {
	city: "北京",
}
fmt.Printf("p7=%#v\n", p7)
//p7=&main.person{name:"", city:"北京", age:0}
```

### 使用值的列表初始化

初始化结构体的时候可以简写，也就是初始化的时候可以不写键，直接写值：

```go
p8 := &person {
    "moqqll",
    "shanghai",
    18,
}
fmt.Printf("p8=%#v\n",p8)
//p8=&main.person{name:"moqqll",city:"shanghai",age:18}
```

使用这种格式初始化时，需要注意：

1、必须初始化结构体的所有字段。

2、初始值的填充顺序必须与字段在结构体中的声明顺序一致。

3、该方式不能和键值对初始化方式混用。

## 结构体内存布局

结构体占用一块连续的内存。

```go
type test struct {
    a int8
    b int8
    c int8
    d int8
}

n := test{
    1
    2
    3
    4
}
fmt.Printf("n.a %p\n",&n.a)
fmt.Printf("n.a %p\n",&n.b)
fmt.Printf("n.a %p\n",&n.c)
fmt.Printf("n.a %p\n",&n.d)
// n.a 0xc0000a0160
// n.a 0xc0000a0161
// n.a 0xc0000a0162
// n.a 0xc0000a0163
```

【进阶知识点】关于 Go 语言中的内存对齐推荐阅读:[在 Go 中恰到好处的内存对齐](https://segmentfault.com/a/1190000017527311?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com)

### 空结构体

空结构体是不占用空间的。

```go
var v struct{}
fmt.Printf("%T\n", v)         //struct {}
fmt.Printf("%#v\n", v)        //struct {}{}
fmt.Println(unsafe.Sizeof(v)) //0
```

### 面试题

```go
type student struct {
    name string
    age int
}

func main() {
    m := make(map[string]*student, 10)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "沙河娜扎", age: 20},
		{name: "张雅婷小王八", age: 1900},
	}

	for _, stu := range stus {
		temp := stu
		// m[stu.name] = &temp
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
		fmt.Println((&stu).name)
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
		fmt.Println(k, "=>", v.name)
		//小王子=>张雅婷小王八
		//沙河娜扎=>张雅婷小王八
		//张雅婷小王八=>张雅婷小王八
	}
}
```

## 构造函数

Go 语言中的结构体没有构造函数，但可以自己实现。下方示例代码实现了一个`person`结构体的构造函数。因为`struct`是值类型，如果结构体比较复杂的话，值拷贝性能开销就会比较大，所以该构造函数的返回值是结构体指针（即，返回值的类型是结构体指针类型）。

```go
func newPerson(name,city stringa, age int8) *person {
    return &person{
        name:name,
        city:city,
        age:age,
    }
}
```

调用构造函数：

```go
p9 := newPerson("张三", "沙河", 90)
fmt.Printf("%T, %#v\n", p9, p9)
//*main.person, &main.person{name:"张三", city:"沙河", age:90}
```

## 方法和接收者

Go 语言中的`方法Method`是作用于特定类型变量的函数。这种特定类型变量叫做`接收者Receiver`，接收者的概念类似于其他语言中的`this`和`self`。

方法的定义格式如下：

```go
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

其中：

- 接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，而不是 self、this 之类的命名。例如，`Person`l 类型的接收者变量应该命名为`p`，`Connector`类型的接收者变量应该命名为`c`。
- 接收者类型：接受者类型和参数类型类似，可以是指针类型和非指针类型。
- 方法名、参数列表、返回参数的具体格式与函数定义相同。

代码示例如下：

```go
type Person struct {
    name string
    age int8
}

func NewPerson(name string, age int8) *Person {
    return &Person{
        name:name,
        age:age,
    }
}

//Dream Person做梦的方法
func (p Person) Dream() {
    fmt.Printf("%s的梦想是学好Go语言！", p.name)
}

//调用Dream方法
func main() {
    p1 := NewPerson("张三"，18)
    p1.Dream()
    //张三的梦想是学好Go语言！
}
```

### 指针类型（结构体指针）的接收者

指针类型的接收者由一个结构体的指针组成，由于指针的特性，调用 Method 方法时，会修改接收者指针所指向的结构体的任意成员变量，在 Method 方法调用结束后，修改仍旧有效。这种方法十分接近于其他语言中面向对象特性中的`this`或者`self`。

例如，我们为`Person`添加一个`SetAge`方法，用来修改实例变量的年龄，代码示例如下：

```go
//SetAge 设置p的年龄
//使用指针类型的接收者
func (p *Person) SetAge(newAge int8) {
    p.age = newAge
}

//调用该方法
func main() {
    p2 := NewPerson("小王子", 25)
    fmt.Printf("%d", p2.age)//25
    p2.SetAge(18)
    fmt.Printf("%#v\n",p2)
    //&main.Person{name:"小王子", age:18}
}
```

### 值（结构体）类型的接收者

当方法作用于值类型的接收者时，Go 语言会在代码运行时将接收者的值复制一份，在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身的成员变量的值。

代码示例如下：

```go
//SetAge2 使用值类型接收者
func (p Person) SetAge2(newAge int8) {
    p.age = newAge
}

func main() {
    p3 := NewPerson("moqqll", 18)
    fmt.Println(p3.age)//18
    p3.SetAge2(20)
    fmt.Println(p3.age)//18
}
```

### 什么时候应该使用指针类型的接收者

1、需要修改接收者中的值。

2、接收者拷贝代价比较大的对象。

3、保证一致性，如果有某个方法使用了指针类型的接收者，那么其他方法也应该使用指针类型的接收者。

## 任意类型添加方法

在 Go 语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。

举例，基于内置`int`类型使用关键字`type`可以定义新的自定义类型，然后为我们的自定义类型添加方法。

代码示例如下：

```go
//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt类型添加一个SayHello方法
func (m MyInt) SayHello() {
    fmt.Println("Hello, 我是一个int.")
}

func main() {
    var m1 MyInt
    m1.SayHello()//Hello, 我是一个int.
    m1 = 100
    fmt.Printf("%#v %T\n", m1, m1)
    //100 main.MyInt
}
```

**注意**：非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。

## 结构体的匿名字段

结构体允许其成员字段在声明时没有字段名只有字段类型，这种没有名字的字段就成为匿名字段。

```go
type Person struct {
    string
    int
}

func main() {
    p1 := Person{
        "moqqll",
        18,
    }
    fmt.Printf("%#v\n", p1)//main.Person{"moqqll", 18}
    fmt.Println(p1.string, p1.int)//moqqll 18
}
```

**注意** ：匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名必须唯一，因此一个结构体中同类型的匿名字段只能有一个。

## 嵌套结构体

一个结构体中可以嵌套包含另外一个结构体或结构体指针。

代码示例如下：

```go
//Address 地址结构体
type Address struct {
    Province string
    City 	 string
}

//User 用户结构体
type User struct {
    Name    string
    Gender  string
    Address Address
}

func main() {
    user1 := User{
        Name:"moqqll",
        Gender:"男",
        Address:Address{
            Province:"浙江",
            City:"杭州",
        },
    }
    fmt.Printf("user1=%#v\n",user1)
    //user1=main.User{Name:"moqqll", Gender:"男", Address:main.Address{Province:"浙江", City:"杭州"}}
}
```

### 嵌套匿名字段

上面 User 结构体中嵌套的 Address 结构体也可以采用匿名字段的方式。

代码示例如下：

```go
//Address 地址结构体
type Address struct {
    Province string
    City 	 string
}

//User 用户结构体
type User struct {
    Name    string
    Gender  string
    Address//匿名字段
}


func main() {
    user1 := User{
        Name:"moqqll",
        Gender:"男",
        Address:Address{
            Province:"浙江",
            City:"杭州",
        },
    }
    fmt.Printf("user1=%#v\n",user1)
    //user1=main.User{Name:"moqqll", Gender:"男", Address:main.Address{Province:"浙江", City:"杭州"}}
    fmt.Println(user1.Address.Province)//间接访问
    fmt.Println(user.Province)//直接访问，当访问结构体成员时会现在结构体中查找该字段，如果找不到再去嵌套的匿名字段中查找。
    //两种访问方式都可以
}
```

### 嵌套结构体的字段名冲突

嵌套结构体的内部可能存在相同的字段名，在这种情况下为了避免歧义，需要通过指定具体的内嵌结构体的字段名来访问内嵌结构体的成员变量。

代码示例如下：

```go
//Address 地址结构体
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱结构体
type Email struct {
	Account    string
	CreateTime string
}

//User 用户结构体
type User struct {
	Name   string
	Gender string
	Address
	Email
}

func main() {
	var user3 User
	user3.Name = "沙河娜扎"
	user3.Gender = "男"
	// user3.CreateTime = "2019" //ambiguous selector user3.CreateTime
	user3.Address.CreateTime = "2000" //指定Address结构体中的CreateTime
	user3.Email.CreateTime = "2000"   //指定Email结构体中的CreateTime
}
```

## 结构体的“继承”

Go 语言中的结构体也可以实现其他编程语言中的面向对象的继承。

代码示例如下：

```go
//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体的指针来实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
```

## 结构体字段的可见性

结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

## 结构体与 JSON 序列化

JSON 是一种轻量级的数据交换格式，易于人的阅读和编写，也易于机器的解析与生成。JSON 键值对是用来保存 JS 对象的一种方式，键/值对组合中的键名写在前面并用双引号`""`包裹，使用冒号`:`分隔，然后紧接着写值；多个键值对之间使用英文逗号`,`分隔。

代码示例如下：

```go
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

func mian() {
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
err = json.Unmarshal([]byte(str), c2)//c2必须是个指针
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
}
```

## 结构体标签(tag)

`tag`是结构体的元信息，可以在运行的时候通过反射的机制读取出来，`tag`在结构体!字段的后方定位，由一**_反引号_**包裹起来，定义格式如下：

```go
`key1:"value1" key2:"value2"`
```

结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号包裹起来。键值对之间使用一个空格分隔。

**注意**：为结构体编写`tag`时，必须严格遵守键值对的规则。结构体标签的解析代码的容差能力很差，一单格式写错，编译和运行不会提示任何错误，但通过反射无法正确取值。例：不要再 key 和 value 之间添加空格。

代码示例：为`Student`结构体的每个字段定义 json 序列化时使用 Tag：

```go
//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}
```

## 结构体和方法补充知识点

因为 slice 和 map 这两种数据类型都包含了指向底层数据的指针，因此我们在需要复制它们的时候要特别注意。

代码示例如下：

```go
type Person struct {
    name string
    age int8
    dreams []string
}

func (p *Person) SetDreams(dreams []string) {
    p.dreams = dreams
}

func main() {
    p1 := &Person{
        name:"moqqll",
        age:18,
    }
    data := []string{"吃饭", "睡觉", "打豆豆"}
    p1.SetDreams(data)

    // 你真的想要修改 p1.dreams 吗？
	data[1] = "不睡觉"
	fmt.Println(p1.dreams)  // ?
 }
```

正确的做法是在方法中使用传入的 slice 的拷贝进行结构体赋值。

```go
func (p *Person) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}
```

同样的问题也存在于返回值 slice 和 map 的情况，在实际编码过程中一定要注意这个问题。
