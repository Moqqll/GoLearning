package main

import "fmt"

// //Sayer ...
// type Sayer interface {
// 	say()
// }

//Mover ...
type Mover interface {
	move()
}

// type dog struct {
// 	name string
// }

// //实现Sayer接口
// func (d *dog) say() {
// 	fmt.Printf("%s会汪汪汪~\n", d.name)
// }

// //实现Mover接口
// func (d *dog) move() {
// 	fmt.Printf("%s会动\n", d.name)
// }

// func main() {
// 	var s Sayer
// 	var m Mover

// 	var d = &dog{
// 		name: "wangcai",
// 	}
// 	s = d
// 	s.say()
// 	m = d
// 	m.move()
// }

type dog struct {
	name string
}

type car struct {
	brand string
}

func (d *dog) move() {
	fmt.Printf("%s会动\n", d.name)
}

func (c *car) move() {
	fmt.Printf("%s会动\n", c.brand)
}

func main() {
	var m Mover
	var d = &dog{
		name: "wangcai",
	}
	var c = &car{
		brand: "fute",
	}
	m = d
	m.move()
	m = c
	m.move()

	//定义一个空接口
	var xx interface{}
	aa := "moqqll"
	xx = aa
	fmt.Printf("%T %v\n", xx, xx)
	bb := 100
	xx = bb
	fmt.Printf("%T %v\n", xx, xx)
	cc := false
	xx = cc
	fmt.Printf("%T %v\n", xx, xx)
}
