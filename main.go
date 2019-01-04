package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

var images = []string{"jpg", "jpeg", "png", "gif"}

const isoDate = "2006-01-02"

func main() {
	startDir := "./"
	files, err := ioutil.ReadDir(startDir)
	if err != nil {
		log.Fatal("Couldn't find directory ", startDir, err)
	}

	for _, fileInfo := range files {
		imagePath := fileInfo.Name()
		if isImage(imagePath) {
			f, err := os.Open(imagePath)
			if err != nil {
				log.Fatal("Couldn't open file ", err)
			}
			defer f.Close()
			x, err := exif.Decode(f)
			if err != nil {
				log.Fatal("Couldn't read exif ", err)
			}
			ts, err := x.DateTime()
			if err != nil {
				log.Fatal("Couldn't read datetime ", err)
			}
			imageFolder := ts.Format(isoDate)
			createDirIfDoesNotExist(imageFolder)
			err = os.Rename(imagePath, path.Join("./", imageFolder, imagePath))
			if err != nil {
				log.Fatal("Couldn't move image ", err)
			}
		}
	}
}

func isImage(img string) bool {
	for _, typ := range images {
		if strings.HasSuffix(strings.ToLower(img), typ) {
			return true
		}
	}
	return false
}

func printExif(f *os.File) {
	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal("Couldn't read exif ", err)
	}
	fmt.Printf("%+q\n", x)
}

func createDirIfDoesNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}
