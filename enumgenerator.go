package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"./analyzer"
	"./translator"
)

func main() {

	// flagの使い方
	// https://qiita.com/Yaruki00/items/7edc04720a24e71abfa2

	var (
		topDir                      string
		enableLocalizableStrings    bool
		enableImageResource         bool
		enableColorResource         bool
		enumString                  string
		enumImage                   string
		enumColor                   string
		localizableStringsExtension string
	)

	flag.StringVar(&topDir, "d", "./", "dir to scan")
	flag.BoolVar(&enableLocalizableStrings, "s", true, "enable generate LocalizableStrings")
	flag.BoolVar(&enableImageResource, "i", false, "enable generate ImageResource")
	flag.BoolVar(&enableColorResource, "c", false, "enable generate ColorResource")

	flag.StringVar(&enumString, "enumStringName", "LocalizableStrings", "enum name for Localizable.strings. If blank, disable output")
	flag.StringVar(&enumImage, "enumImageName", "AppResource.ImageResource", "enum name for Image Assets. If blank, disable output")
	flag.StringVar(&enumColor, "enumColorName", "AppResource.ColorResource", "enum name for Color Assets. If blank, disable output")

	flag.StringVar(&localizableStringsExtension, "ext", "", "You can add some extension after words")

	flag.Parse()

	// ---- LocalizableStrings ----
	if enableLocalizableStrings {
		stringOutput := new(Output)
		stringOutput.Open(fmt.Sprintf("%s/%s.swift", topDir, enumString))
		stringOutput.Print("import Foundation\n\n")
		stringOutput.Print(fmt.Sprintf("internal struct %s {\n", enumString))

		infos := make([]analyzer.AnalyzedInfrmation, 0, 500)
		ScanFile(topDir, analyzer.LocalisableStringsAnalyzer, &infos)
		for _, element := range infos {
			contentText := element.Description
			if contentText == "" {
				continue
			}
			keyword, err := translator.TranslateText("en", contentText)
			if err != nil {
				keyword = ""
			}

			// 空白はアンダースコアに置換
			keyword = strings.Replace(keyword, " ", "_", -1)
			// .はアンダースコアに置換
			keyword = strings.Replace(keyword, ".", "_", -1)
			// +-はアンダースコアに置換
			keyword = strings.Replace(keyword, "+", "_", -1)
			keyword = strings.Replace(keyword, "-", "_", -1)
			// "は消去
			keyword = strings.Replace(keyword, "\"", "", -1)
			// ?は消去
			keyword = strings.Replace(keyword, "?", "", -1)
			// !は消去
			keyword = strings.Replace(keyword, "!", "", -1)
			keyword = strings.Replace(keyword, "“", "", -1)
			keyword = strings.Replace(keyword, "”", "", -1)
			keyword = strings.Replace(keyword, ":", "", -1)
			keyword = strings.Replace(keyword, "[", "", -1)
			keyword = strings.Replace(keyword, "]", "", -1)
			keyword = strings.Replace(keyword, "`", "", -1)
			keyword = strings.Replace(keyword, "'", "", -1)
			keyword = strings.Replace(keyword, "#", "", -1)
			keyword = strings.Replace(keyword, "$", "", -1)
			keyword = strings.Replace(keyword, "%", "", -1)
			keyword = strings.Replace(keyword, "=", "", -1)
			keyword = strings.Replace(keyword, "@", "", -1)
			keyword = strings.Replace(keyword, "\\", "", -1)
			keyword = strings.Replace(keyword, "(", "", -1)
			keyword = strings.Replace(keyword, ")", "", -1)
			keyword = strings.Replace(keyword, ",", "", -1)
			keyword = strings.Replace(keyword, "/", "", -1)

			keyword = strings.Replace(keyword, "&39;", "", -1)
			keyword = strings.Replace(keyword, "&amp;", "", -1)
			keyword = strings.Replace(keyword, "&quot;", "", -1)
			keyword = strings.Replace(keyword, ";", "", -1)

			keyword = convertToCamelCase(keyword)
			stringOutput.Print(fmt.Sprintf("    static let %s = \"%s\"%s\n", keyword, contentText, localizableStringsExtension))
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
		imageAssets := make([]analyzer.AnalyzedInfrmation, 0, 500)
		ScanDir(topDir, analyzer.ImageAssetAnalyzer, &imageAssets)
		for _, element := range imageAssets {
			if element.Description == "" {
				continue
			}
			imageOutput.Print(fmt.Sprintf("imageAssets = \"%s\",\n", element.Description))
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
		colorAssets := make([]analyzer.AnalyzedInfrmation, 0, 500)
		ScanDir(topDir, analyzer.ColorAssetAnalyzer, &colorAssets)
		for _, element := range colorAssets {
			if element.Description == "" {
				continue
			}
			colorOutput.Print(fmt.Sprintf("colorAssets = \"%s\",\n", element.Description))
		}
		colorOutput.Close()
		fmt.Printf("Completed to generate %s\n", enumColor)
	} else {
		fmt.Println("Skipped to scan Color Assets")
	}
}

func convertToCamelCase(text string) string {
	if text == "" {
		return text
	}

	var keyword string
	var foundUnderScore = false
	for i := 0; i < len(text); i++ {
		letter := text[i : i+1]
		if letter == "_" {
			// max 40 letters
			if i > 40 {
				break
			}
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

	head := keyword[:1]
	rest := keyword[1:]
	keyword = strings.ToLower(head) + rest

	return keyword
}

type Output struct{}

var fd *os.File
var err error

func (t Output) Open(path string) {
	return
	if path == "" {
		return
	}

	fd, err = os.Create(path)
	//fd, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fd.Seek(0, 0)

}

func (t Output) Print(text string) {
	fmt.Printf(text)
	return
	if fd != nil {
		fd.WriteString(text)
	} else {
		fmt.Print(text)
	}
}

func (t Output) Close() {
	return
	if fd == nil {
		return
	}
	fd.Close()
}
