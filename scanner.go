package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"./filer"
)

func Scandir(dir string, texts *[]string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		name := file.Name()
		path := filepath.Join(dir, name)
		if file.IsDir() {
			if strings.HasSuffix(name, ".xcodeproj") {
				//fmt.Printf("// ---------------------------------------- xcodeproj\n")
				continue
			}
			if strings.HasSuffix(name, ".xcworkspace") {
				//fmt.Printf("// ---------------------------------------- xcworkspace\n")
				continue
			}

			if name == "build" {
				continue
			}

			// Carthageディレクトリはスキャンしない
			if name == "Carthage" {
				//fmt.Printf("// ---------------------------------------- Carthage\n")
				continue
			}
			// Podsディレクトリはスキャンしない
			if name == "Pods" {
				//fmt.Printf("// ---------------------------------------- Pods\n")
				continue
			}
			Scandir(filepath.Join(path), texts)
			continue
		}
		// xxx.strings のみを解析
		if strings.HasSuffix(path, ".strings") {
			filer.Analyzer(path, texts)
		}
	}
}
