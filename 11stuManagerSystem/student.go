package main

import "fmt"

type student struct {
	id    int //学号唯一
	name  string
	class string
}

//student ...构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

//Students ...
type studentMag struct {
	students []*student
}

//newStudents ...构造函数
func newstudentMag() *studentMag {
	return &studentMag{
		students: make([]*student, 0, 100),
	}
}

//添加学员
func (sm *studentMag) addStudent(newStu *student) {
	sm.students = append(sm.students, newStu)
}

//编辑学员
func (sm *studentMag) editStudent(newStu *student) {
	for i, v := range sm.students {
		if newStu.id == v.id { //当学号相同时，就表示找到了需要编辑的学员
			sm.students[i] = newStu //把新学员信息根据切片索引直接赋能给切片相应元素项，即用新的学员信息替换旧的学员信息。
		}
	}
}

//展示学员
func (sm *studentMag) showStudent() {
	fmt.Println("系统中已有以下学员：")
	for _, v := range sm.students {
		fmt.Printf("id:%d name:%s class:%s\n", v.id, v.name, v.class)
	}
}
