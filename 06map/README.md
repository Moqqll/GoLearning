# map

Go 语言中提供的映射关系容器为`map`，其内部使用`散列表（hash）`实现。

map 是一种**_无序_**的基于`key-value`的数据结构，Go 语言中的 map 是引用类型，必须初始化才是使用。

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

```go
//map默认为nil
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

Go 语言中有个判断 map 中键值对是否存在的特殊写法，格式如下：

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

## map 的遍历

Go 语言中使用`for range`遍历 map

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

**注意**：遍历 map 时的元素顺序和添加键值对的顺序无关。

## 使用 delete()函数删除键值对

Go 语言内建函数`delete()`可以从 map 中删除一组键值对，格式如下：

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

## 按照指定顺序遍历 map

```go
//按照指定顺序遍历map
//将map的key存入slice，对slice进行排序
//根据排序后的slice遍历map
rand.Seed(time.Now().UnixNano()) //初始化随机数种子
var smap = make(map[string]int, 200)
for i := 0; i < 100; i++ { //初始化smap
	temp := rand.Intn(100)
	key := fmt.Sprintf("stu%02d", temp) //生成stu开头的字符串
	//生成0~99的随机整数
	smap[key] = i * 10
}
fmt.Println(smap)
fmt.Println()
//取出map中所有的key存入slice
var keys = make([]string, 0, 200)
for key := range smap {
		keys = append(keys, key)
	}
//对slice进行排序
fmt.Println(keys)
fmt.Println()
sort.Strings(keys)
fmt.Println(keys)
fmt.Println()
//按照排序后的slice元素的顺序遍历may
for _, key := range keys {
	fmt.Println(key, smap[key])
}
// map[stu02:100 stu05:830 stu07:420 stu10:650 stu11:980 stu12:880 stu13:950 stu15:610 stu16:70 stu17:630 stu18:460 stu19:680 stu20:970 stu21:870 stu22:700 stu23:660 stu24:850 stu27:600 stu28:140 stu29:410 stu30:370 stu31:760 stu33:820 stu34:800 stu35:900 stu36:520 stu37:780 stu39:810 stu41:750 stu42:480 stu45:230 stu46:730 stu49:930 stu50:260 stu52:340 stu55:860 stu56:940 stu57:470 stu58:990 stu59:840 stu62:30 stu63:50 stu67:770 stu71:240 stu72:540 stu73:440 stu75:570 stu76:790 stu82:640 stu84:530 stu89:920 stu90:320 stu92:740 stu93:910 stu94:180 stu95:490 stu96:10 stu97:960 stu98:250 stu99:360]

// [stu28 stu89 stu55 stu16 stu76 stu59 stu21 stu12 stu96 stu05 stu24 stu33 stu57 stu92 stu36 stu90 stu49 stu73 stu95 stu29 stu34 stu27 stu63 stu94 stu67 stu45 stu19 stu35 stu84 stu15 stu18 stu10 stu52 stu02 stu41 stu42 stu93 stu11 stu50 stu71 stu97 stu20 stu07 stu46 stu62 stu72 stu99 stu58 stu39 stu22 stu13 stu31 stu17 stu30 stu37 stu56 stu75 stu98 stu23 stu82]

// [stu02 stu05 stu07 stu10 stu11 stu12 stu13 stu15 stu16 stu17 stu18 stu19 stu20 stu21 stu22 stu23 stu24 stu27 stu28 stu29 stu30 stu31 stu33 stu34 stu35 stu36 stu37 stu39 stu41 stu42 stu45 stu46 stu49 stu50 stu52 stu55 stu56 stu57 stu58 stu59 stu62 stu63 stu67 stu71 stu72 stu73 stu75 stu76 stu82 stu84 stu89 stu90 stu92 stu93 stu94 stu95 stu96 stu97 stu98 stu99]

// stu02 100
// stu05 830
// stu07 420
// stu10 650
// stu11 980
// stu12 880
// stu13 950
// stu15 610
// stu16 70
// stu17 630
// stu18 460
// stu19 680
// stu20 970
// stu21 870
// stu22 700
// stu23 660
// stu24 850
// stu27 600
// stu28 140
// stu29 410
// stu30 370
// stu31 760
// stu33 820
// stu34 800
// stu35 900
// stu36 520
// stu37 780
// stu39 810
// stu41 750
// stu42 480
// stu45 230
// stu46 730
// stu49 930
// stu50 260
// stu52 340
// stu55 860
// stu56 940
// stu57 470
// stu58 990
// stu59 840
// stu62 30
// stu63 50
// stu67 770
// stu71 240
// stu72 540
// stu73 440
// stu75 570
// stu76 790
// stu82 640
// stu84 530
// stu89 920
// stu90 320
// stu92 740
// stu93 910
// stu94 180
// stu95 490
// stu96 10
// stu97 960
// stu98 250
// stu99 360
```

## 元素为 map 类型的切片

```go
var mapsli = make([]map[string]int, 3)
for index, value := range mapsli {
	fmt.Printf("index:%d value:%v\n", index, value)
    // index:0 value:map[]
    // index:1 value:map[]
    // index:2 value:map[]
}
fmt.Println("after init...")
//对切片中的map元素初始化
mapsli[0] = make(map[string]string, 10)
mapsli[0]["1"] = "1"
mapsli[0]["2"] = "2"
mapsli[0]["3"] = "3"
mapsli[0]["4"] = "4"
mapsli[0]["5"] = "5"
mapsli[0]["6"] = "6"
mapsli[0]["7"] = "7"
mapsli[0]["8"] = "8"
mapsli[0]["9"] = "9"
mapsli[0]["10"] = "10"
mapsli[0]["11"] = "11"
for index, value := range mapsli {
	fmt.Printf("index:%d value:%v\n", index, value)
// index:0 value:map[1:1 10:10 11:11 2:2 3:3 4:4 5:5 6:6 7:7 8:8 9:9]
// index:1 value:map[]
// index:2 value:map[]
}
```

## 值为切片类型的map

```go
var slimap = make(map[string][]string, 10)
fmt.Println(slimap)
//map[]
fmt.Println("after init ...")
tkey := "中国"
tvalue, tok := slimap[tkey]
if !tok {
	tvalue = make([]string, 0, 3)
}
tvalue = append(tvalue, "北京", "上海")
slimap[tkey] = tvalue
fmt.Println(slimap)
//map[中国:[北京 上海]]
```

