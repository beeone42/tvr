package main

import (
	"encoding/json"
    "html/template"
    "net/http"
	"log"
	"regexp"
	"path"
	"path/filepath"
	"io/ioutil"
	"fmt"
)

type Playlist struct {
	Id	string
	Name  string
	Items []string
}

type Video struct {
	Id string
	Name string
}

var ajaxValidPath = regexp.MustCompile("^/ajax/(load|save)/([a-zA-Z0-9]+).*")


// TEST

/*
func lala() ([]string, error) {
	res, err:=  filepath.Glob(path.Join("playlist", "*.json"))
			if err != nil {
		log.Println(err.Error())
		return nil, err
			}
	fmt.Printf("coucou\n")
	log.Println(res)
 	fd, err := ioutil.ReadFile(res[0])
    if err != nil {
        panic(err)
    }
    fmt.Printf(string(fd))
	fd, err = json.Marshal(fd)
	return res, nil
}

*/

// FIN DES TEST







func listVideo() ([]string, error) {
	res, err := filepath.Glob(path.Join("video", "*.json"))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(res)

	var ext string
	for i := range res {
		res[i] = path.Base(res[i])
		ext = filepath.Ext(res[i])
		res[i] = res[i][0:len(res[i]) - len(ext)]
	}
	return res, nil
}

func listPlaylistDetails() ([]string, error) {
	res, err :=  filepath.Glob(path.Join("playlist", "*.json"))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println("\nSTART")
	log.Println(string(res[0]))
	log.Println("\nFIN")

	var ext string
	for i := range res {
		res[i] = path.Base(res[i])
		ext = filepath.Ext(res[i])
		res[i] = res[i][0:len(res[i]) - len(ext)]
	}
	return res, nil
}

func listPlaylist() ([]string, error) {
	res, err := filepath.Glob(path.Join("playlist", "*.json"))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(res)
	//return []string{"game", "ducourant"}, nil

	var ext string
	for i := range res {
		res[i] = path.Base(res[i])
		ext = filepath.Ext(res[i])
		res[i] = res[i][0:len(res[i]) - len(ext)]
	}

	return res, nil
}
/*

func loadPlaylist(id string) (*Playlist, error) {
	var pl Playlist
    filename := path.Join("playlist", id + ".json")
    body, err := ioutil.ReadFile(filename)
    if err != nil {
		log.Println(err.Error())
        return nil, err
    }
	err = json.Unmarshal(body, &pl)
    if err != nil {
		log.Println(err.Error())
        return nil, err
    }
    return &pl, nil
}
*/
/*
func loadPlaylistDetails(id string) (*Playlist, error) {

}
*/
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

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	m := ajaxValidPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	if (m[1] == "load") {
		var b []byte
		var err error
		if (m[2] == "list" && len(m) == 3) {
			log.Println(string(len(m)))
			res, err := listPlaylist()
			if err != nil {
				log.Println(err.Error())
				return
			}
		b, err = json.Marshal(res)
/*
		} else {

			res, err := loadPlaylist(m[2])
			if err != nil {
				log.Println(err.Error())
				return
			}
			b, err = json.Marshal(res)
*/
		} else if (m[2] == "list" && m[3] == "details") {
			res, err := listPlaylistDetails()
			if err != nil {
				log.Println(err.Error())
				return
			}

			b, err = json.Marshal(res)
			log.Println(string(res[0]))

		} else if (m[2] == "video") {
			res, err := listVideo()
			if err != nil {
				
				return
			}
			b, err = json.Marshal(res)

		} else if (m[2] != "video" && m[2] != "list") {
			errorHandler(w, r, http.StatusNotFound)
        	return
		}

	//	} else {
	//		res, err := loadPlaylist(m[2])
	//		if err != nil {
	//			log.Println(err.Error())
	//			return
	//		}
	//		b, err = json.Marshal(res)
	//	}

		if err != nil {
			log.Println(err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	if (m[1] == "save") {
//		pl = Playlist{m[2], r.FormValue("title")
//		err := savePlaylist(pl)
		return
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}
