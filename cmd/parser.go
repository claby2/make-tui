package main

import (
    "bufio"
    "log"
    "os"
    "regexp"
)

type ParsedContent struct {
    filePath string
    content []string
    rules [][]string
}

func NewParsedContent(newFilePath string) *ParsedContent {
    file, err := os.Open(newFilePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var content []string
    scanner := bufio.NewScanner(file)
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    for scanner.Scan() {
        content = append(content, scanner.Text())
    }

    return &ParsedContent{filePath: newFilePath, content: content, rules: [][]string{}}
}

func (parsedContent *ParsedContent) Parse() {
    for _, line := range parsedContent.content {
        r := regexp.MustCompile(`^([^:\s]+)\s*:\s*([^=].*)?$`)
        match := r.FindStringSubmatch(line)
        if match != nil {
            // Match has been found
            parsedContent.rules = append(parsedContent.rules, match[1:])
        }
    }
}
