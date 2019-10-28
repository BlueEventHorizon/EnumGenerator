package filer

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile2(path string) {
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
		fmt.Printf("%4d行目: %s\n", i, sc.Text())
	}
}
