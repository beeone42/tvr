package main

import (
    "net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/ajax/", ajaxHandler)
	http.ListenAndServe(":8080", nil)
}
