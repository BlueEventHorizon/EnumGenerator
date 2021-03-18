package main

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"./analyzer"
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

func ScanDir(dir string, analyzer func(string, *[]analyzer.AnalyzedInfrmation), infos *[]analyzer.AnalyzedInfrmation) {
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
			analyzer(path, infos)
			ScanDir(filepath.Join(path), analyzer, infos)
			continue
		}
	}
}

func ScanFile(dir string, analyzer func(string, *[]analyzer.AnalyzedInfrmation), infos *[]analyzer.AnalyzedInfrmation) {
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
			ScanFile(filepath.Join(path), analyzer, infos)
			continue
		}

		analyzer(path, infos)
	}
}
