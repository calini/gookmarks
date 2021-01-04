package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	// if string passed as param, return them
	if len(os.Args) != 3 {
		log.Fatal("Usage: gookmarks \"Safari Bookmarks.html\" <output.csv>")
	}

	// Open Safari Bookmarks
	bookmarks, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer bookmarks.Close()

	csvFile, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bookmarks)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find("#com\\.apple\\.ReadingList").Next().Children().NextAll().Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		title := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		writer.Write([]string{url, strings.Replace(title, ",", "", -1)})
	})
}
