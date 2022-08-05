package main

import (
	"encoding/json"
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

	http.HandleFunc("/getSettings", func(w http.ResponseWriter, r *http.Request) {
		val, err := json.Marshal(configuration)
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(w, string(val))
	})

	http.HandleFunc("/setSettings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}
		if r.ParseForm() != nil {
			return
		}
		settingsString := r.PostForm.Get("settings")
		conf := Configuration{}

		err := json.Unmarshal([]byte(settingsString), &conf)
		if err != nil {
			fmt.Println(err)
			return
		}

		configuration = conf
		fmt.Println("Applied settings; ", configuration)
	})

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(configuration.WebPort), nil))
}
