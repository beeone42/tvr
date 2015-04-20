package main

import (
	"path"
	"path/filepath"
	"log"
)


func listVideo() ([]string, error) {
	res, err := filepath.Glob(path.Join("video", "*.mp4, *.MP4"))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(res)

	for i := range res {
		res[i] = path.Base(res[i])
	}
	return res, nil
}

