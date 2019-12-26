package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// SitemapIndex urls in xml
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

// News data from xml
type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// NewsMap k,v
type NewsMap struct {
	Keywords string
	Location string
}

func main() {
	var s SitemapIndex
	var n News
	newsMap := make(map[string]NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		Location = strings.TrimSpace(Location)
		resp, err := http.Get(Location)
		if err != nil {
			fmt.Println(err)
		}
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		//fmt.Printf("\n%s", n.Titles)

		for idx, _ := range n.Titles {
			newsMap[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}

	}
	for idx, data := range newsMap {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keywords)
		fmt.Println("\n", data.Location)

	}
}
