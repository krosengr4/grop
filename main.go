package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Validate command line args
	if len(os.Args) != 3 {
		fmt.Println("ERROR: Unexpected number of arguments!")
		fmt.Println("Usage: ./grop <search_term> <file_path>")
		os.Exit(1)
	}

	searchTerm := os.Args[1]
	filePath := os.Args[2]

	scanFile(searchTerm, filePath)
}

func scanFile(searchTerm, filePath string) {
	// Open the target file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("ERROR opening the file path:  %v\n", err)
	}

	// Scan line-by-line using buffered scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Evaluate string match
		if strings.Contains(line, searchTerm) {
			fmt.Println(line)
		}
	}

	// Catch any errors using scanner
	if err := scanner.Err(); err != nil {
		fmt.Printf("ERROR scanning text: %v\n", err)
		os.Exit(1)
	}
}
