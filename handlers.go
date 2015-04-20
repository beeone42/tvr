package main

import (
	"encoding/json"
    "html/template"
    "net/http"
	"log"
	"regexp"
	"fmt"
	"os"
	"io"
	"path"
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

func videoUploadHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/video/upload"):]
	renderTemplate(w, "header", title)
	renderTemplate(w, "video_upload", title)
	renderTemplate(w, "footer", title)
}

func videoUploadReceiveHandler(w http.ResponseWriter, r *http.Request) {
         file, header, err := r.FormFile("file")
         if err != nil {
                 log.Println(err)
                 return
         }
         defer file.Close()
         out, err := os.Create(path.Join("video", header.Filename))
         if err != nil {
                 log.Println("Unable to create the file for writing. Check your write access privilege")
                 return
         }
         defer out.Close()
         // write the content from POST to the file
         _, err = io.Copy(out, file)
         if err != nil {
                 log.Println(err)
         }

         log.Println( "File uploaded successfully : ")
         log.Println( header.Filename)
         http.Redirect(w, r, "/ajax/upload/ok", http.StatusFound)
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
