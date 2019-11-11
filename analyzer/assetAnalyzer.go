package analyzer

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ImageAssetAnalyzer(path string, texts *[]string) {
	// xxx.xcassets のみを解析
	if !strings.HasSuffix(path, ".xcassets") {
		return
	}
	scanAssetDir(path, ".imageset", texts)
}

func ColorAssetAnalyzer(path string, texts *[]string) {
	// xxx.xcassets のみを解析
	if !strings.HasSuffix(path, ".xcassets") {
		return
	}
	scanAssetDir(path, ".colorset", texts)
}

func scanAssetDir(dir string, assetKey string, texts *[]string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		name := file.Name()
		path := filepath.Join(dir, name)
		if file.IsDir() {
			if strings.HasSuffix(name, assetKey) {
				index := strings.Index(path, ".xcassets")
				if index > 0 {
					text := path[index+len(".xcassets"):]
					*texts = append(*texts, text)
				}
				continue
			}
			scanAssetDir(filepath.Join(path), assetKey, texts)
			continue
		}
	}
}
