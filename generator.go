package main

import (
	"flag"
	"fmt"
	"strings"

	"./analyzer"
	"./env"
)

func main() {

	// flagの使い方
	// https://qiita.com/Yaruki00/items/7edc04720a24e71abfa2

	var (
		topDir      string
		enumName    string
		enableImage bool
		enableColor bool
		//	enableImage = flag.Bool("image", true, "enable scan for image assets")
		//	enableColor = flag.Bool("color", true, "enable scan for color assets")
	)

	flag.StringVar(&topDir, "dir", "./", "dir to scan")
	flag.StringVar(&enumName, "string", "LocalizableStrings", "enum name for Localizable.strings")
	flag.BoolVar(&enableImage, "image", true, "enable scan for image assets")
	flag.BoolVar(&enableColor, "color", true, "enable scan for color assets")

	//flag.Parse()

	output("import Foundation\n\n")
	output(fmt.Sprintf("enum %s: String {\n", enumName))

	texts := make([]string, 100, 500)
	ScanFile(topDir, analyzer.LocalisableStringsAnalyzer, &texts)
	for _, text := range texts {
		if text == "" {
			continue
		}
		// 空白はアンダースコアに置換
		keyword := strings.Replace(text, " ", "_", -1)
		// ピリオドはアンダースコアに置換
		keyword = strings.Replace(keyword, " ", ".", -1)
		// ハイフンはアンダースコアに置換
		keyword = strings.Replace(keyword, " ", "-", -1)

		keyword = convertToCamelCase(keyword)
		output(fmt.Sprintf("    case %s = \"%s\",\n", keyword, text))
	}
	output("}\n")

	if enableImage {
		imageAssets := make([]string, 100, 500)
		ScanDir(topDir, analyzer.ImageAssetAnalyzer, &imageAssets)
		for _, asset := range imageAssets {
			if asset == "" {
				continue
			}
			output(fmt.Sprintf("imageAssets = \"%s\",\n", asset))
		}
	}

	if enableColor {
		colorAssets := make([]string, 100, 500)
		ScanDir(topDir, analyzer.ColorAssetAnalyzer, &colorAssets)
		for _, asset := range colorAssets {
			if asset == "" {
				continue
			}
			output(fmt.Sprintf("colorAssets = \"%s\",\n", asset))
		}
	}
}

func open(text string) {
	if env.OUTPUT_FILE {

	}
}

func output(text string) {
	if env.OUTPUT_FILE {

	} else {
		fmt.Print(text)
	}
}

func close(text string) {
	if env.OUTPUT_FILE {

	}
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
