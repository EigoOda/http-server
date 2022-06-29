package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler struct{}
type thisHandler struct{}
type healthcheck struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo")
}

func (h *thisHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bar")
}

func (h *healthcheck) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	http.Handle("/foo", &myHandler{})
	http.Handle("/bar", &thisHandler{})
	http.Handle("/healthcheck", &healthcheck{})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
