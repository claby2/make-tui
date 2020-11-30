package main

import (
	"bufio"
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
	if len(os.Args[1:]) == 0 {
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
			log.Fatal("no makefile found")
			os.Exit(1)
		}
	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		usage := map[string]string{
			"-h, --help": "Print this message and exit",
		}
		fmt.Printf("Usage: make-tui [options] [file]\n" +
			"Options:\n")
		for key, value := range usage {
			fmt.Println("\t" + key + ": " + value + "\n")
		}
		os.Exit(1)
	} else {
		filePath = os.Args[len(os.Args)-1]
	}

	content := NewParsedContent(filePath, getFileContent(filePath))
	content.Parse()

	Render(content)
}
