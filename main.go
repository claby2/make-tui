package main

import (
	"bufio"
	"log"
	"os"
)

func getContent(filePath string) []string {
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
	if len(os.Args[1:]) == 0 {
		log.Fatal("no input file")
		os.Exit(1)
	}
	filePath := os.Args[1]

	content := NewParsedContent(filePath, getContent(filePath))
	content.Parse()

	Render(content)
}
