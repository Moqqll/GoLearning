# Channel

单纯地将函数并发执行是没有意义的，函数与函数之间需要交换数据才能体现并发执行函数的意义。

虽然可以使用共享内存的方式进行数据交换，但是共享内存在不同的`goroutine`中容易发生竞态问题，为了保证数据交换的正确性与稳定性，必须使用互斥量对内存进行加锁，这种做法势必会造成性能问题。

Go 语言的并发模型是`CSP(communicating sequential process)`，提倡**_通过通信共享内存_**而不是**_通过共享内存实现通信_**。

如果说`goroutine`是 Go 程序并发的执行体，那`channel`就是他们之间的连接。`channel`是可以让一个`goroutine`发送特定值给另一个`goroutine`的通信机制

Go 语言中的通道（channel）是一种特殊的类型，通道像一个传动带或者队列，总是遵循先入先出（First In First Out）的规则，从而保证收发数据的顺序。**_每一个通道都是一个具体类型的导管，也就是声明 channel 的时候需要为其指定元素类型_**。

## channel 类型

`channel`是一种类型，一种引用类型。声明通道类型的格式如下：

```go
var 变量 chan 元素类型
```

举几个例子：

```go
var ch1 chan int//声明一个传递整型的通道
var ch2 chan bool//声明一个传递布尔类型的通道
var ch3 chan []int//声明一个传递int切片的通道
```

## 创建 channel

通道是**_引用类型_**，通道类型的空值是`nil`。

```go
var ch chan int
fmt.Println(ch)//<nil>
```

声明的通道**_需要使用`make`函数初始化后才能使用_**

格式如下：

```go
变量 := make(chan 元素类型， [缓冲区大小])
```

channel 的缓冲区大小是可选的。

举几个例子：

```go
ch1 := make(chan int, [])
ch2 := make(chan bool, [])
ch3 := make(chan []int, [])
```

## channel 操作

通常通常有发送（send），接收（receive）和关闭（close）三个操作。

发送和接受都使用符号`<-`

现在我们先使用以下语句定义一个通道：

```go
ch := make(chan int)
```

### 发送

将一个值发送到通道中。

```go
ch <- 10
```

### 接收

从通道中接收一个值

```go
x := <- ch//从通道ch中接收值并复制给变量x
<- ch //从通道中接收值并忽略
```

### 关闭

我们可以调用内置的`close()`函数关闭通道

```go
close(ch)
```

**注意**：关于关闭通道需要注意的是，只有在通知接收方 goroutine 所有的数据都发送完毕的时候才需要关闭通道；**_通道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的_**。

关闭后的通道有以下特点：

- 对一个关闭的通道再发送值会触发 panic。
- 对一个关闭的通道进行接收会一直获取值，直到通道为空。
- 对一个关闭且没有值的通道执行接收操作，会得到对应类型的零值。
- 关闭一个已经关闭的通道会触发 panic。

## 无缓冲的通道

无缓冲的通道又成为阻塞的通道。我们来看下面的代码示例：

```go
func main() {
	ch := make(chan int)
	ch <- 10
	fmt.Println("发送成功")
}
```

上面这段代码能够通过编译，但是执行的时候会出现以下错误：

```go
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        .../src/github.com/Q1mi/studygo/day06/channel02/main.go:8 +0x54
```

为什么会出现`deadlock`错误呢？

因为我们使用`ch := make(chan int)`创建的是无缓冲的通道，无缓冲的通道只有在有人接收值的时候才能发送值，简单来说就是无缓冲的通道必须有接收才能发送。

上面的代码会阻塞在`ch <- 10`这一行代码，形成死锁，那如何解决这个问题呢？

一种方式是启用一个`goroutine`去接收值，例如：

```go
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
```

无缓冲通道上的发送操作会阻塞，直到另一个`goroutine`在该通道上执行接收操作，这时值才能发送成功，两个`goroutine`将继续执行。相反，如果接收操作先执行，接收方的`goroutine`将阻塞，直到另一个`goroutine`在该通道上发送一个值。

使用无缓冲通道进行通信将导致发送和接收的`goroutine`同步化。因此，无缓冲通道也被称为`同步通道`。

## 有缓冲的通道

解决上面问题的方法还有一种就是使用有缓冲区的通道。我们可以在使用 make 函数初始化通道的时候为其指定通道的容量，例如：

```go
func main() {
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功")
}
```

只要通道的容量大于零，那么该通道就是有缓冲的通道，通道的容量表示通道中能存放元素的数量。

我们可以使用内置的`len()`函数获取通道内元素的数量，使用`cap()`函数获取通道的容量，但我们很少会这么做。

## for range 从通道循环取值

当向通道中发送完数据时，我们可以通过`close()`函数来关闭通道。

当通道被关闭时，再往该通道发送值会引发`panic`，从该通道取值的操作会先取完通道中的值，最后取到的值会一直是通道存储元素对应类型的零值。

那如何判断一个通道是否被关闭了呢？

代码示例如下：

```go
//channel练习
func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    //开启goroutine将0~100的数发送到ch1中
    go func(){
        for i:=0;i<100;i++{
            ch1<-i
        }
        close(ch1)
    }()

    //开启goroutine取出ch1的值，并求平方后发送到ch2中
    //利用for无限循环从通道中依次取值
    go func() {
        for {
            i,ok := <-ch1
            if !ok {//如果值去完了，就退出for循环
                break
             }
            ch2 <- i*i
        }
        close(ch2)
    }()

    //利用for range从通道中依次取值
    for i := range ch2{//取完值且通道关闭，自动退出for range循环
        fmt.Println(i)
    }
}
```

从上面的例子中我们可以看到有两种方式在接收通道值的时候判断该通道的值是否被取完且通道是否被关闭，通常使用的是`for range`的方式；使用`for range`遍历通道，当通道的值被取完且通道关闭了的时候就会退出`for range`。

## 单向通道

有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或只能接收。

Go 语言中提供了**_单向通道_**来处理这种情况。例如，我们把上面的例子改造如下：

```go
func couter(out chan<- int) {
    for i:=0;i<100;i++{
        out <- i
    }
    close(out)
}

func squarer(out chan<- int, in <-chan int) {
    for i := range in {
        out <- i*i
    }
    close(out)
}

func printer(in <-chan int) {
    for i := in {
        fmt.Println(i)
    }
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)

    go counter(ch1)
    go squarer(ch2, ch1)
    printer(ch2)
}
```

其中：

- `chan<-`是一个只写单向通道（只能对其写入 int 类型值），可以对其执行发送操作，但不能执行接收操作。
- `<-chan`是一个只读单向通道（只能从中读取 int 类型值），可以对其执行接收操作，但不能执行发送操作。

在函数传参及任何赋能操作中可以将双向通道转换为单向通道，但反过来是不可以的。

# 通道总结

`channel`常见的异常总结，如下图：

![channel异常总结](https://www.liwenzhou.com/images/Go/concurrence/channel01.png)

关闭已经关闭的`channel`，也会引发`panic`。

# worker pool （goroutine 池）

在工作中我们通常会使用可以指定启动的`goroutine`数量--`worker pool`模式，控制`goroutine`的数量，防止`goroutine`的泄漏和暴涨。

一个简易的`work pool`代码示例如下：

```go
func worker(id int, in <-chan int, out chan<- int) {
    for i := range in {
        fmt.Printf("worker:%d start job:%d\n", id, i)
        time.Sleep(time.Second)
        fmt.Printf("worker:%d stop job:%d\n", id, i)
        out <- i*i
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    //开启3个goroutine
    for w := 1; w <3; w++ {
        go worker(w, jobs, results)
    }

    //5个任务
    for i:= 1; i < 5; i++ {
        jobs <- i
    }
    close(jobs)

    //打印结果
    for r := 0; r<5; r++ {
        ret := <- results
        fmt.Println(ret)
    }
}
```

# select 多路复用

在某些场景下我们需要同时从多个通道接收数据，通道在接收数据时，如果没有数据可以接收将会发生阻塞。

你也许会写出如下代码，使用遍历方式实现从多个通道接收数据：

```go
for {
    //尝试从ch1接收值
    data, ok := <-ch1

    //尝试从ch2接受值
    data, ok := <-ch2

    ...
}
```

这种方式虽然可以实现从多个通道接收值的需求，但是运行性能会差很多。为了应对这种场景，Go 内置了`select`关键字，可以同时响应多个通道的操作。

`select`的使用类似于 switch 语句，它有一系列 case 分支和一个默认的分支，每个 case 分支会对应一个通道的通信（接收或发送）过程。`select`会一直等待，直到某个`case`的通信操作完成时，就会执行`case`分支对应的语句。具体格式如下：

```go
switch {
    case <-ch1:
    	...
    case data := <-ch2:
    	...
    case ch3<-data:
    	...
    default:
    	默认操作
}
```

举个小例子来演示下`select`的使用：

```go
func main() {
    ch := make(chan int, 1)
    for i:=0;i<10;i++ {
        select{
            case data:=<-ch:
            	fmt.Println(data)
            case ch<-i:
        }
    }
}
```

使用`select`语句能提高代码的可读性：

- 可处理一个或多个 channel 的发送/接收操作。
- 如果多个`case`同时满足，`select`会随机选择一个执行。
- 对于没有`case`的`select{}`会一直等待，可用于阻塞 main()函数。
