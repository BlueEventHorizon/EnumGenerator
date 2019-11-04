package main

import (
	"flag"
	"fmt"
	"strings"

	"./analyzer"
)

func main() {
	flag.Parse()
	topDir := flag.Arg(0)
	enumName := flag.Arg(1)
	if topDir == "" {
		topDir = "./"
	}
	if enumName == "" {
		enumName = "LocalizableStrings"
	}

	output(fmt.Sprintf("import Foundation\n\n"))
	output(fmt.Sprintf("enum %s: String {\n", enumName))

	texts := make([]string, 100, 500)
	Scandir(topDir, analyzer.LocalisableStringsAnalyzer, &texts)
	for _, text := range texts {
		if text == "" {
			continue
		}
		// 空白はアンダースコアに置換
		keyword := strings.Replace(text, " ", "_", -1)
		keyword = convertToCamelCase(keyword)
		output(fmt.Sprintf("    case %s = \"%s\",\n", keyword, text))
	}
	output(fmt.Sprintf("}\n"))

	assets := make([]string, 100, 500)
	Scandir(topDir, analyzer.AssetAnalyzer, &assets)
}

func output(text string) {
	fmt.Print(text)
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
