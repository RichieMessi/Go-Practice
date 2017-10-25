package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/boo/", boo)
	http.ListenAndServe(":1000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I am the best guy in the world")
}

func boo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is from function boo")
}
