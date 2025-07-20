package main

import (
	"flag"
	"fmt"
	"greppo/internal/search"
	"os"
)

func main() {

	iFlag := flag.Bool("i", false, "case-insensitive match")
	vFlag := flag.Bool("v", false, "non-matching lines")
	cFlag := flag.Bool("c", false, "count matching lines")
	nFlag := flag.Bool("n", false, "matching line with line number")
	rFlag := flag.Bool("r", false, "Search all files in a dir/subdir")

	flag.Parse()

	args := flag.Args()
	if len(flag.Args()) < 2 {
		fmt.Println("Usage: greppo [flags] <pattern> <file-or-directory>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	pattern := args[0]
	fileName := args[1]

	options := search.Options{
		IgnoreCase:     *iFlag,
		Invert:         *vFlag,
		CountLines:     *cFlag,
		ShowLineNumber: *nFlag,
		Recursive:      *rFlag,
	}

	if options.Recursive {
		search.SearchRecursive(fileName, pattern, options)
	} else {
		search.Search(fileName, pattern, options)
	}
}
