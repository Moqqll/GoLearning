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

**注意**：在Go语言中支持对结构体指针直接使用`.`来访问结构体的成员。

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

`p3.name = "七米"`，在底层是`(*p3).name = "七米"`，这是Go语言实现的语法糖。

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

【进阶知识点】关于Go语言中的内存对齐推荐阅读:[在 Go 中恰到好处的内存对齐](https://segmentfault.com/a/1190000017527311?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com)

### 空结构体

空结构体是不占用空间的。

```go

```



















