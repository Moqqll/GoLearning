# 标识（zhi）符与关键字

## 标识符

在编程语言中标识符就是程序员定义的具有特殊意义的词，比如变量名、常量名、函数名等。Go 语言中标识符由字母、数字以及 _ （下划线）组成，并且只能以字母或 _ 开头。

## 关键字

关键字是指编程语言中预先定义好的具有特殊含义的标识符。关键字和保留字都不建议用作变量名

### Go 语言中有 25 个关键字

```go
break default func interface select case defer go map struct

chan else goto package switch const fallthrough if range type

continue for import return var
```

### Go 语言中有 37 个保留字

#### constants

```go
true false iota nil
```

#### Types

```go
int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr

float32 float64 complex64 bool byte rune string error
```

#### Functions

```go
make len cap new append copy close delete complex real imag

panic recover
```

## 变量

### 变量的来历

程序运行过程中的数据都是保存在内存中，我们想要在代码中操作某个数据时就需要去内存上找到这个变量，但是如果我们直接在代码中通过内存地址去操作变量的话，代码的可读性就会非常差而且还容易出错，所以我们就**_利用变量将这个数据的内存地址保存起来_**，以后可以直接通过这个变量就能找到内存上对应的数据了。

### 变量类型

变量（variable）的功能是存储数据，不同的变量保存的数据类型可能会不一样。经过半个多世纪的发展，编程语言已经基本形成了一套固定的类型，常见的变量数据类型有：整型、浮点型、布尔型等。

### 变量声明

Go 语言中的变量需要声明后才能使用，同一**_作用域_**内不支持重复声明，并且 Go 语言的变量声明后必须使用，否则编译时会报错。

#### 标准声明

var 变量名 变量类型

变量声明以关键字 `var`开头，变量类型放在变量的后面，行尾无需分号。举例如下：

```go
var name string

var age int

var isOK bool
```

#### 批量声明

每声明一个变量就需要写`var`关键字比较繁琐，因此 go 语言中支持变量的批量声明

```go
var(
    a string
    b int
    c bool
    d float32
)
```

#### 变量的初始化

Go 语言中在声明变量的时候，**_会自动对变量对应的内存区域进行初始化操作，每个变量会被初始化成器类型的默认值_**，例如：整型和浮点型变量的默认值是`0`，字符串变量的默认值是`空字符串`，布尔类型变量的默认值是`false`，切片、函数、指针变量的默认值是`nil`。

当然，我们也可以在声明变量的时候为其指定初始值。变量初始化的标准格式如下：

```go
var 变量名 变量类型 = 表达式
```

举例如下：

```go
var name string = "Moqqll"
var age int = 18
```

或者一次初始化多个变量：

```go
var name age = "Moqqll",18
```

#### 类型推导

有时候我们会将变量的类型省略，此时编译器会根据等号右边的值来推导变量的类型从而完成变量初始化

```go
var name = "Moqqll"
var age = 18
```

#### 短变量声明

在**_函数内部_**，可以使用简略的`:=`方式声明变量

```go
package main

import "fmt"

var m = 100 //全局变量
var y = "moqqll"

func main(){
    n := 20
    m := 200 //覆盖全局变量m的值
    y = "moke" //覆盖全局变量y的值
    fmt.Println(m,n)
}
```

#### 匿名变量

在使用多重赋值时，如果想要忽略掉某个值，可以使用`匿名变量（annoymous variable`。匿名变量用一个下划线`_`表示，例如：

```go
package main

import "fmt"

func foo()(int string){
    return 10, "moqqll"
}

func main(){
    x,_ := foo()
    _,y := foo()
    fmt.Println("x=",x)
    fmt.Println("y=",y)
}
```

匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。

**注意事项**

1、函数外的每个语句都必须以关键字开始（var，const，func 等）

2、:= 不能用在函数外

3、\_ 多用于占位，表示忽略掉值

## 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的值。常量的声明和变量的声明类似，只是将`var`换成了`const`。

**注意事项**

常量在声明时必须赋值

```go
const pi = 3.1415
const e = 2.7182
```

```go
const (
	pi = 3.1415
	e = 2.7182
)
```

const声明多个变量时，如果省略了值则表示和上面一行的值相同，例如：

```go
const (
    n1 = 100
    n2
    n3
)
```

上述实例中，常量`n1`，`n2`，`n3`的值都是1000。

### iota

`iota`是go语言中的常量计数器，只能在常量的表达式中使用

`iota`在const关键字出现时被重置为0。const中每新增一行常量声明`iota`计数一次（iota可以理解为const语句块中的行索引），使用iota能简化定义，在定义枚举时很有用。

```go
const(
    n1 = iota //n1=0
    n2 //n2=1
    n3 //n3=2
    n4 //n4=3
)
```

#### 几个常见iota实例

使用 _  跳过某些值

```go
const(
    n1 = iota //n1=0
    n2 //n2=1
    _
    n4 //n4=3
)
```

iota声明中间插队

```go
const(
    n1 = iota //n1=0
    n2 //n2=1
    n3 = 100//n3=2
    n4 =iota//n4=3
)
const n5 =iota//0
```

iota定义数量级

```go
const(
    _ = iota
    KB = 1 << (10*iota)
    MB = 1 << (10*iota)
    GB = 1 << (10*iota)
    TB = 1 << (10*iota)
    PB = 1 << (10*iota)
)
```

```go
const(
    a,b=iota+1,iota+2
    c,d
    e,f
)
```



