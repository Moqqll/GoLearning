package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

//随机出现慢响应
func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	if number == 0 {
		time.Sleep(time.Second * 10) //耗时10s的慢响应
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprintf(w, "quick response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
