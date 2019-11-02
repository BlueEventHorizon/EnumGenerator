package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	firstFlag := flag.Arg(0)
	secondFlag := flag.Arg(1)
	if firstFlag == "" {
		firstFlag = "./"
	}
	if secondFlag == "" {
		secondFlag = "LocalizableStrings"
	}

	fmt.Printf("import Foundation\n\n")
	fmt.Printf("enum %s: String {\n", secondFlag)
	texts := make([]string, 100, 500)
	Scandir(firstFlag, &texts)
	for _, text := range texts {
		if text == "" {
			continue
		}
		// 空白はアンダースコアに置換
		keyword := strings.Replace(text, " ", "_", -1)
		keyword = convertToCamelCase(keyword)
		fmt.Printf("    case %s = \"%s\",\n", keyword, text)
	}
	fmt.Printf("}\n")
}

func convertToCamelCase(text string) string {
	var keyword string
	var foundUnderScore = false
	for i := 0; i < len(text); i++ {
		letter := text[i : i+1]
		if letter == "_" {
			foundUnderScore = true
			continue
		}
		if foundUnderScore {
			foundUnderScore = false
			keyword = keyword + strings.ToUpper(letter)
		} else {
			keyword = keyword + letter
		}
	}

	return keyword
}
