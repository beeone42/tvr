package main

import (
	"encoding/json"
    "html/template"
    "net/http"
	"log"
	"regexp"
	"fmt"
)


type Video struct {
	Id string
	Name string
}

var ajaxValidPath = regexp.MustCompile("^/ajax/(load|details|save)/([a-zA-Z0-9]+)$")

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

func videoHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/video"):]
	renderTemplate(w, "header", title)
	renderTemplate(w, "video", title)
	renderTemplate(w, "footer", title)
}

func videoCreateHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/video/create"):]
	renderTemplate(w, "header", title)
	renderTemplate(w, "video_create", title)
	renderTemplate(w, "footer", title)
}


func ajaxListHandler(w http.ResponseWriter, r *http.Request) {
	var b []byte
	res, err := listPlaylist()
	if err != nil {
		log.Println(err.Error())
		return
	}
	b, err = json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func ajaxVideoHandler(w http.ResponseWriter, r *http.Request) {
	var b []byte
	res, err := listVideo()
	if err != nil {
		log.Println(err.Error())
		return
	}
	b, err = json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func ajaxLoadHandler(w http.ResponseWriter, r *http.Request) {
	m := ajaxValidPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	if (m[1] == "load") {
		var b []byte
		res, err := loadPlaylist(m[2])
		if err != nil {
			log.Println(err.Error())
			return
		}
		b, err = json.Marshal(res)
		if err != nil {
			log.Println(err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}
