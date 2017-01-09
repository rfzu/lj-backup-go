package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Download(url string) {
	fmt.Println("donwload", url)
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	fmt.Printf("%T", resp.Body)

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
	default:
		reader = resp.Body
	}

	fmt.Printf("%T", reader)
	body, _ := ioutil.ReadAll(reader)
	fmt.Printf("%T", body)
	ioutil.WriteFile("text.html", body, 0644)
}

func main() {
	Download("http://tema.livejournal.com/2408414.html")
}
