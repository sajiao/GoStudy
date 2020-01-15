package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", home)

	http.HandleFunc("/sum", add)
	http.ListenAndServe(":80", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to yanzi travel")
}

func add(w http.ResponseWriter, r *http.Request) {

	n1, err := strconv.Atoi(r.URL.Query().Get("a"))
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	n2, err := strconv.Atoi(r.URL.Query().Get("b"))
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, strconv.Itoa(n1+n2))
}

func multiply(n1, n2 int) int {
	return n1 * n2
}
