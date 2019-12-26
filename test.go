package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type sitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprint(l.Loc)
}

/*
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> hola </h1>")
	fmt.Fprintf(w, "<p> %s  </p>", "sarasa")
}*/

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	//stringBody := string(bytes)
	resp.Body.Close()

	var s sitemapIndex
	xml.Unmarshal(bytes, &s)
	for _, l := range s.Locations {
		fmt.Printf("\n%s", l)

	}

	//http.HandleFunc("/", indexHandler)
	//http.ListenAndServe(":8000", nil)
}
