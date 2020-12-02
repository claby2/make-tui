package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func getFileContent(filePath string) []string {
	var fileContent []string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}
	return fileContent
}

func main() {
	var filePath string
	flag.StringVar(&filePath, "f", "", "Parse given file as Makefile")
	helpFlag := flag.Bool("h", false, "Print this message and exit")
	allFlag := flag.Bool("a", false, "Display all targets including special built-in targets")
	flag.Parse()

	if *helpFlag {
		// Print help message
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	} else if filePath == "" {
		// Attempt to find makefile in current directory
		defaultMakefileNames := []string{"GNUmakefile", "makefile", "Makefile"}
		foundFile := false
		for _, name := range defaultMakefileNames {
			if _, err := os.Stat(name); os.IsNotExist(err) == false {
				// File exists
				filePath = name
				foundFile = true
				break
			}
		}
		if !foundFile {
			log.Fatal("no Makefile found")
			os.Exit(1)
		}
	}

	content := NewParsedContent(filePath, getFileContent(filePath))
	if *allFlag {
		content.SetIncludeSpecialTargets(true)
	}
	content.Parse()

	Render(content)
}
