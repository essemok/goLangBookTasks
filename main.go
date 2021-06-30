package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("Please, provide filename as call argument")
		return
	}

	for _, fileName := range files {
		counts := make(map[string]int)
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
