package main

import (
    "net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/ajax/", ajaxHandler)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
