# Point

区别于 C/C++中的指针，Go 语言中的指针不能进行偏移和运算，是安全指针。

要搞明白 Go 语言中的指针，需要先了解 3 个概念：指针地址、指针类型、指针取值。

任何程序数据载入内存后，在内存中都会有它们的地址，这就是指针，而为了保存一个数据在内存中的地址，我们就需要指针变量。

Go 语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据时使用指针变量，而无须拷贝数据。类型指针不能进行偏移和运算，Go 语言中的指针操作非常简单，只需记住两个符号即可：`&`（取地址）和`*`（根据地址取值）。

## 地址、指针类型和指针变量

每个变量在运行时都拥有一个地址，这个地址表示变量在内存中的位置。Go 语言中使用`&`（取地址符）放在变量前面对变量进行“取地址”操作。Go 语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型（`*int`、`*int64`、`*string`等）。

取变量指针的语法格式如下：

```go
ptr := &v
```

其中：

- v：代表被取地址的变量，类型为 T。
- ptr：代表接收地址的变量，类型为 `*T`，称为 T 的指针类型。\*代表指针。

示例如下：

```go
a := 10
b := &a
fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
fmt.Println(&b)                    //0xc00000e018
```

图示：

![取变量地址图示](https://www.liwenzhou.com/images/Go/pointer/ptr.png)

## 指针取值

在对普通变量使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用\*操作，也就是指针取值，代码如下。

```go
//指针取值
a := 10
b := &a // 取变量a的地址，将指针保存到b中
fmt.Printf("type of b:%T\n", b)
c := *b // 指针取值（根据指针去内存取值）
fmt.Printf("type of c:%T\n", c)
fmt.Printf("value of c:%v\n", c)

//输出结果
//type of b:*int
//type of c:int
//value of c:10
```

**总结**： 取地址操作符`&`和取值操作符`*`是一对互补操作符，`&`取出地址，`*`根据地址取出地址指向的值。

变量、地址、指针变量、取地址、取值的关系如下：

- 对变量进行取地址（&）操作，可以拿到这个变量的地址，并存入指针类型的变量（即指针变量）中。
- 指针变量的值是原变量的地址。
- 对指针变量进行取值（\*）操作，可以获得这个指针变量所指向的原变量的值。

```go
func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a) // 100
}
```

## new 和 make

看一个代码示例：

```go
func main() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
```

执行上述代码会引发 panic，为什么呢？

在 Go 语言中对于**引用类型**的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。而对于**值类型**的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配了内存空间。

要分配内存，就引出来`new`和`make`，Go 语言中的 new 和 make 是内建的两个函数，主要用来分配内存。

### new

new 是一个内建的函数，它的函数签名如下：

```go
func new(Type) *Type
```

其中：

- Type 表示类型，new 函数只接受一个参数，这个参数是一个类型。
- \*Type 表示类型指针，new 函数返回一个指向该类型内存地址的指针。

new 函数不太常用，使用 new 函数得到的是一个类型指针，并且该指针对应的值为该类型的零值。示例如下：

```go
func main(){
    a := new(int)
    b := new(bool)
    fmt.Printf("%T\n",a)// *int
    fmt.Printf("%T\n",b)// *bool
    fmt.Println(*a)		// 0
    fmt.Println(*b)		//false
}
```

示例中的代码`var a *int`只是声明了一个指针变量 a，未初始化。指针作为引用类型需要初始化后才会拥有内存空间，才可以给指针指向变量赋值。应该按照如下方式使用内置的 new 函数对 a 进行初始化后才可以对其正常赋值。

```go
func main() {
	var a *int
	a = new(int)
	*a = 10
	fmt.Println(*a) //10
}
```

### make

make 函数也是用于分配内存的，区别于 new 函数，它只用于 slice、map 以及 chan 的内存创建，并且它返回的类型就是这三个类型本身，而不是它们的指针。因为这三个类型就是也能用类型，因此没有必要返回它们的指针了。

make 函数的函数签名如下：

```go
func make(t Type, size ...IntegerType) Type
```

make 函数是无可替代的，我们在使用 slice、map 和 chan 的时候，都需要使用 make 函数进行初始化（即分配内存），然后才可以对它们进行操作。

本节开始的示例中`var b map[string]int`只是声明变量 b 是一个 map 类型的变量，需要像下面的代码一样使用 make 函数对 b 初始化后，才能对其进行赋值操作。

```go
func main(){
    var b map[string]int
    b = make(map[string]int, 10)
    b["沙河娜扎"] = 100
    fmt.Println(b)//map[沙河娜扎:100]
}
```

### new 和 make 的异同

1、二是都是用来做内存分配的。

2、make 只用于 slice、map 和 chan 的初始化，返回的也是这三个引用类型自身。

3、new 用于值类型的内存分配，并且初始化后的变量的值为类型零值，返回的是指向该类型的指针。
