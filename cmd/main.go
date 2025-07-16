package main

import (
	"flag"
	"fmt"
	"os"
	"greppo/internal/search"
)

/*
Add recursion
Cool features + nice errors in case
Test and optimize
Format : greppo -flag "arguments" /path/to/file
*/
func main(){

	iFlag := flag.Bool("i", false, "case-insensitive match")
	vFlag := flag.Bool("v", false, "non-matching lines")
	cFlag := flag.Bool("c", false, "count matching lines")
	nFlag := flag.Bool("n", false, "matching line with line number")

	flag.Parse()
	
	args := flag.Args()
	if len(os.Args) < 2{
		fmt.Println("Usage: greppo [flags] <pattern> <filename>")
		os.Exit(1)
	}

	pattern := args[0]
	fileName := args[1]
	
	options := search.Options{
		IgnoreCase:  *iFlag,
		Invert:		 *vFlag,	
		CountLines:  *cFlag,
		ShowLineNumber: *nFlag,
	}

	search.Search(fileName, pattern, options)
}
