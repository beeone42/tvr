package main

import (
    "net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/video", videoHandler)
	http.HandleFunc("/video/create", videoCreateHandler)
	http.HandleFunc("/ajax/list/", ajaxListHandler) // la liste des playlists
	http.HandleFunc("/ajax/load/", ajaxLoadHandler) // le contenu d'une playlist
//	http.HandleFunc("/ajax/save/", ajaxSaveHandler) // sauver une playlist
	http.ListenAndServe("127.0.0.1:3000", nil)
}
