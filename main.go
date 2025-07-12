package main

import (
	"flag"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Phase 2 : add useful flags like or -n or -c
-n line numbers
-c count matching lines
and maybe * in the /path/to/file to recursively search
Test and optimize
Format : greppo -flag "arguments" /path/to/file
*/
func main(){

	iFlag := flag.Bool("i", false, "case-insensitive match")
	vFlag := flag.Bool("v", false, "non-matching lines")
	flag.Parse()
	
	args := flag.Args()
	if len(os.Args) < 2{
		fmt.Println("Usage: greppo <pattern> <filename>")
		os.Exit(1)
	}

	pattern := args[0]
	fileName := args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()

		if *iFlag {
			line = strings.ToLower(line)	
			pattern = strings.ToLower(pattern)
		}
		matched := strings.Contains(line, pattern)

		if *vFlag {
			if !matched {
				fmt.Println(line)
			}
		} else {
			if matched {
				fmt.Println(line)
			}
		}
	
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
