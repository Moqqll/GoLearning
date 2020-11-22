# time

时间和日期是我们编程中经常会用到的，本文主要介绍Go语言内置的time标准库的基本用法。

time标准库提供了时间的显示、测量等函数，其日历的计算采用的是公历。

## 时间类型

`time.Time`类型表示时间对象，我们可以通过`time.Now()`函数获取当前的时间对象，然后获取时间对象的年月日分秒等信息。代码示例如下：

```go
func timeDemo() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%-2d\n", year, month, day, hour, minute, second)
}
```

```
current time:2020-11-22 00:26:06.8178371 +0800 CST m=+0.004986101
2020-11-22 00:26:6
```

## 时间戳

时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数，它也被称为Unix时间戳（UnixTime stamp）。

基于时间对象获取时间戳的示例代码如下：

```go
func timestampDemo() int64 {
	now := time.Now()
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	return timestamp1
}

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%-2d\n", year, month, day, hour, minute, second)
}
```

```
current timestamp1:1605976840
current timestamp2:1605976840219329400
2020-11-22 00:40:40 +0800 CST                               git
2020-11-22 00:40:40
```

## 时间间隔

`time.Duration`是`time`标准库定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位。`time.Duration`表示一段时间间隔，可表示的最长时间段大约为290年。

time标准库定义的时间间隔类型的常量如下：

```go
const (
    Nanosecond Duration = 1
    Microsecond = 1000 * Nanosecond
    Millisecond = 1000 * Microsecond
    Second = 1000 * Millisecond
    Minute = 60 * Second
    Hour = 60 * Minute
)
```

例如，`time.Duration`表示1纳秒，`time.Second`表示1秒。

## 时间操作

### Add

我们在日常的编码过程中可能会遇到要求时间+时间间隔的需求，Go语言的时间对象有提供Add方法如下：

```go
func (t Time) Add(d Duration) Time	
```

代码示例如下：

```go
func mian() {
	now := time.Now()
	late := now.Add(time.Hour)
	fmt.Println(late)
}
```

### Sub

求两个时间之间的差值：

```go
func (t Time) Sub(u Time) Duration
```

返回一个时间段t-u，如果结果超出了Duration可以表示的最大值/最小值，就爱那个返回最大值/最小值。要获取时间点t-d（d 为Duration），可以使用t.Add(-d)

### Equal

```go
func (t Time) Equal(u Time) bool
```

判断两个事件是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用 t==u不同，此方面还会比较地点和时区信息。

### Before

```go
func (t Time) Before(u Time) bool
```

如果t代表的时间点在u之前，返回真，否则返回假。

### After

```go
func (t Time) After(U Time) bool
```

如果t代表的时间点在u之后，返回真，否则返回假。

## 定时器

使用`time.Tick(时间间隔)`来设置定时器，定时器本质上是一个通道（channel）

```go 
func tickDemo(){
	ticker := time.Tick(time.Second)
	for i := range ticker {
		fmt.Println(i)//每秒都会执行的任务
	}
}
```

## 时间格式化

时间类型有一个自带的方法`Format`进行格式化，需要注意的是Go语言中格式化时间模板不是常见的`Y-m-d H:M:S`，而是使用Go的诞生时间2006年1月2号 15点04分（记忆口诀：2006 1234）。也学这就是技术人员的浪漫吧。

补充，如果想格式化为12小时方式，需要指定`PM`。

代码示例如下：

```go
func formatDemo() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}
```

## 解析字符串格式的时间

```

```



