package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func isSkipDir(name string) bool {
	if strings.HasSuffix(name, ".xcodeproj") {
		//fmt.Printf("// ---------------------------------------- xcodeproj\n")
		return true
	}
	if strings.HasSuffix(name, ".xcworkspace") {
		//fmt.Printf("// ---------------------------------------- xcworkspace\n")
		return true
	}
	if name == "build" {
		return true
	}
	// Carthageディレクトリはスキャンしない
	if name == "Carthage" {
		//fmt.Printf("// ---------------------------------------- Carthage\n")
		return true
	}
	// Podsディレクトリはスキャンしない
	if name == "Pods" {
		//fmt.Printf("// ---------------------------------------- Pods\n")
		return true
	}
	if name == ".git" {
		return true
	}
	return false
}

func ScanDir(dir string, analyzer func(string, *[]string), texts *[]string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		name := file.Name()
		path := filepath.Join(dir, name)
		if file.IsDir() {
			if isSkipDir(name) {
				continue
			}
			analyzer(path, texts)
			ScanDir(filepath.Join(path), analyzer, texts)
			continue
		}
	}
}

func ScanFile(dir string, analyzer func(string, *[]string), texts *[]string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		name := file.Name()
		path := filepath.Join(dir, name)
		if file.IsDir() {
			if isSkipDir(name) {
				continue
			}
			ScanFile(filepath.Join(path), analyzer, texts)
			continue
		}

		analyzer(path, texts)
	}
}
