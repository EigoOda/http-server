package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler struct{}
type thisHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo")
}

func (h *thisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar")
}

func main() {
	http.Handle("/foo", &myHandler{})
	http.Handle("/bar", &thisHandler{})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
