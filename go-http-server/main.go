package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Home page\n")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/hello", hello)
	fmt.Println("Server is running on http://localhost:8090")
	http.ListenAndServe(":8090", nil)
}
