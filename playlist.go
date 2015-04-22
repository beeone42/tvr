package main

import (
	"path"
	"path/filepath"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Playlist struct {
	Id	string
	Name  string
	Author string
	Items []string
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
    return ioutil.WriteFile(filename, body, 0644)
}
