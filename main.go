package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	httpPrefix = "http://"
)

func main() {
	for _, url := range os.Args[1:] {
		url = checkUrlForPrefix(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch: %v\n", err)
			os.Exit(1)
		}
		_, er := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if er != nil {
			fmt.Printf("fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("HTTP Code: %v", resp.Status)
	}
}

func checkUrlForPrefix(url string) string {
	hasPrefix := strings.HasPrefix(url, httpPrefix)
	if !hasPrefix {
		url = httpPrefix + url
	}

	return url
}
