package pkg1

import "fmt"

//City ...
var City = "杭州"

//Test1 ...
func Test1() {
	fmt.Println("pkg1")
}

//Student ...
type Student struct {
	id   int
	name string
}

//NewStudent ...
func NewStudent(id int, name string) Student {
	return Student{
		id:   id,
		name: name,
	}
}

func init() {
	fmt.Println("这是pkg1的init...")
}
