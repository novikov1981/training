package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for i, url := range os.Args[1:] {
		if strings.HasPrefix(url, "https://") != true {
			url = "https://" + url
			os.Args[i+1] = url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		for _, url := range os.Args[1:] {
			fmt.Printf("%s\n", url)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch чтение %ы: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("Статус https:// \t%v", resp.Status)

	}
}
