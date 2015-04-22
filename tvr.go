package main

import (
    "net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/video", videoHandler)					// la page pour editer une playlist
	http.HandleFunc("/video/execute", videoExecuteHandler)	// la page pour sauver suite a l'edition d'une playlist
	http.HandleFunc("/video/create", videoCreateHandler)	// formulaire de creation d'une playlist vide
	http.HandleFunc("/video/create/execute", videoCreateExecuteHandler) // sauvegarde de la playlist vide
	http.HandleFunc("/video/upload", videoUploadHandler)	// formulaire d'upload d'une video dans la bibliotheque
	http.HandleFunc("/video/upload/receive", videoUploadReceiveHandler)	
	http.HandleFunc("/ajax/list/", ajaxListHandler)			// la liste des playlists
	http.HandleFunc("/ajax/video/", ajaxVideoHandler)		// la liste des videos
	http.HandleFunc("/ajax/load/", ajaxLoadHandler)			// le contenu d'une playlist
	http.HandleFunc("/ajax/publish/", ajaxPublishHandler)	// publie une playlist sur une tele
	http.HandleFunc("/ajax/state/", ajaxStateHandler)		// recupere le status courant sur une tele
//	http.HandleFunc("/ajax/save/", ajaxSaveHandler)			// sauver une playlist
	http.ListenAndServe("127.0.0.1:3000", nil)
}
