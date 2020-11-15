# Slice

切片（slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装，它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含`地址`(指向底层数组），`长度`，`容量`，切片一般用于快速地操作一块数据集合。

##　切片的声明

### 标准声明

```go
var name []T
```

其中：

- name 表示变量名
- T 表示切片中的元素类型

示例如下：

```go
func main(){
    //声明切片类型
    var a[]string //声明一个字符串切片
    var b=[]int{} //声明一个整型切片并初始化
    var c=[]bool{false,true} //声明一个布尔类型切片并初始化
    var d=[]bool{false,true} //声明一个布尔类型切片并初始化
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(a==nil) //true
    fmt.Println(b==nil) //false
    fmt.Println(c==nil) //false
    fmt.Println(c==d) //切片是引用类型，不支持直接比较，只能和nil比较
}
```

### 基于数组得到切片

切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。它有两种变体：一种**指定 low 和 high 两个索引界限值的简单的形式**，另一种是**除了 low 和 high 索引界限值外还指定容量的完整的形式**。

**_得到的切片`长度=high-low`，容量等于得到的切片的底层数组的容量。_**

```go
m := [5]int{55, 56, 57, 58, 59}
//简单切片表达式
n := m[1:3]           //设置索引值（左闭右开），从数组m得到切片n
fmt.Println(n)        //[56,57]
fmt.Printf("%T\n", n) //[]int
fmt.Println(len(n))   //2
fmt.Println(cap(n))   //4
//完整切片表达式
t := m[1:3:4]
fmt.Printf("%T,%v,%v\n", t, len(t), cap(t)) //[]int,2,3
```

### 切片再次切片

```go
q := n[0:1]
q1 := n[3:4]
fmt.Println(q)        //[56]
fmt.Println(q1)       //[59]
fmt.Printf("%T\n", q) //[]int
fmt.Println(len(q))   //1
fmt.Println(cap(q))   //4
fmt.Println(cap(q1))  //1
```

### 通过 make 函数声明切片

```go
p := make([]int, 5, 10) //make(切片类型,切片长度,切片容量)
fmt.Println(p)
fmt.Printf("%T\n", p)
fmt.Println(len(p)) //获取切片长度len()=5
fmt.Println(cap(p)) //获取切片容量cap()=10
```

### 切片的长度和容量

获取长度：`len()`

获取容量：`cap()`

## 切片的本质

![slice_01](https://www.liwenzhou.com/images/Go/slice/slice_01.png)

![slice_02](https://www.liwenzhou.com/images/Go/slice/slice_02.png)

## 判断切片是否为空

使用`len(s)==0`来判断，而不能使用`s==nil`。因为**_长度为 0 的切片不一定值为 nil_**。

## 切片不能直接比较

切片之间是不能直接比较的，不能使用`==`操作符来判断两个切片是否含有全部相等元素。切片唯一合法的比较操作是和`nil`比较。一个`nil`值的切片没有底层数组，因此该切片的长度和容量都是 0。但是一个长度和容量都是 0 的切片不一定是`nil`。示例如下：

```go
var s1 []int                    //声明切片未初始化，长度、容量为0，值为nil
s2 := []int{}                   //声明且初始化切片，长度、容量为0，值不为nil
s3 := make([]int, 0)            //make()切片，长度、容量为0，值不为nil
fmt.Printf("%v,%v,%v\n", s1, len(s1), cap(s1)) //[], 0, 0
fmt.Println(s1 == nil)                         //s1==nil,true
fmt.Printf("%v,%v,%v\n", s2, len(s2), cap(s2)) //[], 0, 0
fmt.Println(s2 == nil)                         //s2!=nil,false
fmt.Printf("%v,%v,%v\n", s3, len(s3), cap(s3)) //[], 0, 0
fmt.Println(s3 == nil)                         //s3!=nil,false
```

## 切片的复制拷贝

对切片进行复制拷贝，拷贝前后的两个切片共享底层数组，对一个切片的修改会影响另一个切片的内容，这点需要特别注意。

```go
s1 := make([]int, 3) //[0 0 0]
s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
s2[0] = 100
fmt.Println(s1) //[100 0 0]
fmt.Println(s2) //[100 0 0]
```

## 切片的遍历

切片的遍历方式和数组是一致的，支持索引遍历和`for range`遍历。

```go
s := []int{1, 3, 5}

for i := 0; i < len(s); i++ {
	fmt.Println(i, s[i])
}

for index, value := range s {
	fmt.Println(index, value)
}
```

## append()方法为切片添加元素

Go 语言的内建函数`append()`可以为切片动态添加元素，可以一次添加一个元素，也可以添加多个元素，还可以另一个切片中的元素（后面加 ...）

```go
var st []int
st = append(st, 1)//通过var声明的零值切片可以在append()函数中直接使用，无需初始化。
fmt.Println(st) //[1]
st = append(st, 2, 3, 4)
fmt.Println(st) //[1,2,3,4]
sv := []int{6, 7, 8}
st = append(st, sv...)
fmt.Println(st) //[1,2,3,4,6,7,8]
```

**注意**：通过`var`声明的零值切片可以在 append()函数中直接使用，无需初始化。

每个切片都会指向一个底层数组，这个数组的容量足够用就添加新增元素。当底层数组的容量不足以容纳新增元素时，切片会自动按照一定的策略进行“扩容”，**_此时切片指向的底层数组就会更换_**。“扩容”发生于`append()`函数调用时，所以我们通常都需要用原变量接收 appen()函数的返回值。

```go
//append()添加元素和切片扩容
var numSlice []int
for i := 0; i < 10; i++ {
	numSlice = append(numSlice, i)
	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, 		len(numSlice), cap(numSlice), numSlice)
}
```

## 使用 copy()函数复制切片

Go 语言内建的`copy()`函数可以迅速地将一个切片的数据复制到另外一个切片空间中，使用方式如下：

`copy(destSlice,srcSlice[]T)`

其中：

- srcSlice：数据来源切片
- destSlice：目标切片

```go
// copy()复制切片
a := []int{1, 2, 3, 4, 5}
c := make([]int, 5, 5)
copy(c, a)     //使用copy()函数将切片a中的元素复制到切片c
fmt.Println(a) //[1 2 3 4 5]
fmt.Println(c) //[1 2 3 4 5]
c[0] = 1000
fmt.Println(a) //[1 2 3 4 5]
fmt.Println(c) //[1000 2 3 4 5]
```

## 从切片中删除元素

Go 语言中并没有删除切片元素的专用方法，我们可以使用切片自身的特性来删除元素

```go
// 从切片中删除元素
a := []int{30, 31, 32, 33, 34, 35, 36, 37}
// 要删除索引为2的元素
a = append(a[:2], a[3:]...)
fmt.Println(a) //[30 31 33 34 35 36 37]
fmt.Printf("len(at):%d, cap(at):%d\n", len(at), cap(at)) //7, 8
```

## 写出下面代码的输出结果

```go
var aa = make([]string, 5, 10)
for i := 0; i < 10; i++ {
	aa = append(aa, fmt.Sprintf("%v", i))
	fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", aa, len(aa), cap(aa), aa)
}
fmt.Println(aa)
fmt.Printf("len:%d, cap:%d\n", len(aa), cap(aa))
// [     0]  len:6  cap:10  ptr:0xc00005a0a0
// [     0 1]  len:7  cap:10  ptr:0xc00005a0a0
// [     0 1 2]  len:8  cap:10  ptr:0xc00005a0a0
// [     0 1 2 3]  len:9  cap:10  ptr:0xc00005a0a0
// [     0 1 2 3 4]  len:10  cap:10  ptr:0xc00005a0a0
// [     0 1 2 3 4 5]  len:11  cap:20  ptr:0xc00005c140
// [     0 1 2 3 4 5 6]  len:12  cap:20  ptr:0xc00005c140
// [     0 1 2 3 4 5 6 7]  len:13  cap:20  ptr:0xc00005c140
// [     0 1 2 3 4 5 6 7 8]  len:14  cap:20  ptr:0xc00005c140
// [     0 1 2 3 4 5 6 7 8 9]  len:15  cap:20  ptr:0xc00005c140
// [     0 1 2 3 4 5 6 7 8 9]
// len:15, cap:20
```
