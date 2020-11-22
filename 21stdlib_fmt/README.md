# fmt

fmt标准库是我们学习Go语言过程中接触最早最频繁的一个标准库，本文介绍fmt标准库的一些常用函数。

fmt标准库实现了类似C语言printf和scanf的格式化I/O，主要分为向外输出内容和获取输入内容两大部分。

## 向外输出

标准库`fmt`提供了以下几种输出相关的函数。

### Print

`Print`系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。

```go
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```

​	举个简单的例子：

```go
func main() {
	fmt.Print("在终端打印该信息。")
	name := "沙河小王子"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")
}
```

执行上面的代码输出：

```
在终端打印该信息。我是：沙河小王子
在终端打印单独一行显示
```

### Fprint

Fprint系列函数会将内容输出到一个io.writer接口类型的变量w中，我们通常使用这个函数往文件中写入内容。

```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
```

举个例子：

```go
//Fprint
fmt.Fprintln(os.Stdout, "向标准输出写入内容")
fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil {
	fmt.Println("打开文件错误, err：", err)
	return
}
name := "沙河小王子"
//向打开的文件句柄中写入内容main
fmt.Fprintf(fileObj, "往文件中写入信息：%s", name)

```

**注意**：只要满足`io.Writer`接口的类型都支持写入

### Sprint

Sprint系列函数会把传入的数据生成并返回一个字符串。

```go
func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
```

举例如下：

```go
s1 := fmt.Sprint("moqqll")
name1 := "沙河娜扎"
age := 18
s2 := fmt.Sprintf("name:%s,age:%d", name1, age)
s3 := fmt.Sprintln("沙河小王子")
fmt.Println(s1, s2, s3)
```

### Errorf

`Errorf`函数根据format参数生成格式化字符串并返回一个包含该字符串的错误

```go
func Errorf(format string, a ...interface(){}) error
```

通常使用这种方式来自定义错误类型，例如：

```go
err := fmt.Errorf("这是一个错误")
```

Go1.13版本为`fmt.Errorf`函数新加了一个`%w`占位符用来生成一个包裹Error的Wrapping Error。

```go
e := errors.New("原始错误err")
w := fmt.Errorf("Wrap了一个错误%w", err)
```

## 格式化占位符

*printf系列函数都支持format格式化参数，在这里我们按照占位符将被替换的变量类型划分，方便查询和记忆。

### 通用占位符

| 占位符 | 说明                               |
| ------ | ---------------------------------- |
| %v     | 值的默认格式                       |
| %+v    | 类似%v，但输出结构体时会添加字段名 |
| %#v    | 值的Go语法表示                     |
| %T     | 值的类型                           |
| %%     | 百分号                             |

代码示例如下：

```go
//格式化占位符
fmt.Printf("%v\n", 100)
fmt.Printf("%v\n", false)
o := struct{ name string }{"moqqll"}
fmt.Printf("%v\n", o)
fmt.Printf("%+v\n", o)
fmt.Printf("%#v\n", o)
fmt.Printf("%T\n", o)
fmt.Printf("100%%\n")
```

### 布尔型

| 占位符 | 说明                        |
| ------ | --------------------------- |
| %t     | 只能接收bool型：true或false |

### 整型

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| %b     | 二进制                                                       |
| %c     | Unicode码值                                                  |
| %d     | 十进制                                                       |
| %o     | 八进制                                                       |
| %x     | 十六进制，使用a-f                                            |
| %X     | 十六进制，使用A-F                                            |
| %U     | Unicode格式                                                  |
| %q     | 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示 |

代码示例如下：

```go
n := 999999
fmt.Printf("二进制：%b\n", n)
fmt.Printf("Unicode码值：%c\n", n)
fmt.Printf("十进制：%d\n", n)
fmt.Printf("八进制：%o\n", n)
fmt.Printf("十六进制：%x\n", n)
fmt.Printf("十六进制：%X\n", n)
fmt.Printf("Unicode：%U\n", n)
```

输出结果如下：

```
二进制：11110100001000111111
Unicode码值：󴈿
十进制：999999
八进制：3641077
十六进制：f423f
十六进制：F423F
Unicode：U+F423
```

### 浮点数与复数

| 占位符 | 说明                                                 |
| ------ | ---------------------------------------------------- |
| %b     | 无小数部分，二进制指数的科学记数法                   |
| %e     | 科学计数法                                           |
| %E     | 科学计数法                                           |
| %f     | 有小数部分但无指数部分                               |
| %F     | 等价于%f                                             |
| %g     | 根据实际情况采用%e或%f格式，以获得更简洁、准确地输出 |
| %G     | 根据实际情况采用%E或%F格式，以获得更简洁、准确地输出 |

### 字符串和[]byte

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| %s     | 直接输出字符串或者[]byte                                     |
| %q     | 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示 |
| %x     | 每个字节用两字符十六进制数表示（使用a~f）                    |
| %X     | 每个字节用两字符十六进制                                     |

### 指针

| 占位符 | 说明                           |
| ------ | ------------------------------ |
| %p     | 表示为十六进制，并加上前导的0x |

### 宽度标识符

宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不做填充。精度通过（可选的）宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0。

| 占位符 | 说明               |
| ------ | ------------------ |
| %f     | 默认宽度，默认精度 |
| %9f    | 宽度9，默认精度    |
| %.2f   | 默认宽度，精度2    |
| %9.2f  | 宽度9，精度2       |
| %9.f   | 宽度9，精度0       |

代码示例如下：

```go
n := 12.34
fmt.Printf("%f\n", n)
fmt.Printf("%9f\n", n)
fmt.Printf("%.2f\n", n)
fmt.Printf("%9.2f\n", n)
fmt.Printf("%9.f\n", n)
```

输出结果如下：

```
12.340000
12.340000
12.34
    12.34
       12
```

### 其他flag

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| ‘+‘    | 总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义） |
| ’ ‘    | 对数值，正数前加空格，负数前加负号；对字符串采用%x活%X时，会给各打印的字节之间加空格 |
| ’_'    | 在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐） |
| '#'    | 八进制数前加0(%#o)，十六进制数前加0x(%#x)，指针去掉前面的0x(%#p)，对%#U会输出空格和单引号括起来的go字面值 |
| '0'    | 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面   |

## 获取输入

Go语言`fmt`标准库下有fmt.Scan，fmt.Scanf，fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入。

### fmt.Scan

函数的签名如下：

```go
func Scan(a ...interface{}) (n int, err error)
```

* Scan从标准输入扫描文本，读取由空白符分隔的值保存到传递给本函数的参数中，换行符视为空白符
* 本函数返回成功扫描的数据个数和遇到的任何错误，如果读取的数据个数比提供的参数少，会返回一个错误报告原因。

代码实力如下：

```go
//Scan
var (
	name    string
	age     int
	married bool
)
fmt.Scan(&name, &age, &married)
fmt.Printf("name:%s, age:%d, married:%t\n", name, age, married)
```

将上面的代码编译后在终端执行，在终端依次输入`小王子`、`28`和`false`使用空格分隔。

```
小王子 28 false
扫描结果 name:小王子 age:28 married:false
```

`fmt.Scan`从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。

### fmt.Scanf

函数签名如下：

```go
funC Scanf(format string, a ...interface{}) (n int, err error)
```

* Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中
* 本函数返回成功扫描的数据个数和遇到的任何错误

代码示例如下：

```go
var (
	name    string
	age     int
	married bool
)
fmt.Scanf("name:%s age:%d married:%t\n", &name, &age, &married)
fmt.Printf("扫描结果 name:%s age:%d married:%t", name, age, married)
```

```
name:moqqll age:18 married:false
扫描结果 name:moqqll age:18 married:false
```

`fmt.Scanf`不同于`fmt.Scan`简单的以空格作为输入数据的分隔符，`fmt.Scanf`为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。

例如，我们还是按照上个示例中以空格分隔的方式输入，`fmt.Scanf`就不能正确扫描到输入的数据。

```
moqqll 18 false
扫描结果 name: age:0 married:false
```

### fmt.Scanln

函数签名如下：

```go
func Scanln(a ...interface{}) (n int, err error)
```

- Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。
- 本函数返回成功扫描的数据个数和遇到的任何错误。

将上面的代码编译后在终端执行，在终端依次输入`小王子`、`28`和`false`使用空格分隔。

```
小王子 28 false
扫描结果 name:小王子 age:28 married:false 
```

`fmt.Scanln`遇到回车就结束扫描了，这个比较常用。

### bufio.NewReader

有时候我们想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用`bufio`标准库来实现

```go
func bufioDemo() {
	reader := bufio.NewReader(os.Stdin) //从标准输入生成读对象
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n')
	// text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
```

```
请输入内容：moqqll 18 true
"moqqll 18 true"

请输入内容：moqqll 18 true
"moqqll 18 true\r\n"
```

### Fscan系列

这几个函数功能分别类似于`fmt.Scan`，`fmt.Scanf`，`fmt.Scanln`三个函数，只不过他们不是从标准输入中读取数据，而是从`io.Reader`中读取数据。

```go
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int,err error)
func Fscanf(r io.Reader, a ...interface{}) (n int,err error)
```

### Sscan系列

这几个函数功能分别类似于`fmt.Scan`，`fmt.Scanf`，`fmt.Scanln`三个函数，只不过他们不是从标准输入中读取数据，而是从指定字符串中读取数据。

```go
func Scan(str string, a ...interface{}) (n int, err error)
func Scanln(str string, a ...interface{}) (n int, err error)
func Scanf(str string, a ...interface{}) (n int, err error)
```

