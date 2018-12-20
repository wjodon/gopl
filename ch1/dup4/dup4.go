package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, fileCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, fileCounts)
			f.Close()
		}
	}
	for line, sf := range fileCounts {
		total := 0
		sublist := ""
		for file, n := range sf {
			if n > 1 {
				sublist += fmt.Sprintf("%s(%d)\t", file, n)
				total += n
			}
		}
		if total > 0 {
			fmt.Printf("%s\ttotal(%d)\t<-\t%s\n", line, total, sublist)
		}
	}
}

func countLines(f *os.File, fileCounts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if fileCounts[input.Text()] == nil {
			fileCounts[input.Text()] = make(map[string]int)
		}
		fileCounts[input.Text()][f.Name()]++
	}
}
