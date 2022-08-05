package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func webServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, index)
	})

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		if r.PostForm.Has("x") {

		}
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(configuration.WebPort), nil))
}
