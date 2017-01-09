package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Read(url string) {
	// from https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("HTML:\n\n", string(bytes))
	ioutil.WriteFile("test_2.html", bytes, 0644)

	resp.Body.Close()
}

func Download(url string) {
	// piece of of shit
	fmt.Println("donwload", url)
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
	default:
		reader = resp.Body
	}

	body, _ := ioutil.ReadAll(reader)
	fmt.Printf("%T", body)
	ioutil.WriteFile("text.html", body, 0644)
}

func main() {
	// Download("http://tema.livejournal.com/2408414.html")
	Read("http://tema.livejournal.com/2408414.html")
}
