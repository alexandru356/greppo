package main

import (
//	"flag"
	"fmt"
	"os"
	"strings"
)

/*
Phase 2 : add useful flags like -i or -n or -c
-i case insensitivity
-n line numbers
-c count matching lines
-v non-matching lines
and maybe * in the /path/to/file to recursively search
Test and optimize
Format : greppo -flag "arguments" /path/to/file 
*/
func main(){
		
	if len(os.Args) < 3{
		fmt.Println("Usage: greppo <pattern> <filename>")
		os.Exit(1)
	}

	//Access to command-line arguments
	pattern := os.Args[1]
	fileName := os.Args[2]
	
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	
	lines := strings.Split(string(data), "\n")
	for _, line := range lines{

		if strings.Contains(string(line), pattern){
			fmt.Println(line)
		}
	}
}
