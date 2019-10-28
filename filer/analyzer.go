package filer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Analyzer(path string, texts *[]string) {

	//var counter int = 0

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

			// 空白はアンダースコアに置換
			//keyword := strings.Replace(text, " ", "_", -1)
			// case
			//fmt.Printf("    case %s = \"%s\",\n", keyword, text)
		}
	}
}