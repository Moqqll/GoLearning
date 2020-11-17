package main

import (
	"fmt"
	"os"
)

//学员信息管理系统

//需求：
//1、添加学员信息
//2、编辑学员信息
//3、展示所有学员信息
//4、退出系统

func showMenu() {
	fmt.Println()
	fmt.Println("欢迎来到学员信息管理系统：")
	fmt.Println("1、添加学员信息")
	fmt.Println("2、编辑学员信息")
	fmt.Println("3、展示所有学员信息")
	fmt.Println("4、退出系统")
	fmt.Println()
}

func getInput() *student {
	defer fmt.Println("输入完成！")
	var (
		id    int
		name  string
		class string
	)
	//拿到用户输入的学员的信息
	fmt.Println("请按要求输入学员信息")
	fmt.Print("请输入学员的id：")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学员的name：")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学员的class：")
	fmt.Scanf("%s\n", &class)
	tmpstu := newStudent(id, name, class)
	return tmpstu
}

func main() {
	sm := newstudentMag()
	fmt.Print("初始学员信息为空：")
	fmt.Println(sm.students)
	for {
		//1、打印系统菜单
		showMenu()
		//2、等待用户选择要执行的选项
		var input int
		fmt.Print("请输入你的选择：")
		fmt.Scanf("%d\n", &input)
		fmt.Print("你输入的是：")
		fmt.Println(input)
		fmt.Println()
		//3、执行用户选择的动作
		switch input {
		case 1:
			//添加学员
			tmpstu := getInput()
			sm.addStudent(tmpstu)
		case 2:
			//编辑学员
			tmpstu := getInput()
			sm.editStudent(tmpstu)
		case 3:
			//展示所有学员
			sm.showStudent()
		case 4:
			//退出系统
			os.Exit(0)
		default:
			//输入值不满足系统要求
			fmt.Println()
			fmt.Println("输入错误，请重新输入。")
			fmt.Println()
		}
	}
}
