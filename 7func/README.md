# Func

函数是组织好的、可重复使用的、用于执行指定任务的代码块。

Go 语言中支持**_函数、匿名函数和闭包_**，并且函数在 Go 语言中属于“一等公民”。

## 函数定义

Go 语言中定义函数使用`func`关键字，格式如下：

```go
func 函数名(参数)(返回值){
  函数体
}
```

- 函数名：由字母、数字、下划线组成，但函数名的第一个字符不能为数字，在同一个包内，函数名不能重名。
- 参数：参数由`参数变量`和`参数变量的类型`组成，多个参数之间使用`,`分隔
- 返回值：返回值由`返回值变量`和`返回值变量类型`组成，也可以只写返回值的类型，多个返回值必须用`()`包裹，并用`,`分隔。
- 函数体：实现特定功能的代码块。

## 函数的调用

定义了函数之后，可以通过`函数名()`方式调用函数。

**注意**：调用有返回值的函数时，可以不接收其返回值

## 参数

###　类型简写

函数的参数中如果相邻变量的类型相同，则可以简写类型，示例如下：

```go
func intSum(x,y int)int{
  return x + y
}
```

### 可变参数

可变参数是指函数的参数数量不固定，Go 语言中的可变参数通过在参数名后加`...`来标识。

**注意**：可变参数通常要作为函数的最后一个参数。

**本质上**：函数的可变参数是通过切片来实现的。

示例如下：

```go
func intSum(x ...int) (sum int) {
	fmt.Println(x)
    //[]
	//[10]
	//[10 20]
	//[10 30]
	for _, v := range x {
		sum += v
	}
	return sum
}

func main() {
	ret1 := intSum()
	ret2 := intSum(10)
	ret3 := intSum(10, 20)
	ret4 := intSum(10, 30)
	fmt.Println(ret1, ret2, ret3, ret4)
    //0 10 30 40
}
```

固定参数搭配可变参数使用时，可变参数要放在固定参数的后面，示例如下：

```go
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

func main() {
	ret5 := intSum1(100)
	ret6 := intSum1(100, 10)
	ret7 := intSum1(100, 10, 20)
	ret8 := intSum1(100, 10, 20, 30)
	fmt.Println(ret5, ret6, ret7, ret8) 
    //100 110 130 160
}
```

Go语言中的函数是没有默认参数的。

## 返回值

Go语言中通过`return`关键字向外输出返回值。

### 多返回值

Go语言中函数支持多返回值，函数如果有多个返回值时必须用`()`将所有返回值包裹起来。

示例如下：

```go
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
```

### 返回值命名

函数定义时可以给返回值命名，并在函数中直接使用这些变量，最后通过`return`关键字返回。

示例如下：

```go
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return //直接写一个return即可,但建议写上，代码会比较清晰
}
```

### 返回值补充

当一个函数的返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。

```go
func someFunc(x string) []int {
	if x == "" {
		return nil // 返回nil即可，没必要返回[]int{}
	}
	...
}
```



## 函数进阶

### 变量作用域

#### 全局变量

全局变量是定义在函数外部的变量，它是程序整个运行周期内都有效。在函数中可以访问全局变量。

```go
package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testGlobalVar() {
	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
}
func main() {
	testGlobalVar() //num=10
}
```

#### 局部变量

局部变量分两种：一是函数内定义的变量，二是代码块中定义的变量。

***函数内定义的变量***无法在函数外使用，示例如下：

```go
func testLocalVar() {
	//定义一个函数局部变量x,仅在该函数内生效
	var x int64 = 100
	fmt.Printf("x=%d\n", x)
}

func main() {
	testLocalVar()
	fmt.Println(x) // 此时无法使用变量x
}
```

如果全局变量和局部变量重名，优先访问局部变量。

```go
package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testNum() {
	num := 100
	fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量
}
func main() {
	testNum() // num=100
}
```

***语句块定义的变量***，通常我们会在if条件判断、for循环、switch语句中使用语句块定义的变量。示例如下：

```go
func testLocalVar2(x, y int) {
	fmt.Println(x, y) //函数的参数也是只在本函数中生效
	if x > 0 {
		z := 100 //变量z只在if语句块生效
		fmt.Println(z)
	}
	//fmt.Println(z)//此处无法使用变量z
}
```

以及for循环语句中定义的变量，也只在for语句块中生效：

```go
func testLocalVar3() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) //变量i只在当前for语句块中生效
	}
	//fmt.Println(i) //此处无法使用变量i
}
```

### 函数类型与变量

定义的函数名可以直接作为变量赋值给一个新变量`a`，新变量`a`的类型则为`func()`类型，在函数中则可以直接调用新变量`a()`执行函数。

#### 定义函数类型

使用`type`关键字来定义一个函数类型，格式如下：

```go
type calculation func(int, int) int
```

上面语句定义了一个`calculation`类型，它是一种函数类型，这种函数接收两个int类型参数并返回一个int类型返回值。

简单来说，凡是满足这个条件的函数都是`calculation`类型的函数，例如下面的`add`和`sub`是calculation类型。

```go
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
```

add和sub都能赋值给calculation类型的变量。

```go
var c calculation
c = add
```

#### 函数类型变量

可以声明函数类型的变量并且为该变量赋值：

```go
func main() {
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f1
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}
```

### 高阶函数

#### 函数作为参数

```go
func add(x, y int) int {
	return x + y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main() {
	ret2 := calc(10, 20, add)
	fmt.Println(ret2) //30
}
```

#### 函数作为返回值

```go
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
```

### 匿名函数和闭包

#### 匿名函数

在Go语言中，函数作为返回值时，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数定义格式如下：

```go
func(参数)(返回值){
    函数体
}
```

匿名函数没有函数名，因此无法像调用普通函数那样调用匿名函数，所以匿名函数需要保存到某个变量中，或者作为立即执行函数使用：

```go
func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
```

匿名函数多用于实现回调函数和闭包。

#### 闭包

闭包值的是一个函数和与其相关的引用环境组合而成的实体，简单来说，`闭包=函数+引用环境`。先看一个例子：

```go
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

func main(){
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
}
```

## defer语句

Go语言中的`defer`语句会将其后面跟随的语句进行延迟处理。在`defer`语句所在的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行。也就是说，先被`defer`的语句最后被执行，最后被`defer`的语句最先被执行。

```go
    fmt.Println("start")
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    fmt.Println("end")
    //start
	//end
	//3
	//2
	//1

    fmt.Println("start")
    defer fmt.Println("end")
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
	//start
	//3
	//2
	//1
	//end
```

`defer`语句延迟执行的特性非常适用于资源释放场景，比如：资源清理、文件关闭、解锁、及记录时间等。

### defer执行时机

Go语言函数中的`return`语句在底层并不是原子操作，它分为给返回值赋能和`RET`指令两步。而`defer`语句执行的时机就在返回值复制操作后、`RET`指令执行前，如下图所示：

![defer执行时机](https://www.liwenzhou.com/images/Go/func/defer.png)

### defer经典案例

阅读下面的代码，写出最后的打印结果。

```go
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
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
//5
//6
//5
//5
```

### defer面试题

问，下面的代码的输出结果是？***（提示：defer注册要延迟执行的函数时该函数`所有的参数`都需要确定其值）***

```GO
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
//A 1 2 3
//B 10 2 12
//BB 10 12 22
//AA 1 3 4
```

## 内置函数介绍

| 内置函数       | 介绍                                                         |
| -------------- | ------------------------------------------------------------ |
| close          | 主要用来关闭channel                                          |
| len            | 用来求长度，比如string、array、slice、map、channel           |
| new            | 用来分配内存，主要用来分配值类型，比如int、struct，返回的是指针 |
| make           | 用来分配内存，主要用来分配引用类型，比如chan、map、slice     |
| append         | 用来追加元素到数组、slice中                                  |
| panic和recover | 用来处理程序错误                                             |

### panic和recover

Go语言（1.12版本及以前）是没有异常机制的，但是可以使用`panic/recover`模式来处理错误。`panic`可以在任何地方已发，但是`recover`只有在`defer`定义的函数（匿名函数）中才有效。

示例如下：

```go
func funcA() {
	fmt.Println("func A")
}

func funcB() {
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}

 //依次调用funcA，funcB，funcC
//funcB会抛出panic，导致程序崩溃退出
funcA()
funcB()
funcC()
//func A
//...
```

程序运行期间`funcB`中引发了`panic`导致程序崩溃，异常退出了。这个时候我们可以利用`recover`将捕获`err`并将程序恢复，继续往后执行。

```go
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

//依次调用funcA，funcB，funcC
funcA()
funcB()
funcC()
//func A
//panic in func B
//func C
```

**注意**

* `recover()`必须搭配`defer`使用。
* `defer`一定要在可能引发`panic`的语句之前定义。























