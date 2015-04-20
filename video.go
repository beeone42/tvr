package main

import (
	"path"
	"path/filepath"
	"log"
)


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

