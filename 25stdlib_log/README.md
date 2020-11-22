# log

无论是软件开发的调试阶段还是软件上线后的运行阶段，日志一直都是一个非常重要的环节，我们也应该养成在程序中实现日志记录的好习惯。

Go 语言内置的`log`包实现了简单的日志服务，本文介绍了`log`标准库的基本使用。

## 使用 Logger

log 标准库定义了 Logger 类型，该类型提供了一些格式化输出的方法。该标准库也提供了一个预定义的“标准 logger“，可以通过调用函数`Print`系列（Print、Printf、Println）、`Fatal`系列（Fatal、Fatalf、Fatalln）和`Panic`系列（Panic、Panicf、Panicln），比自行创建一个 logger 对象更容易

例如，我们可以向下面的代码一样直接通过`log`包来调用上面提到的方法，默认他们会将日志信息打印到终端界面：

```go
func lg() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志\n", v)
	log.Printf("这是一条%s狗\n", v)
	log.Fatalln("这是一条会触发Fatal的日志")
	//log.Panicln("这是一条会触发Panic的日志")
}
```

```
2020/11/22 02:22:24 这是一条很普通的日志。
2020/11/22 02:22:24 这是一条很普通的日志
2020/11/22 02:22:24 这是一条很普通的狗
2020/11/22 02:22:24 这是一条会触发Fatal的日志
exit status 1
//2020/11/22 02:25:06 这是一条会触发Panic的日志
//panic: 这是一条会触发Panic的日志

//goroutine 1 [running]:
//log.Panicln(0xc00007df58, 0x1, 0x1)
        //D:/Go/src/log/log.go:365 +0xb3
//main.lg()
        //F:/8DevOps/github.com/Moqqll/02goLearning/25stdlib_log/main.go:13 +0x180
//main.main()
        //F:/8DevOps/github.com/Moqqll/02goLearning/25stdlib_log/main.go:17 +0x27
//exit status 2
```

logger 会打印每条日志信息的日期、时间，默认输出到系统的标准错误。Fatal 系列函数会在写入日志信息后调用 os.Exit(1)，Panic 系列函数会在写入日志信息后 panic。

## 配置 logger

### 标准 lgger 的配置

默认情况下的 logger 只会提供日志的时间信息等，但是很多情况下我们希望得到更多新，比如记录该日志的文件名和行号等。`log`标准库中为我们提供了定制这些设置的方法。

`log`标准库中的`Flags`函数会返回标准的 logger 输出配置，而`SetFlags`函数用来设置标准 logger 的输出配置。

### flag 选项

`log`标准库提供了如下的 flag 选项，它们是一系列定义好的常量。

```
const (
	//控制输出日志信息的细节，不能控制输出的顺序和格式
	//输出的日志在每一项后会又一个冒号分隔，例如：2009/
	Ldate	= 1 << iota
	Ltime
	Lmicroseconds
	Llongfile
	Lshortfile
	LUTC
	LstdFlags	= Ldate | Ltime
)
```

下面，我们在记录日志之前先设置一下标准 logger 的输出选项：

```
log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
log.Println("这是一条很普通的日志。")
```

```
2020/11/22 02:47:23.418395 F:/8DevOps/github.com/Moqqll/02goLearning/25stdlib_log/main.go:21: 这是一条很普通的日志。
```

### 配置日志前缀

`log`标准库中还提供了关于日志信息前缀的两个方法：

```go
func Prefix() string
func SetPrefix(prefix string)
```

其中`Prefix`函数用来查看标准 logger 的输出前缀，`SetPrefix`函数用来设置输出前缀。

```
log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.L date)
log.Println("这是一条很普通的日志。")
// fmt.Println(log.Prefix())
log.SetPrefix("[moqqll]")
log.Println("这是一条很普通的日志。")
```

```
2020/11/22 02:50:36.508239 main.go:21: 这是一条很普通的日志
。
[moqqll]2020/11/22 02:50:36.527241 main.go:24: 这是一条很普
通的日志。
```

这样我们就能够在代码中为我们的日志信息添加指定的前缀，方便之后对日志信息进行检索和处理。

### 配置日志输出位置

```go
func SetOutput(w io.Writer)
```

`SetOutput`函数用来设置标准 logger 的输出目的地，默认是标准错误输出。

例如，下面的代码会把日志输出到同目录下的 xx.log 文件中：

```go
logFile, err := os.OpenFile("xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil {
	fmt.Println("Open log file failed, err:", err)
	return
}
log.SetOutput(logFile)
log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
log.Println("这是一条很普通的日志。")
log.SetPrefix("[moqqll]")
log.Println("这是一条很普通的日志")
```

如果你要使用标准的 logger，我们通常会把上面的配置操作写到`init()`函数中：

```go
func init() {
	logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
```

## 创建 logger

`log`标准库中还提供了一个创建新 logger 对象的构造函数–`New`，支持我们创建自己的 logger 示例。`New`函数的签名如下：

```go
func New(out io.Writer, prefix string, flag int) *logger
```

New 创建一个 Logger 对象。其中，参数 out 设置日志信息写入的目的地。参数 prefix 会添加到生成的每一条日志前面。参数 flag 定义日志的属性（时间、文件等等）。

举个例子：

```go
logger := log.New(os.Stdout, "<NEW>", log.Lshortfile|log.Ldate|log.time)
log.Println("这是自定义的logger记录的日志")
```

将上面的代码编译执行之后，得到结果如下：

```
<NEW>2020/11/22 03:03:41 main.go:33: 这是自定义的logger记录
的日志
```

## 总结

Go 内置的 log 库功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如[logrus](https://github.com/sirupsen/logrus)、[zap](https://github.com/uber-go/zap)等。
