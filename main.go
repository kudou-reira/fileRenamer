package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/djherbis/times"
)

type fileInfo struct {
	birthTime time.Time
	name      string
}

func main() {
	// exe, err := os.Getwd()
	// if err != nil {
	// 	panic(err)
	// }
	// exePath := filepath.Dir(exe)
	// fmt.Println(exePath)

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var fileCollection []fileInfo

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".mp4") {
			t, err := times.Stat(file.Name())
			if err != nil {
				log.Fatal(err.Error())
			}

			if t.HasBirthTime() {
				bTime := t.BirthTime()
				tempFileInfo := fileInfo{
					birthTime: bTime,
					name:      file.Name(),
				}
				fileCollection = append(fileCollection, tempFileInfo)
			}
		}
	}

	fileCollection = sortByBirthDate(fileCollection)
	rename(fileCollection)
}

func rename(collection []fileInfo) {
	for i, j := range collection {
		newVal := fmt.Sprintf("%03d", i+1)
		newName := newVal + " " + j.name
		err := os.Rename(j.name, newName)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func sortByBirthDate(collection []fileInfo) []fileInfo {
	sort.Slice(collection, func(i, j int) bool {
		return collection[i].birthTime.Unix() < collection[j].birthTime.Unix()
	})

	return collection
}
