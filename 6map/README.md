# map

Go 语言中提供的映射关系容器为`map`，其内部使用`散列表（hash）`实现。

map 是一种无序的的基于`key-value`的数据结构，Go 语言中的 map 是引用类型，必须初始化才是使用。

## map 定义

Go 语言中`map`的定义语法如下：

```go
map [keyType]valueType
//keyType：键的类型
//ValueType:键对应的值的类型
```

map 类型的变量默认初始值为 nil，需要使用 make()函数来分配内存。语法为：

```go
make(map[keyType]valueType,[cap])
//cap表示map的容量
//该参数不是必须的，但是我们应该在初始化map的时候就为其指定一个合适容量
```

```django
var a map[string]int
fmt.Println(a == nil) //true
```

## map 基本使用

map 中的数据都是成对出现的，map 的基本使用示例如下：

```go
scoreMap := make(map[string]int, 8)
scoreMap["张三"] = 90
scoreMap["李四"] = 100
fmt.Println(scoreMap)
fmt.Println(scoreMap["张三"])
fmt.Printf("type of scoreMap: %T\n", scoreMap)
// map[张三:90 李四:100]
// 90
// type of scoreMap: map[string]int
```

map 也支持在声明的时候填充元素，示例如下：

```go
userInfo:=map[string]string{
	"user_name":"moqqll",
	"password":"moqqll",
}
fmt.Println(userInfo)
//map[password:moqqll user_name:moqqll]
```

## 判断某个键值对是否存在

Go语言中有个判断map中键值对是否存在的特殊写法，格式如下：

```go
value,ok := map[key]
```

示例如下：

```go
v, ok := scoreMap["李四"] //不存在时v的值为0
if ok {
	fmt.Println(v)
} else {
	fmt.Println("查无此人")
	fmt.Println(v)
}
```

## map的遍历

Go语言中使用`for range`遍历map

```go
for k, v := range scoreMap {
	fmt.Println(k, v)
}
//李四 100
//张三 90

for k := range scoreMap { //只遍历key
	fmt.Println(k)
}
for _, v := range scoreMap { //忽略key，只取value值
	fmt.Println(v)
}
```

**注意**：遍历map时的元素顺序和添加键值对的顺序无关。

## 使用delete()函数删除键值对

Go语言内建函数`delete()`可以从map中删除一组键值对，格式如下：

```go
delete(map,key)
//map：要删除键值对的map
//key：要删除的键值对的键
```

示例如下：

```go
delete(scoreMap, "张三")
for k, v := range scoreMap {
	fmt.Println(k, v) //李四 100
}
```

## 按照指定顺序遍历map









