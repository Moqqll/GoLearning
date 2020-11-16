# typealiasAndcustomtype

## 自定义类型

在Go语言中有一些基本的数据类型，如`string`，`int`，`float`，`bool`等数据类型，我们可以使用`type`关键字来定义自定义类型。

自定义类型是定义了一个全新的类型。

我们可以基于内置的基本类型来自定义类型，也可以通过struct来定义自定义类型。

```go
//将MyInt定义为int类型
type MyInt int
```

通过`type`关键字，定义`MyInt`为一种新的类型，它具有`int`的特性。

## 类型别名

类型别名是`Go1.9`版本添加的新功能。

类型别名规定：

typealias只是Type的别名，本质上typealias与Type是同一个类型。就如一个孩子小时候有小名、乳名，作者有笔名，出国后有英文名等，这些名字都指向的是其本人。

类型别名定义格式如下：

```go
type typealias = Type
```

Go语言中自带的`rune`和`byte`都是类型别名，它们的定义如下；

```go
type byte = int8
type rune = int32
```

## 自定义类型和类型别名的异同

自定类型和类型别名表面上看只有一个等号的差异，但其本质区别我们可以通过下面这段代码来理解。

```go
//类型定义
type NewInt int

//类型别名
type MyInt = int

func main() {
	var a NewInt
	var b MyInt
	
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}
```

结果显示a的类型是`main.NewInt`，表示main包下定义的`NewInt`类型。b的类型是`int`，也就是说`MyInt`类型只会在代码中存在，编译完成时并不会有`MyInt`类型。

