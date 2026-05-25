package main

import (
	_ "embed" // ブランクインポート
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed udhr/text/chinese_rights.txt
var chineseRights string

//go:embed udhr/text/english_rights.txt
var englishRights string

//go:embed udhr/text/japanese_rights.txt
var japaneseRights string

//go:embed udhr/text/korean_rights.txt
var koreanRights string

func main() {
	if len(os.Args) < 2 {
		log.Fatal("実行時に言語名を引数として指定してください")
	}

	printRights(os.Args[1])
}

func printRights(lang string) {
	switch strings.ToLower(lang) {
	case "chinese":
		fmt.Println(chineseRights)
	case "english":
		fmt.Println(englishRights)
	case "japanese":
		fmt.Println(japaneseRights)
	case "korean":
		fmt.Println(koreanRights)
	default:
		log.Fatalf("対応していない言語: %s", lang)
	}
}
