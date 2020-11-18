# Concurrent and Lock

有时候在Go代码中可能会存在多个`goroutine`同时操作一个资源（临界区）的情况，这种情况会发生`竞态问题`（数据竞态）。类比现实生活中的例子有十字路口被各个方向的汽车竞争，火车上的卫生间被车厢里的人竞争等。

代码示例如下：

```go
var x int64
var wg sync.WaitGroup

func add(){
    for i:=0;i<500;i++{
        x+=i
    }
    wg.Done()
}

func main() {
    wg.Add(2)
    go add()
    go add()
    wg.Wait()
    fmt.Println(x)
}
```

上面的代码中我们开启了两个`goroutine`去累加变量`x`的值，这两个`goroutine`在访问和修改`x`变量的时候就会存在数据竞争，导致最后的结果与期望不符。

## 互斥锁

互斥锁是一种常用的用于控制共享资源访问的方法，它能保证同时只有一个`goroutine`可以访问共享的资源。Go语言中使用`sync`包的`Mutex`类型来实现互斥锁。

使用互斥锁修复上面的代码，示例如下：

```go
var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add(){
    for i:=0;i<500;i++{
        lock.Lock()//加锁
        x+=i
        lock.Unlock()//解锁
    }
    wg.Done()
}

func main() {
    wg.Add(2)
    go add()
    go add()
    wg.Wait()
    fmt.Println(x)
}
```

使用互斥锁能够保证同一时间有且只有一个`goroutine`进入临界区，其他的`goroutine`则在等待；当互斥所释放后，等待的`goroutine`才可以进入临界区，多个`goroutine`同时等待一个锁时，唤醒的策略是随机地。

## 读写互斥锁

互斥锁是完全互斥的，但是有很多实际场景下是 读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用`sync`包中的`RwMutex`类型。

读写锁分两种：读锁和写锁。当一个`goroutine`获取读锁之后，其他的`goroutine`如果获取读锁则会继续获取锁，如果获取写锁就会等待；当一个`goroutine`获取写锁之后，其他的`goroutine`无论是获取读锁还是写锁都会等待。

读写锁代码示例如下：

```go
var {
    x int64
    wg sync.WaitGroup
    lock sync.Mutex
    rwlock sync.RWMutex
}

func write() {
    rwlock.Lock()//加写锁
    x = x + 1
    time.Sleep(10*time.Millisecond)//假设读操作耗时10ms
    rwlock.Unlock()//解写锁
    wg.Done()
}

func read() {
    rwlock.RLock()//加读锁
    time.Sleep(time.Millisecond)//假设读操作耗时1ms
    rwlock.RUnlock()//解读锁
    wg.Done()
}

func main() {
    start := time.Now()
    for i:=0;i<10;i++ {
        wg.Add(1)
        go write()
    }
    for i:=0;i<1000;i++{
        wg.Add(1)
        go read()
    }
    wg.Wait()
    end := time.Now()
    fmt.Println(end.Sub(start))
}
```

**注意**：读写锁非常适合“读多写少”的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。

## sync.WaitGroup

在代码中生硬的使用`time.Sleep`肯定是不合适的，Go语言中可以使用`sync.WaitGroup`来实现并发任务的同步。`sync.WaitGroup`有以下几个方法：

| 方法名                         | 功能                |
| ------------------------------ | ------------------- |
| (wg *WaitGroup) Add(delta int) | 计数器+delta        |
| (wg *WaitGroup) Done()         | 计数器-1            |
| (wg *WaitGroup) Wait()         | 阻塞知道计数器变为0 |

`sync.WaitGroup`内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了N个并发任务时，就将计数器增加N，每个任务完成时通过调用Done()方法将计数器减1，通过调用Wait()来等待并发任务执行完，当计数器值为0时，表示所有并发任务已经完成。

我们利用`sync.WaitGroup`将上面的代码优化一下：

```go
var wg sync.WaitGroup

func hello() {
    defer wg.Done()
    fmt.Println("Hello Goroutine!")
}

func main() {
    wg.Add(1)
    go hello()
    fmt.Println("main goroutine done!")
    wg.Wait()
}
```

**注意**：`sync.WaitGroup`是一个结构体，传递的时候要传递指针。

## sync.Once

***这是一个进阶知识点***

在编程的很多场景下，我们需要确保某些操作在高并发场景下只执行一次，例如只加载一次配置文件，只关闭一次通道等。

Go语言中的`sync`包提供了一个针对***只执行一次场景***的解决方案---`sync.Once`。`sync.Once`只有一个`Do`方法，其签名如下：

```go
func (o *Once) Do(f func()) {}
```

**备注**：如果要执行的函数`f`需要传递参数就需要搭配闭包来使用。

### 加载配置文件示例

延迟***一个开销很大的初始化操作到真正用到它的时候再执行***是一个很好的实践。因为预先初始化一个变量（比如在init函数中完成初始化）会增加程序的启动耗时，而且有可能实际执行过程中这个变量没有用上，那么这个初始化操作就不是必须要做的。

代码示例如下：

```go
var icons map[string]image.Image

func loadIcons(){
    icons = map[string]image.Image{
        "left":loadIcon("left.png"),
        "up":loadIcon("up.png"),
        "right":loadIcon("right.png"),
        "down":loadIcon("down.png"),
    }
}

//Icon被多个goroutine调用时不是并发安全的
func Icon(name string) image.Image{
    if icons == nil{
        loadIcons()
    }
    return icons[name]
}
```

多个`goroutine`并发调用`Icon`函数时并不是并发安全的，现代的编译器和CPU可能会保证每个`goroutine`都满足串行一致的基础上自由地重排访问内存的顺序。loadIcons函数可能会被重排为以下结果：

```go
func loadIcons() {
    icons = make(map[string]image.Image)
    icons["left"] = loadIcon("left.png")
    icons["up"] = loadIcon("up.png")
    icons["right"] = loadIcon("right.png")
    icons["down"] = loadIcon("down.png")
}
```

在这种情况下就会出现***即使判断了`icons`不是nil，也不意味着初始化完成了***。考虑到这种情况，我们能想到的办法就是添加互斥锁，保证初始化`icons`的时候不会被其他的`goroutine`操作，但是这样做又会引发性能问题。

使用`sync.Once`改造的代码示例如下：

```go

```











