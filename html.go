package main

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func getHeadingFromHTML(html string) string {
	r := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "err"
	}
	title := doc.Find("h1").Text()
	return title
}

func getFirstParagraphFromHTML(html string) string {
	r := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "err"
	}
	sel1 := doc.Find("main")
	title := sel1.Find("p").Text()
	return title
}
