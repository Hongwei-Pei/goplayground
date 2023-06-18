package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func AULFHandler(w http.ResponseWriter, r *http.Request) {
	type T struct {
		Name string
	}
	for i := 0; i < 10000; i++ {
		n := fmt.Sprintf("test: %d", i)
		t := T{Name: n}
		fmt.Println(t.Name)
		w.Write([]byte(t.Name))
	}
}

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/h", AULFHandler)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
