package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func a(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/", a)
	http.ListenAndServe("127.0.0.1:9000", nil)
}
