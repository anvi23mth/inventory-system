package handler

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello World")
		return
	}

	fmt.Fprintf(w, "Hello, %s", name)
}