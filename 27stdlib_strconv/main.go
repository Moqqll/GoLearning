package main

import (
	"fmt"
	"strconv"
)

func myatoi() {
	s1 := "100"
	il, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("cant't convert to int, err:", err)
		return
	}
	fmt.Printf("type:%T value:%#v\n", il, il)
}

func myitoa() {
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T,value:%#v\n", s2, s2)
}

func myparse() {
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	i, _ := strconv.ParseInt("2", 10, 64)
	u, _ := strconv.ParseUint("2", 10, 64)
	fmt.Println(b, f, i, u)
}

func mystrconv() {
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)
	s4 := strconv.FormatUint(2, 16)
	fmt.Printf("%#v,%#v,%#v,%#v", s1, s2, s3, s4)
}

func main() {
	// myatoi()
	// myitoa()
	// myparse()
	mystrconv()

}
