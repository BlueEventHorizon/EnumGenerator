package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./analyzer"
)

func main() {

	// flagの使い方
	// https://qiita.com/Yaruki00/items/7edc04720a24e71abfa2

	var (
		topDir                   string
		enableLocalizableStrings bool
		enableImageResource      bool
		enableColorResource      bool
		enumString               string
		enumImage                string
		enumColor                string

		//	enableImage = flag.Bool("image", true, "enable scan for image assets")
		//	enableColor = flag.Bool("color", true, "enable scan for color assets")
	)

	flag.StringVar(&topDir, "d", "./", "dir to scan")
	flag.BoolVar(&enableLocalizableStrings, "s", true, "enable generate LocalizableStrings")
	flag.BoolVar(&enableImageResource, "i", false, "enable generate ImageResource")
	flag.BoolVar(&enableColorResource, "c", false, "enable generate ColorResource")

	flag.StringVar(&enumString, "enumStringName", "LocalizableStrings", "enum name for Localizable.strings. If blank, disable output")
	flag.StringVar(&enumImage, "enumImageName", "AppResource.ImageResource", "enum name for Image Assets. If blank, disable output")
	flag.StringVar(&enumColor, "enumColorName", "AppResource.ColorResource", "enum name for Color Assets. If blank, disable output")
	flag.Parse()

	// ---- LocalizableStrings ----
	if enableLocalizableStrings {
		stringOutput := new(Output)
		stringOutput.Open(fmt.Sprintf("%s/%s.swift", topDir, enumString))
		stringOutput.Print("import Foundation\n\n")
		stringOutput.Print(fmt.Sprintf("enum %s: String {\n", enumString))

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
			stringOutput.Print(fmt.Sprintf("    case %s = \"%s\",\n", keyword, text))
		}
		stringOutput.Print("}\n")
		stringOutput.Close()
		fmt.Printf("Completed to generate %s\n", enumString)
	} else {
		fmt.Println("Skipped to scan Localizable.strings")
	}

	// ---- imageAssets ----
	if enableImageResource {
		imageOutput := new(Output)
		imageOutput.Open(fmt.Sprintf("%s/%s", topDir, enumImage))
		imageAssets := make([]string, 0, 500)
		ScanDir(topDir, analyzer.ImageAssetAnalyzer, &imageAssets)
		for _, asset := range imageAssets {
			if asset == "" {
				continue
			}
			imageOutput.Print(fmt.Sprintf("imageAssets = \"%s\",\n", asset))
		}
		imageOutput.Close()
		fmt.Printf("Completed to generate %s\n", enumImage)
	} else {
		fmt.Println("Skipped to scan Image Assets")
	}

	// ---- colorAssets ----
	if enableColorResource {
		colorOutput := new(Output)
		colorOutput.Open(fmt.Sprintf("%s/%s", topDir, enumColor))
		colorAssets := make([]string, 0, 500)
		ScanDir(topDir, analyzer.ColorAssetAnalyzer, &colorAssets)
		for _, asset := range colorAssets {
			if asset == "" {
				continue
			}
			colorOutput.Print(fmt.Sprintf("colorAssets = \"%s\",\n", asset))
		}
		colorOutput.Close()
		fmt.Printf("Completed to generate %s\n", enumColor)
	} else {
		fmt.Println("Skipped to scan Color Assets")
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

type Output struct{}

var fd *os.File
var err error

func (t Output) Open(path string) {

	if path == "" {
		return
	}

	fd, err = os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	fd.Seek(0, 0)

}

func (t Output) Print(text string) {
	if fd != nil {
		fd.WriteString(text)
	} else {
		fmt.Print(text)
	}
}

func (t Output) Close() {
	if fd == nil {
		return
	}
	fd.Close()
}
