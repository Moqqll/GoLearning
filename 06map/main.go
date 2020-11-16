package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	var a map[string]int
	fmt.Println(a == nil)
	fmt.Println()

	//map基础使用
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of scoreMap: %T\n", scoreMap)
	fmt.Println()

	//map也支持在声明的时候填充元素
	userInfo := map[string]string{
		"user_name": "moqqll",
		"password":  "moqqll",
	}
	fmt.Println(userInfo)
	fmt.Println()

	//判断某个键值对是否存在
	v, ok := scoreMap["李四"] //不存在时v的值为0
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
		fmt.Println(v)
	}
	fmt.Println()

	//使用for range 遍历map
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	for k := range scoreMap { //只遍历key
		fmt.Println(k)
	}
	for _, v := range scoreMap { //忽略key，只取value值
		fmt.Println(v)
	}
	fmt.Println()

	//使用delete()函数删除map中的一组键值对
	delete(scoreMap, "张三")
	for k, v := range scoreMap {
		fmt.Println(k, v) //李四 100
	}
	fmt.Println()

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
	fmt.Println()

	//元素为map类型的切片
	var mapsli = make([]map[string]string, 3)
	for index, value := range mapsli {
		fmt.Printf("index:%d value:%v\n", index, value)
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
	}
	fmt.Println()

	//值为切片类型的map
	var slimap = make(map[string][]string, 3)
	fmt.Println(slimap)
	fmt.Println("after init ...")
	tkey := "中国"
	tvalue, tok := slimap[tkey]
	if !tok {
		tvalue = make([]string, 0, 3)
	}
	tvalue = append(tvalue, "北京", "上海")
	slimap[tkey] = tvalue
	fmt.Println(slimap)
}
