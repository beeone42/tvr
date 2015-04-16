package main

import (
	"encoding/json"
    "html/template"
    "net/http"
	"log"
)

func renderTemplate(w http.ResponseWriter, tmpl string, title string) {
    t, err := template.ParseFiles("views/" + tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, title)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/"):]
	renderTemplate(w, "header", title)
	renderTemplate(w, "main", title)
	renderTemplate(w, "footer", title)
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal([]byte("salut"))
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(b);
}

