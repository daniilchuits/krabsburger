package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет из ServeHTTP!")
}

func main() {
	handler := MyHandler{}
	http.ListenAndServe(":8080", handler)
}
