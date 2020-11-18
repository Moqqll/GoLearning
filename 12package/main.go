package main

import (
	"12package/pkg1"
	"12package/test"

	//import包排序遵循“字典序排序”策略
	"fmt"
)

func main() {
	fmt.Println(test.Name) //moqqll
	test.Test2()           //test
	fmt.Println(pkg1.City) //杭州
	pkg1.Test1()           //pkg1

	// fmt.Println(m.Name)
	// m.Test2()
	// fmt.Println()

	// stu1 := pkg1.NewStudent(1, "小王子")
	// fmt.Println(stu1)
}

func init() {
	fmt.Println("这是main包的init...")
}
