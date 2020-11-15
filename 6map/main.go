package main

import (
	"fmt"
)

func main() {

	var a map[string]int
	fmt.Println(a == nil)

	//map基础使用
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["张三"])
	fmt.Printf("type of scoreMap: %T\n", scoreMap)

	//map也支持在声明的时候填充元素
	userInfo := map[string]string{
		"user_name": "moqqll",
		"password":  "moqqll",
	}
	fmt.Println(userInfo)

	//判断某个键值对是否存在
	v, ok := scoreMap["李四"] //不存在时v的值为0
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
		fmt.Println(v)
	}

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

	//使用delete()函数删除map中的一组键值对
	delete(scoreMap, "张三")
	for k, v := range scoreMap {
		fmt.Println(k, v) //李四 100
	}
}
