package main

import (
    "net/http"
)

func main() {
//	lala()
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/video", videoHandler)
	http.HandleFunc("/video/create", videoCreateHandler)
	http.HandleFunc("/ajax/", ajaxHandler)
	http.ListenAndServe("127.0.0.1:3000", nil)
}
