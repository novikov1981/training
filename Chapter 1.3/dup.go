package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	fmt.Println(files)
	fmt.Println(len(files))
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			for _, arg := range files {
				f, _ := os.Open(arg)
				input := bufio.NewScanner(f)
				for input.Scan() {
					if input.Text() == line {
						fmt.Printf("%d\t%s\t%s\n", n, line, arg)
						f.Close()
					}
				}
			}
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	i := 1
	for input.Scan() {
		counts[input.Text()]++
		i++
		if i > 5 {
			break
		}
	}

}
