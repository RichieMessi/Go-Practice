package main

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/", naruto)

	http.HandleFunc("/pain", pain)

	http.ListenAndServe(":8000", nil)

}

func naruto(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I am from village hidden in the leaf\n")
}

func pain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Village hidden in the rain\n")
}
