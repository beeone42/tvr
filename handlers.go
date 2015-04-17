package main

import (
	"encoding/json"
    "html/template"
    "net/http"
	"log"
	"regexp"
	"path"
	"io/ioutil"
)

type Playlist struct {
	Id	string
	Name  string
	Items []string
}

var ajaxValidPath = regexp.MustCompile("^/ajax/(load|save)/([a-zA-Z0-9]+)$")

func loadPlaylist(id string) (*Playlist, error) {
	var pl Playlist
    filename := path.Join("playlist", id + ".json")
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
	err = json.Unmarshal(body, &pl)
    if err != nil {
        return nil, err
    }
    return &pl, nil
}

func savePlaylist(pl Playlist) (error) {
    filename := path.Join("playlist", pl.Id + ".json")
	body, err := json.Marshal(pl)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(filename, body, 0600)
}

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

	m := ajaxValidPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}

	pl, err := loadPlaylist(m[2])
	if err != nil {
		log.Println(err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	//Playlist{m[2], "GoT", []string{"episode1.mp4", "episode2.mp4"}}

	b, err := json.Marshal(pl)
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

