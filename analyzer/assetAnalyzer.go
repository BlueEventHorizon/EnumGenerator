package analyzer

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ImageAssetAnalyzer(path string, infos *[]AnalyzedInfrmation) {
	// xxx.xcassets のみを解析
	if !strings.HasSuffix(path, ".xcassets") {
		return
	}
	scanAssetDir(path, ".imageset", infos)
}

func ColorAssetAnalyzer(path string, infos *[]AnalyzedInfrmation) {
	// xxx.xcassets のみを解析
	if !strings.HasSuffix(path, ".xcassets") {
		return
	}
	scanAssetDir(path, ".colorset", infos)
}

func scanAssetDir(dir string, assetKey string, infos *[]AnalyzedInfrmation) {
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
				if index >= 0 {
					text := path[index+len(".xcassets"):]
					//*texts = append(*texts, text)
					*infos = append(*infos, AnalyzedInfrmation{path, 0, text})
				}
				continue
			}
			scanAssetDir(filepath.Join(path), assetKey, infos)
			continue
		}
	}
}
