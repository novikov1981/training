package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	file := os.Args[1:]
	for _, arg := range file {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "arguments is wrong: %v\n", err)
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			url := input.Text()
			go fetch(url, ch)
			f.Close()
		}
	}
	//for _, url := range os.Args[1:] {
	//	go fetch(url, ch)
	//}
	out, err := os.Create("F:/Projects/Go training/Go training Book/Chapter 1.6/yourfile.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer out.Close()
	for _, arg := range file {
		f, _ := os.Open(arg)
		input := bufio.NewScanner(f)
		for input.Scan() {
			out.WriteString(<-ch)
		}

	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs\n %7d %s\n", secs, nbytes, url)

}
