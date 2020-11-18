# Go 并发

并发是编程中一个非常重要的概念，Go 语言在语言层面天生支持并发，这也是 Go 语言流行的一个很重要的原因。

## 并发和并行

并发：同一时间段内执行多个任务（你在用微信和两个女朋友聊天）

并行：同一时刻执行多个任务（你和你朋友都在用微信和你的女朋友聊天）

Go 语言的并发通过`goroutine`实现，`goroutine`类似于线程，属于用户态的线程，我们可以根据需要创建成千上万个`goroutine`并发工作。`goroutine`是由 Go 语言的运行时（runtime）调度完成，而线程是由操作系统调度完成。

Go 语言还提供`channel`（通道）在多个`goroutine`之间进行通信。`goroutine`和`channel`是 Go 语言秉承 CSP（communicating sequential process）并发模式的重要实现基础。

## goroutine

在 java/C++中，我们要实现并发编程的时候，我们通常需要自己去维护一个线程池，并且需要自己包装一个又一个的任务，同时需要自己去调度线程执行任务并维护上下文切换，这一切通常会耗费程序员大量的心智。

那么能不能实现一种机制，程序员只需要定义很多个人任务，让系统去帮助我们把这些任务分配到 CPU 上实现并发执行呢？

Go 语言中的`goroutine`就是这样一种机制，`goroutine`的概念类似于线程，但是`goroutine`是由 Go 的运行时（runtime）调度和管理的。Go 程序会智能地把 goroutine 中的任务合理的分配给每个 CPU。

Go 语言之所以被称为现代化的编程语言，就是因为它在语言层面已经内置了调度和上下文切换的机制。

在 Go 语言编程中，你不再需要自己去写进程、线程、协程，你的技能包里只有一个技能`goroutine`，当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，并开启一个`goroutine`

去执行这个函数即可。

### 使用 goroutine

Go 语言中使用`goroutine`非常简单，只需要在调用函数的时候在前面加上`go`关键字，就可以为一个函数创建一个`goroutine`。

一个`goroutine`必定对应一个函数，可以创建多个`goroutine`去执行相同的函数。

### 启动单个 goroutine

启动`goroutine`的方式非常简单，只需要在调用的函数（普通函数和匿名函数）的前面加上一个`go`关键字即可。

代码示例如下：

```go
func hello() {
    fmt.Println("Hello Goroutine!")
}

func main() {
    hello()
    fmt.Println("main goroutine done!")
}
```

上面这段代码中的`hello()`函数和`fmt.Println()`语句是串行的，执行的结构是打印完`Hello Goroutine!`再打印`main goroutine done!`。

接下来我们在调用函数`hello()`的时候在它前面加上关键字`go`，也就是启动一个`goroutine`去执行`hello()`这个函数。

```go
func main() {
    go hello()
    fmt.Println("main goroutine done!")
}
```

这一次的执行结果只打印了`main goroutine done!`，并没有打印`Hello Goroutine!`。为什么呢？

在程序启动时，Go 程序会为`main()`函数创建一个默认的`goroutine`，`main()`函数会在这个`goroutine`中执行，当`main()`函数返回的时候这个`goroutine`就结束了，所有在`main()`函数中启动的`goroutine`会一同结束。

`main`函数所在的`goroutine`就像是权利的游戏中的夜王，其他的`goroutine`都是异鬼，夜王一死它转化的那些异鬼也就全部 GG 了。

所以我们得想办法让 main 函数等一等 hello 函数，最简单粗暴的方式就是`time.Sleep`了。

```go
func main() {
    go hello()
    fmt.Println("main goroutine done!")
    time.Sleep(time.Second)//粗暴的等待方式
}
```

执行上面的代码你会发现，这一次先打印`main goroutine done!`，然后紧接着打印`Hello Goroutine!`。

为什么会先打印`main goroutine done!`呢？

是因为我们在创建新的 goroutine 的时候需要花费一些时间，而此时 main 函数所在的`goroutine`是继续执行的。

### 启动多个 goroutine

在 Go 语言中实现并发就是这样简单，我们还可以启动多个`goroutine`，让我们再来看一个例子：（这里使用了`sync.WaitGroup`来实现比较优雅的`goroutine`同步）

```go
var wg sync.WaitGroup

func hello(i int){
    defer wg.Done()//goutine结束就登记-1
    fmt.Println("Hello Goroutine",i)
}

func main(){
    for i:=0;i<10;i++{
        wg.Add(1)//启动一个goroutine就等级+1
        go hello(i)
    }
    wg.Wait()//优雅等待所有等级的goroutine结束
}
```

多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为 10 个`goroutine`是并发执行的，而`goroutine`的调度是随机的。

## goroutine 与线程

### 可增长的栈

OS 线程（操作系统线程）一般都有固定的栈内存（通常为 2MB），一个`goroutine`的栈在其生命周期开始时只有很小的栈（典型情况下为 2KB），`goroutine`的栈不是固定的，它可以按需增大或减小，`goroutine`的栈大小限制可以达到 1GB，虽然极少会用到这么大。所以在 Go 语言中一次创建十万几十万左右的`goroutine`也是可以的。

### goroutine 的调度

`GPM`是 Go 语言运行时（runtime）层面的实现，是 go 语言自己实现的一套调度系统，区别于操作系统调度 OS 线程：

- `G`很好理解，就是一个`goroutine`，里面除了存放本`goroutine`信息外，还存放有与所在 P 的绑定等信息。
- `P(Proc进程)`管理着一组 goroutine 队列，P 里面会存储当前 goroutine 运行的上下文环境（函数指针、堆栈地址及地址边界），P 对自己管理的 goroutine 队列做一些调度（比如把占用 CPU 时间较长的 goroutine 暂停、运行后续的 goroutine 等等），当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了就会去其他 P 的队列里抢任务。
- `M(Machine)`是 Go 运行时（runtime）对操作系统内核线程的虚拟，M 与内核线程一般是一个映射的关系，一个 goroutine 最终是要放到 M 上执行的。

P 与 M 一般也是一一对应的，它们的关系是：P 管理着一组 G 挂载在 M 上运行，当一个 G 长久阻塞在一个 M 上时，runtime 会新建一个 M，阻塞 G 所在的 P 会把其他的 G 挂载在新建的 M 上，当旧的 G 阻塞完成或者认为已经死掉时，回收旧的 M。

P 的个数是通过`runtime.GOMAXPROCS`设定的（最大 256），Go1.5 版本之后默认为物理线程数，在并发量大的时候会增加 I 写 P 和 M，但不会太多，切换太频繁的话得不偿失。

单从线程调度将，Go 语言相比其他语言的优势在于 OS 线程是由 OS 内核来调度的，`goroutine`则是由 Go 运行时（runtime）自己的调度器调度的，这个调度器使用一个**_称为`m:n`调度的技术（复用/调度 m 个 goroutine 到 n 个 OS 线程）_**。其一大特点是 goroutine 的调度是在用户态下完成的，不涉及内核态与用户态之间的频繁切换，包括内存的分配与释放，也是在用户态维护着的一块大的内存池，不直接调用系统的 malloc 函数（除非内存池需要改变），成本比调度 OS 线程低很多。另一方面 goroutine 充分利用了多核的硬件资源，近似的把若干 goroutine 均分在物理线程上，再加上 goroutine 的超轻量特性，如此种种保证了 Go 调度方面的性能。

### GOMAXPROCS

Go 运行时的调度器使用`GOMAXPROCS`参数来确定需要使用多少个 OS 线程来同时执行 Go 代码。默认值是机器上的 CPU 核心数，例如在一个 8 核的机器上，调度器会把 Go 代码同时调度到 8 个 OS 线程上（GOMAXPROCS 是 m:n 调度中的 n）。

Go 语言中可以通过`runtime.GOMAXPROCS()`函数设置当前程序并发时占用的 CPU 逻辑核心数。

Go1.5 版本之前，默认使用的是单核心执行。Go1.5 版本之后，默认使用全部的 CPU 逻辑核心数。

我们可以通过将任务分配到不同的 CPU 逻辑核心上实现并行的效果，这里举个例子：

```go
func a() {
    for i := 1; i < 10; i++{
        fmt.Println("A:", i)
    }
}

func b() {
    for i := 1; i < 10; i++{
        fmt.Println("B:", i)
    }
}

func main() {
    runtime.GOMAXPROCSZ(1)
    go a()
    go b()
    time.Sleep(time.Second)
}
```

两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。 将逻辑核心数设为 2，此时两个任务并行执行，代码如下：

```go
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
```

**Go 语言中的操作系统线程和 goroutine 的关系**：

1、一个操作系统线程对应用户态多个 goroutine

2、go 程序可以同时使用多个操作系统线程

3、goroutine 和 OS 线程是多对多的关系，即 m:n。
