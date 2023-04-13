package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getURLandFilename() (string, string) {
	urlPath := flag.String("u", "", "url")
	flag.Parse()
	_, err := url.Parse(*urlPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*urlPath)
	splitedURL := strings.Split(*urlPath, "/")
	return *urlPath, splitedURL[len(splitedURL)-1]
}

func createFile(filename string) *os.File {
	fmt.Printf("filename: %s\n", filename)
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func getData(urlPath string, client *http.Client, file *os.File) int64 {
	resp, err := client.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return size
}

func main() {
	urlPath, filename := getURLandFilename()
	fmt.Println("Loh")
	file := createFile(filename)
	fmt.Println("Hol")
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	size := getData(urlPath, &client, file)

	fmt.Printf("Downloaded a file %s with size %d", urlPath, size)
}
