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
	"io/ioutil"
	"bytes"
	"path"
	"strings"
//	"strconv"
	"os/exec"
)


type Video struct {
	Id string
	Name string
}

type Hash struct {
    PlayerId int
    Type string
}

type KodiResult struct {
	Id int
	JsonVersion string
	Result []Hash
}

var ajaxValidPath = regexp.MustCompile("^/ajax/(load|details|save)/([a-zA-Z0-9]+)$")
var ajaxValidPublishPath = regexp.MustCompile("^/ajax/publish/(left|right)/([a-zA-Z0-9]+)$")
var ajaxValidStatePath = regexp.MustCompile("^/ajax/state/(left|right)$")

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
	renderTemplate(w, "header", "main")
	renderTemplate(w, "main", title)
	renderTemplate(w, "footer", title)
}

func videoHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/video"):]
	renderTemplate(w, "header", "navbar_video")
	renderTemplate(w, "video", title)
	renderTemplate(w, "footer", title)
}

func videoExecuteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("videoExecuteHandler")
	r.ParseForm()

	id := strings.Join(r.Form["id"], "")
    pl, err := loadPlaylist(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

	log.Println("id=", pl.Id)
	log.Println("name=", pl.Name)
	log.Println("author=", pl.Author)

    pl.Items = r.Form["items"]
    savePlaylist(*pl)

    http.Redirect(w, r, "/video", http.StatusFound)
	w.Write([]byte("ok"))
}

func videoCreateHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/video/create"):]
	renderTemplate(w, "header", "navbar_video_create")
	renderTemplate(w, "video_create", title)
	renderTemplate(w, "footer", title)
}

func videoCreateExecuteHandler(w http.ResponseWriter, r *http.Request) {
	var pl Playlist

	log.Println("videoCreateExecute")
	r.ParseForm()
	pl.Name = strings.Join(r.Form["inputName"], "")
	pl.Id = pl.Name
	pl.Author = strings.Join(r.Form["inputAuthor"], "")

	log.Println("id=", pl.Id)
	log.Println("name=", pl.Name)
	log.Println("author=", pl.Author)

	savePlaylist(pl)

    http.Redirect(w, r, "/video", http.StatusFound)
	w.Write([]byte("ok"))
}

func videoUploadHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/video/upload"):]
	renderTemplate(w, "header", "navbar_video_upload")
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

func ajaxPublishHandler(w http.ResponseWriter, r *http.Request) {
	m := ajaxValidPublishPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	cmd := "scripts/publish_" + m[1] + ".py";
	log.Println(cmd)
	log.Println(m[2])

	c := exec.Command(cmd, m[2])
	if err := c.Run(); err != nil { 
		fmt.Println("Error: ", err)
	}

	w.Write([]byte("ok"))
}


func ajaxStateHandler(w http.ResponseWriter, r *http.Request) {
	var host string;
	m := ajaxValidStatePath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return
	}
	if (m[1] == "left") {
		host = "10.62.1.2"
	} else {
		host = "10.62.1.3"
	}

    url := "http://" + host + "/jsonrpc"
    log.Println("URL:>", url)

// {"jsonrpc": "2.0", "method": "Player.GetActivePlayers", "id": 1}

    players_str, _ := jsonQuery(url + "?Player.GetActivePlayers", []byte(`{"jsonrpc": "2.0", "method": "Player.GetActivePlayers", "id": 1}`))
    //fmt.Println("players_str:", string(players_str))

// {"id":1,"jsonrpc":"2.0","result":[{"playerid":2,"type":"picture"}]}

    var kl KodiResult

	err := json.Unmarshal(players_str, &kl)
    if err != nil {
		log.Println(err.Error())
        return 
    }
	//fmt.Printf("%#v", kl)

	if ((kl.Result != nil) && (len(kl.Result) > 0)) {
		//fmt.Println("PlayerId:", kl.Result[0].PlayerId)
	    //body, _ := jsonQuery(url + "?Playlist.GetItems", []byte("{\"jsonrpc\": \"2.0\", \"method\": \"Playlist.GetItems\", \"params\": { \"properties\": [ \"runtime\" ], \"playlistid\": " + strconv.Itoa(kl.Result[0].PlayerId) + "}, \"id\": 1}"))
	    body, _ := jsonQuery(url + "?Playlist.GetItems", []byte(`{"jsonrpc": "2.0", "method": "Playlist.GetItems", "params": { "properties": [ "runtime" ], "playlistid": 0}, "id": 1}`))
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}

func jsonQuery(url string, jsonStr []byte) ([]byte, error) {

    fmt.Println("request JSON:", string(jsonStr))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
    	log.Println(err.Error())
    	return nil, err
    }
    defer resp.Body.Close()

    //fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    return body, nil
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
