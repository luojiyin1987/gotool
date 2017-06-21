//
// downloadfile.go
// Copyright (C) 2017 root <root@localhost.localdomain>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	fmt.Println("Download file.....")
	rawURL := "http://104.224.184.178/Planet%20China%20-%20Teaser%20Langde%20-%20Guizhou%20province-65737719.mp4"
	fileURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fmt.Println(segments)
	fileName := segments[2]
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	check := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := check.Get(rawURL)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println(resp.Status)

	size, err := io.Copy(file, resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s with %v bytes download", fileName, size)
}
