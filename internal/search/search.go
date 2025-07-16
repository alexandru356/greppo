package search

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
)

// Struct to hold the flags for search
type Options struct {
	IgnoreCase bool
	Invert     bool
	CountLines bool
	ShowLineNumber bool
}

func Search(fileName, pattern string, opt Options) {
	
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	matchCount := 0
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		testLine := line
		testPattern := pattern

		if opt.IgnoreCase {
			testLine = strings.ToLower(line)
			testPattern = strings.ToLower(pattern)
		}

		matched := strings.Contains(testLine, testPattern)
		
		if !opt.CountLines && (matched != opt.Invert) {
			matchCount++
			if opt.ShowLineNumber {
				fmt.Printf("%d:%s\n", lineNumber, line)
			} else {
				fmt.Println(line)
			}
		}
		
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if opt.CountLines {
		fmt.Println(matchCount)
	}

}
