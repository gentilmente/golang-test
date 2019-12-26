package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> hola </h1>")
	fmt.Fprintf(w, "<p> %s  </p>", "sarasa")
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	resp.Body.Close()

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
