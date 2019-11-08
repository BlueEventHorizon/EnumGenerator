package analyzer

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func AssetAnalyzer(path string, texts *[]string) {
	// xxx.xcassets のみを解析
	if !strings.HasSuffix(path, ".xcassets") {
		return
	}
	scanAssetDir(path, assetAnalyzer, texts)
}

func scanAssetDir(dir string, analyzer func(string, *[]string), texts *[]string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	println(dir)

	for _, file := range files {
		name := file.Name()
		path := filepath.Join(dir, name)
		if file.IsDir() {
			if strings.HasSuffix(name, ".imageset") {
				println(name)
				continue
			}
			if strings.HasSuffix(name, ".colorset") {
				println(name)
				continue
			}

			//analyzer(path, texts)
			scanAssetDir(filepath.Join(path), analyzer, texts)
			continue
		}
	}
}

func assetAnalyzer(path string, texts *[]string) {
	// xxx.strings のみを解析
	if !strings.HasSuffix(path, ".xcassets") {
		return
	}

	// ファイルをOpenする
	file, err := os.Open(path)
	// 読み取り時の例外処理
	if err != nil {
		fmt.Println("error")
	}
	// 関数が終了した際に確実に閉じるようにする
	defer file.Close()

	sc := bufio.NewScanner(file)
	for i := 1; sc.Scan(); i++ {
		if err := sc.Err(); err != nil {
			// エラー処理
			break
		}

		text := sc.Text()
		text = strings.TrimSpace(text)
		index := strings.Index(text, "//")
		if index > 0 {
			text = text[:index]
			text = strings.TrimSpace(text)
		}

		if strings.HasSuffix(text, ";") {
			index := strings.Index(text, "=")
			if index > 0 {
				text = text[:index]
				text = strings.TrimSpace(text)
				text = strings.Trim(text, "\"")
			}

			var result bool = false
			for _, element := range *texts {
				if element == text {
					result = true
					break
				}
			}
			if result == false {
				*texts = append(*texts, text)
			}
		}
	}
}
